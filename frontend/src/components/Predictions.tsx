import { formatDate } from "date-fns";
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "../@/components/ui/table";
import { truncateAddress } from "../utils";
import { useBitcoin } from "../providers/bitcoin";

export type Prediction = {
    Addr: `0x${string}`; //who bet
    Timestamps: number[];
    PlacedAt: number;
    Raw: { transactionHash: `0x${string}` };
};

type UsablePrediction = {
    address: `0x${string}`; //who bet
    amount: number;
    timestamp1: number;
    timestamp2: number;
};

const Predictions = () => {
    const { predictions, currentTip, blocks, predictionBlockIndex, tokenInfo } =
        useBitcoin();

    const isPreviousBlock =
        (currentTip?.height || 0) >=
        (blocks[predictionBlockIndex]?.height || 0);

    const usablePredictions: UsablePrediction[] = (predictions || [])
        .map((prediction) => ({
            ...prediction,
            amount: prediction.Timestamps.length,
            timestamps: prediction.Timestamps.sort((a, b) => a - b),
        }))
        .map((prediction) => ({
            ...prediction,
            address: prediction.Addr,
            timestamp1: prediction.timestamps[0],
            timestamp2: prediction.timestamps[prediction.timestamps.length - 1],
        }));

    return (
        <div className="bg-white w-[90%] md:w-[50%] lg:w-[40%] mx-auto rounded-xl mt-4 p-4">
            <h2 className="text-[#433C53] text-center font-bold">
                See other's predictions
            </h2>
            {usablePredictions.length == 0 ? (
                <span className="flex justify-center text-[#817A90] mt-4 text-center">
                    {!isPreviousBlock
                        ? "No predictions yet. Be the first one to bet!"
                        : "No predictions for this block."}
                </span>
            ) : (
                <div className="bg-white p-4">
                    <Table>
                        <TableHeader>
                            <TableRow className="border-black">
                                {/* <TableHead className="w-2/3 bg-gradient-to-r from-purple-600 to-pink-500 text-transparent bg-clip-text animate-gradient font-bold"> */}
                                <TableHead className="w-4/5 text-[#817A90] font-semibold">
                                    Predictions
                                </TableHead>
                                {/* <TableHead className="bg-gradient-to-r from-pink-500 to-orange-400 text-transparent bg-clip-text animate-gradient font-bold"> */}
                                <TableHead className="text-[#817A90] font-semibold">
                                    Amount
                                </TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {usablePredictions.map((prediction, index) => (
                                <TableRow
                                    key={index}
                                    className="border-black"
                                >
                                    <TableCell className="text-left">
                                        <div className="space-y-1">
                                            <div className="flex items-center gap-2 text-base">
                                                <span className="font-semibold">
                                                    {truncateAddress(
                                                        prediction.address
                                                    )}
                                                </span>
                                            </div>
                                            <p className="text-sm ">
                                                Predicted:{" "}
                                                <span className="bg-gradient-to-r from-purple-600 via-pink-500 to-orange-400 text-transparent bg-clip-text animate-gradient">
                                                    {formatDate(
                                                        prediction.timestamp1,
                                                        "dd/MM/yyyy HH:mm:ss"
                                                    )}{" "}
                                                    -{" "}
                                                    {formatDate(
                                                        prediction.timestamp2,
                                                        "dd/MM/yyyy HH:mm:ss"
                                                    )}
                                                </span>
                                            </p>
                                        </div>
                                    </TableCell>
                                    <TableCell className="text-center font-bold text-sm text-[#433C53] ">
                                        {prediction.amount}{" "}
                                        {tokenInfo?.symbol || "USDC"}
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </div>
            )}
        </div>
    );
};

export default Predictions;
