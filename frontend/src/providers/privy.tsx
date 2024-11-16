import { PrivyProvider } from "@privy-io/react-auth";
import {
    celoAlfajores,
    sapphireTestnet,
    zircuitTestnet,
    baseSepolia,
    polygonZkEvmTestnet,
    morphHolesky,
    scrollSepolia,
} from "viem/chains";

export const networks = [
    {
        ...celoAlfajores,
        contractAddress: "0xfb8cf1e65b8d5a699ffa86235941319d44156701",
    },
    {
        ...sapphireTestnet,
        contractAddress: "0xfb8cf1e65b8d5a699ffa86235941319d44156701",
    },
    {
        ...zircuitTestnet,
        contractAddress: "0xfb8cf1e65b8d5a699ffa86235941319d44156701",
    },
    {
        ...baseSepolia,
        contractAddress: "0xfb8cf1e65b8d5a699ffa86235941319d44156701",
    },
    {
        ...scrollSepolia, 
        contractAddress: "0xfb8cf1e65b8d5a699ffa86235941319d44156701",   
    }
    {
        ...polygonZkEvmTestnet,
        contractAddress: "0xfb8cf1e65b8d5a699ffa86235941319d44156701",
    },
    {
        ...morphHolesky,
        contractAddress: "0xfb8cf1e65b8d5a699ffa86235941319d44156701",
    },
];

export default function Privyprovider({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <PrivyProvider
            appId="cm3jk7wbx00b9vmvuysteb3j4"
            config={{
                defaultChain: celoAlfajores,
                supportedChains: networks,
                appearance: {
                    theme: "light",
                    accentColor: "#FB923C",
                },
            }}
        >
            {children}
        </PrivyProvider>
    );
}
