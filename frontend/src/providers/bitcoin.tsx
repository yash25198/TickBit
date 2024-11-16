import {
    createContext,
    useCallback,
    useContext,
    useEffect,
    useState,
} from "react";
import axios from "axios";
import { Prediction } from "../components/Predictions.js";
import config from "../config.json";
import { useWallets } from "@privy-io/react-auth";
import { decodeFunctionResult, encodeFunctionData, erc20Abi } from "viem";
import tickBitAbi from "../abi/TickBit.json";
import { celoAlfajores } from "viem/chains";
import { networks } from "../providers/privy";

export interface Block {
    id: string | null;
    height: number;
    timestamp: number | null;
    previousblockhash: string | null;
    isFirstPreviousBlock: boolean;
}

export interface TokenInfo {
    address: `0x${string}`;
    symbol: string;
    decimals: number;
}

interface BitcoinContextType {
    currentTip: Block | null;
    loading: boolean;
    blocks: Block[];
    predictionBlockIndex: number;
    setPredictionBlockIndex: React.Dispatch<React.SetStateAction<number>>;
    setBlocks: React.Dispatch<React.SetStateAction<Block[]>>;
    winAmount: number;
    setWinAmount: React.Dispatch<React.SetStateAction<number>>;
    predictions: Prediction[];
    setPredictions: React.Dispatch<React.SetStateAction<Prediction[]>>;
    tokenInfo: TokenInfo | undefined;
    setTokenInfo: React.Dispatch<React.SetStateAction<TokenInfo | undefined>>;
    chainId: number;
    setChainId: React.Dispatch<React.SetStateAction<number>>;
}

const BitcoinContext = createContext<BitcoinContextType | undefined>(undefined);

export const BitcoinProvider = ({
    children,
}: {
    children: React.ReactNode;
}) => {
    const [currentTip, setCurrentTip] = useState<Block | null>(null);
    //prediction block is the block you want to show predictions for
    const [blocks, setBlocks] = useState<Block[]>([]);
    const [predictionBlockIndex, setPredictionBlockIndex] =
        useState<number>(10);
    const [winAmount, setWinAmount] = useState(0);
    const [loading, setLoading] = useState(true);
    const [predictions, setPredictions] = useState<Prediction[]>([]);
    const [tokenInfo, setTokenInfo] = useState<TokenInfo>();
    const { wallets, ready } = useWallets();
    const [chainId, setChainId] = useState<number>(celoAlfajores.id);
    const wallet = wallets[0];

    useEffect(() => {
        const fetchTokenInfo = async () => {
            if (!wallet || !ready) return;
            const provider = await wallet.getEthereumProvider();
            const network = networks.find((n) => n.id === chainId);
            if (!network) return;
            const tokenAddress = decodeFunctionResult({
                abi: tickBitAbi,
                functionName: "token",
                data: await provider.request({
                    method: "eth_call",
                    params: [
                        {
                            to: network.contractAddress,
                            data: encodeFunctionData({
                                abi: tickBitAbi,
                                functionName: "token",
                                args: [],
                            }),
                        },
                        "latest",
                    ],
                }),
            }) as `0x${string}`;

            console.log(tokenAddress);

            const tokenSymbol = decodeFunctionResult({
                abi: erc20Abi,
                functionName: "symbol",
                data: await provider.request({
                    method: "eth_call",
                    params: [
                        {
                            to: tokenAddress,
                            data: encodeFunctionData({
                                abi: erc20Abi,
                                functionName: "symbol",
                                args: [],
                            }),
                        },
                        "latest",
                    ],
                }),
            }) as string;

            const tokenDecimals = decodeFunctionResult({
                abi: erc20Abi,
                functionName: "decimals",
                data: await provider.request({
                    method: "eth_call",
                    params: [
                        {
                            to: tokenAddress,
                            data: encodeFunctionData({
                                abi: erc20Abi,
                                functionName: "decimals",
                                args: [],
                            }),
                        },
                        "latest",
                    ],
                }),
            }) as number;

            const tokenInfo = {
                address: tokenAddress,
                symbol: tokenSymbol,
                decimals: Number(tokenDecimals),
            };

            setTokenInfo(tokenInfo);
        };
        fetchTokenInfo();
    }, [chainId, ready, wallet]);

    const fetchPredictions = useCallback(async () => {
        const height = blocks[predictionBlockIndex]?.height;
        if (!height) return;
        if (!tokenInfo) return;

        try {
            const response = await axios.get(
                `${config.predictionsApi}${height}&chain=${chainId}`
            );

            if (predictions == response.data.bets) return;

            setPredictions(response.data.bets || []);
            setWinAmount(
                Number(
                    BigInt(response.data.pool.AcruedAmount) /
                        BigInt(10 ** tokenInfo.decimals)
                )
            );
        } catch (error) {
            console.error("Error fetching predictions:", error);
        }
    }, [blocks, predictionBlockIndex, tokenInfo, chainId, predictions]);

    useEffect(() => {
        fetchPredictions();
    }, [fetchPredictions]);

    const initializeState = useCallback(async () => {
        const blocks = await axios.get("https://mempool.space/api/blocks/tip");

        const predictionBlockNumber: number = blocks.data[0].height + 1;
        const newCurrentTip = blocks.data[0];

        if (newCurrentTip.height == currentTip?.height) return;

        setCurrentTip(newCurrentTip);

        const futureBlocks = new Array(10)
            .fill(0)
            .map(
                (_, index) =>
                    ({
                        height: predictionBlockNumber + index + 1,
                        timestamp: null,
                        id: null,
                    } as Block)
            )
            .reverse();

        const predictionBlock = {
            height: predictionBlockNumber,
            timestamp: null,
            id: null,
            previousblockhash: null,
        };

        const previousBlocks = await axios.get(
            `https://mempool.space/api/blocks/${predictionBlockNumber - 1}`
        );
        previousBlocks.data[0].isFirstPreviousBlock = true;

        setBlocks([...futureBlocks, predictionBlock, ...previousBlocks.data]);
    }, [currentTip?.height]);

    useEffect(() => {
        const interval = setInterval(() => {
            (async () => {
                const blocks = await axios.get(
                    "https://mempool.space/api/blocks/tip"
                );
                if (blocks.data[0].height === currentTip?.height) return;
                await initializeState();
            })();
        }, 5000);
        return () => clearInterval(interval);
    }, [currentTip?.height, initializeState]);

    useEffect(() => {
        (async () => {
            await initializeState();

            setLoading(false);
        })();
    }, [currentTip, initializeState]);

    return (
        <BitcoinContext.Provider
            value={{
                currentTip,
                predictionBlockIndex,
                setPredictionBlockIndex,
                predictions,
                setPredictions,
                winAmount,
                setWinAmount,
                loading,
                blocks,
                setBlocks,
                tokenInfo,
                setTokenInfo,
                chainId,
                setChainId,
            }}
        >
            {children}
        </BitcoinContext.Provider>
    );
};

export const useBitcoin = (): BitcoinContextType => {
    const context = useContext(BitcoinContext);
    if (context === undefined) {
        throw new Error("useBitcoin must be used within a BitcoinProvider");
    }
    return context;
};
