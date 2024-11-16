import React, { useEffect, useState } from "react";
import { usePrivy, useWallets } from "@privy-io/react-auth";
import { truncateAddress } from "../utils.js";
import {
    createWalletClient,
    custom,
    decodeFunctionResult,
    encodeFunctionData,
    erc20Abi,
} from "viem";
import { useBitcoin } from "../providers/bitcoin.js";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "../@/components/ui/dropdown-menu.js";
import { Button } from "../@/components/ui/button.js";
import { Check, ChevronDown } from "lucide-react";
import { networks } from "../providers/privy.js";

const WalletIcon = () => {
    return (
        <svg
            width="20"
            height="20"
            viewBox="0 0 20 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <mask
                id="mask0_69_174"
                style={{ background: "mask-type:alpha" }}
                maskUnits="userSpaceOnUse"
                x="0"
                y="0"
                width="20"
                height="20"
            >
                <rect
                    width="20"
                    height="20"
                    fill="#D9D9D9"
                />
            </mask>
            <g mask="url(#mask0_69_174)">
                <path
                    d="M4.16667 17.5C3.70833 17.5 3.31597 17.3368 2.98958 17.0104C2.66319 16.684 2.5 16.2917 2.5 15.8333V4.16667C2.5 3.70833 2.66319 3.31597 2.98958 2.98958C3.31597 2.66319 3.70833 2.5 4.16667 2.5H15.8333C16.2917 2.5 16.684 2.66319 17.0104 2.98958C17.3368 3.31597 17.5 3.70833 17.5 4.16667V6.25H15.8333V4.16667H4.16667V15.8333H15.8333V13.75H17.5V15.8333C17.5 16.2917 17.3368 16.684 17.0104 17.0104C16.684 17.3368 16.2917 17.5 15.8333 17.5H4.16667ZM10.8333 14.1667C10.375 14.1667 9.98264 14.0035 9.65625 13.6771C9.32986 13.3507 9.16667 12.9583 9.16667 12.5V7.5C9.16667 7.04167 9.32986 6.64931 9.65625 6.32292C9.98264 5.99653 10.375 5.83333 10.8333 5.83333H16.6667C17.125 5.83333 17.5174 5.99653 17.8438 6.32292C18.1701 6.64931 18.3333 7.04167 18.3333 7.5V12.5C18.3333 12.9583 18.1701 13.3507 17.8438 13.6771C17.5174 14.0035 17.125 14.1667 16.6667 14.1667H10.8333ZM16.6667 12.5V7.5H10.8333V12.5H16.6667ZM13.3333 11.25C13.6806 11.25 13.9757 11.1285 14.2188 10.8854C14.4618 10.6424 14.5833 10.3472 14.5833 10C14.5833 9.65278 14.4618 9.35764 14.2188 9.11458C13.9757 8.87153 13.6806 8.75 13.3333 8.75C12.9861 8.75 12.691 8.87153 12.4479 9.11458C12.2049 9.35764 12.0833 9.65278 12.0833 10C12.0833 10.3472 12.2049 10.6424 12.4479 10.8854C12.691 11.1285 12.9861 11.25 13.3333 11.25Z"
                    fill="#817A90"
                />
            </g>
        </svg>
    );
};

const Twitter = () => {
    return (
        <svg
            stroke="currentColor"
            fill="currentColor"
            stroke-width="0"
            viewBox="0 0 512 512"
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
        >
            <path d="M459.37 151.716c.325 4.548.325 9.097.325 13.645 0 138.72-105.583 298.558-298.558 298.558-59.452 0-114.68-17.219-161.137-47.106 8.447.974 16.568 1.299 25.34 1.299 49.055 0 94.213-16.568 130.274-44.832-46.132-.975-84.792-31.188-98.112-72.772 6.498.974 12.995 1.624 19.818 1.624 9.421 0 18.843-1.3 27.614-3.573-48.081-9.747-84.143-51.98-84.143-102.985v-1.299c13.969 7.797 30.214 12.67 47.431 13.319-28.264-18.843-46.781-51.005-46.781-87.391 0-19.492 5.197-37.36 14.294-52.954 51.655 63.675 129.3 105.258 216.365 109.807-1.624-7.797-2.599-15.918-2.599-24.04 0-57.828 46.782-104.934 104.934-104.934 30.213 0 57.502 12.67 76.67 33.137 23.715-4.548 46.456-13.32 66.599-25.34-7.798 24.366-24.366 44.833-46.132 57.827 21.117-2.273 41.584-8.122 60.426-16.243-14.292 20.791-32.161 39.308-52.628 54.253z"></path>
        </svg>
    );
};

