export const trimBlockHash = (hash: string) =>
    hash.slice(0, 10) + "..." + hash.slice(-10);

export function truncateAddress(address: `0x${string}`) {
    if (!address) return "";
    return `${address.slice(0, 6)}...${address.slice(-4)}`;
}

export const chainIdToLink = (chainId: number): `${string}/` => {
    switch (chainId) {
        case 44_787:
            return "https://celo-alfajores.blockscout.com/";
        case 23_295:
            return "https://explorer.oasis.io/mainnet/sapphire/";
        case 48_899:
            return "https://evm-testnet.flowscan.io/";
        case 84_532:
            return "https://base-sepolia.blockscout.com/";
        case 534_351:
            return "https://scroll-sepolia.blockscout.com/";
        case 1_442:
            return "https://explorer-ui.cardona.zkevm-rpc.com/";
        case 2_810:
            return "https://explorer-holesky.morphl2.io/";
        default:
            throw new Error("unreachable code");
    }
};