const Header: React.FC = () => {
    const { ready, login, user, authenticated, logout } = usePrivy();
    const [loading, setLoading] = useState(false);
    const { wallets } = useWallets();
    const [balance, setBalance] = useState(0);
    const { tokenInfo, setChainId, chainId } = useBitcoin();
    // const [selectedNetwork, setSelectedNetwork] = useState(celoAlfajores);
    const selectedNetwork = networks.find((n) => n.id === chainId);

    const wallet = wallets[0];

    useEffect(() => {
        if (!user || !tokenInfo) return;

        (async () => {
            while (!ready || !wallet) {
                await new Promise((resolve) => setTimeout(resolve, 100));
            }

            const provider = await wallet.getEthereumProvider();
            const data = encodeFunctionData({
                abi: erc20Abi,
                functionName: "balanceOf",
                args: [wallet.address as `0x${string}`],
            });

            const transactionRequest = {
                to: tokenInfo.address,
                data,
            };
            const result = await provider.request({
                method: "eth_call",
                params: [transactionRequest],
            });

            const decodedResult = decodeFunctionResult({
                abi: erc20Abi,
                functionName: "balanceOf",
                data: result,
            });

            const decimals = tokenInfo.decimals;
            const humanReadableBalance =
                decodedResult / BigInt(Math.pow(10, decimals));

            setBalance(Number(humanReadableBalance));

            console.log(result);
        })();
    }, [user, ready, tokenInfo, wallet]);

    const handleNetworkChange = async (chainId: number) => {
        if (!wallet) return;

        try {
            await wallet.switchChain(chainId);
            setChainId(chainId);
        } catch (error) {
            console.log("here");
            if ((error as { code: number }).code === -32002) return;
            try {
                const network = networks.find((n) => n.id === chainId);
                if (!network) return;

                const provider = await wallet.getEthereumProvider();

                const walletClient = createWalletClient({
                    chain: network,
                    transport: custom(provider),
                });

                await walletClient.addChain({
                    chain: network,
                });
                console.log("succewss");
                // await wallet.switchChain(chainId);
                // setChainId(chainId);
            } catch (error) {
                console.error(error);
            }
        }
    };

    return (
        <header className="flex items-center justify-between p-4 z-10">
            {/* Left Side */}
            <div className="flex items-center space-x-4">
                <h1 className="text-lg font-bold text-black">
                    <span className="bg-gradient-to-r from-purple-600 via-pink-500 to-orange-400 text-transparent bg-clip-text animate-gradient">
                        Tick
                    </span>
                    <span>Bit</span>
                </h1>
                <a
                    className="text-gray-500 hover:text-gray-700"
                    href="https://x.com/CremaLabs"
                    target="_blank"
                    rel="noreferrer"
                >
                    <Twitter />
                </a>
            </div>

            {!authenticated ? (
                <div className="flex items-center space-x-4">
                    <button
                        className="bg-gradient-to-r from-purple-600 via-pink-500 to-orange-400 pt-3 pb-3 pl-6 pr-6 rounded-full text-white font-semibold"
                        onClick={() => {
                            setLoading(true);
                            while (!ready);
                            login({ loginMethods: ["wallet"] });
                            setLoading(false);
                        }}
                        disabled={loading}
                    >
                        Connect Wallet
                    </button>
                </div>
            ) : (
                <div className="flex items-center gap-4  text-[#433C53] z-10">
                    <div className="flex gap-2 bg-white py-3 px-6 rounded-full items-center">
                        <WalletIcon />
                        <div className="flex gap-1">
                            <span>{balance}</span>
                            <span>{tokenInfo?.symbol || "USDC"}</span>
                        </div>
                    </div>
                    <button
                        className="bg-white pt-3 pb-3 pl-6 pr-6 rounded-full font-semibold "
                        onClick={async () => {
                            setLoading(true);
                            await logout();
                            setLoading(false);
                        }}
                        disabled={loading}
                    >
                        {truncateAddress(
                            user?.wallet?.address as `0x${string}`
                        ) || "Log Out"}
                    </button>
                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <Button
                                variant="outline"
                                className="w-48 justify-between z-10"
                            >
                                <span className="flex items-center gap-2">
                                    {selectedNetwork?.name}
                                </span>
                                <ChevronDown className="h-4 w-4 opacity-50" />
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent className="w-48">
                            {networks.map((network) => (
                                <DropdownMenuItem
                                    key={network.id}
                                    className="flex items-center justify-between"
                                    onClick={() =>
                                        handleNetworkChange(network.id)
                                    }
                                >
                                    <span className="flex items-center gap-2">
                                        {network.name}
                                    </span>
                                    {selectedNetwork?.id === network.id && (
                                        <Check className="h-4 w-4" />
                                    )}
                                </DropdownMenuItem>
                            ))}
                        </DropdownMenuContent>
                    </DropdownMenu>
                </div>
            )}
        </header>
    );
};

export default Header;
