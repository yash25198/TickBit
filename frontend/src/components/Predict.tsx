/* eslint-disable @typescript-eslint/no-unused-expressions */
import React, { useState } from "react";
import {
    Tooltip,
    TooltipContent,
    TooltipProvider,
    TooltipTrigger,
} from "../@/components/ui/tooltip.js";
import { usePrivy, useWallets } from "@privy-io/react-auth";
import tickBitAbi from "../abi/TickBit.json";
import { useBitcoin } from "../providers/bitcoin.js";
import tsd from "../abi/tsd.json";
import {
    createPublicClient,
    createWalletClient,
    custom,
    encodeFunctionData,
} from "viem";
import { networks } from "../providers/privy";

const Predict = () => {
    const [hours1, setHours1] = useState(0);
    const [minutes1, setMinutes1] = useState(0);
    const [seconds1, setSeconds1] = useState(0);

    const [hours2, setHours2] = useState(0);
    const [minutes2, setMinutes2] = useState(0);
    const [seconds2, setSeconds2] = useState(0);

    const { ready } = usePrivy();
    const { wallets } = useWallets();
    const {
        blocks,
        predictionBlockIndex,
        currentTip,
        winAmount,
        setPredictions,
        tokenInfo,
        chainId,
    } = useBitcoin();
    const wallet = wallets[0];

    const hoursLess = hours1 < hours2;
    const hoursEqual = hours1 === hours2;

    const minutesLess = minutes1 < minutes2;
    const minutesEqual = minutes1 === minutes2;

    const secondsLess = seconds1 <= seconds2;

    const isPreviousBlock =
        (currentTip?.height || 0) >=
        (blocks[predictionBlockIndex]?.height || 0);

    let totalStake = 0;

    if (
        !(
            hoursLess ||
            (hoursEqual && minutesLess) ||
            (hoursEqual && minutesEqual && secondsLess)
        )
    ) {
        totalStake = 0;
    } else {
        const time1 = new Date();
        time1.setHours(time1.getHours() + hours1);
        time1.setMinutes(time1.getMinutes() + minutes1);
        time1.setSeconds(time1.getSeconds() + seconds1);

        const time2 = new Date();
        time2.setHours(time2.getHours() + hours2);
        time2.setMinutes(time2.getMinutes() + minutes2);
        time2.setSeconds(time2.getSeconds() + seconds2);

        const unixTimestamp1 = Math.floor(time1.getTime() / 1000);
        const unixTimestamp2 = Math.floor(time2.getTime() / 1000);

        totalStake = unixTimestamp2 - unixTimestamp1 + 1;
    }

    const handleTimeChange = (
        e: React.ChangeEvent<HTMLInputElement>,
        type: string,
        which: "from" | "to"
    ) => {
        const value = parseInt(e.target.value);
        if (type === "hours") {
            if (value > 23) which === "from" ? setHours1(0) : setHours2(0);
            else if (value < 0)
                which === "from" ? setHours1(23) : setHours2(23);
            else which === "from" ? setHours1(value) : setHours2(value);
        }

        if (type === "minutes")
            if (value > 59) which === "from" ? setMinutes1(0) : setMinutes2(0);
            else if (value < 0)
                which === "from" ? setMinutes1(59) : setMinutes2(59);
            else which === "from" ? setMinutes1(value) : setMinutes2(value);

        if (type === "seconds")
            if (value > 59) which === "from" ? setSeconds1(0) : setSeconds2(0);
            else if (value < 0)
                which === "from" ? setSeconds1(59) : setSeconds2(59);
            else which === "from" ? setSeconds1(value) : setSeconds2(value);
    };

    const submitPrediction = async () => {
        if (!wallet || !ready || !tokenInfo) return;
        if (
            !(
                hoursLess ||
                (hoursEqual && minutesLess) ||
                (hoursEqual && minutesEqual && secondsLess)
            )
        ) {
            return alert(
                "Invalid Range: Please make sure range2 is greater than range1"
            );
        }

        const time1 = new Date();
        time1.setHours(time1.getHours() + hours1);
        time1.setMinutes(time1.getMinutes() + minutes1);
        time1.setSeconds(time1.getSeconds() + seconds1);

        const time2 = new Date();
        time2.setHours(time2.getHours() + hours2);
        time2.setMinutes(time2.getMinutes() + minutes2);
        time2.setSeconds(time2.getSeconds() + seconds2);

        const unixTimestamp1 = Math.floor(time1.getTime() / 1000);
        const unixTimestamp2 = Math.floor(time2.getTime() / 1000);

        const timestamps = Array.from({
            length: unixTimestamp2 - unixTimestamp1 + 1,
        })
            .fill(0)
            .map((_, index) => unixTimestamp1 + index);

        const provider = await wallet.getEthereumProvider();
        const network = networks.find((n) => n.id === chainId);
        if (!network) return;
        const { contractAddress, ...newNetwork } = { ...network };

        const walletClient = createWalletClient({
            chain: newNetwork,
            transport: custom(provider),
        });

        const approveData = encodeFunctionData({
            abi: tsd,
            functionName: "approve",
            args: [
                contractAddress as `0x${string}`,
                BigInt(totalStake) * BigInt(10 ** tokenInfo.decimals),
            ],
        });

        const approveTxRequest = {
            to: tokenInfo.address as `0x${string}`,
            data: approveData,
        };

        const hash = await walletClient.sendTransaction({
            ...approveTxRequest,
            account: wallet.address as `0x${string}`,
        });

        const publicClient = createPublicClient({
            chain: newNetwork,
            transport: custom(provider),
        });

        await publicClient.waitForTransactionReceipt({
            hash,
        });

        const data = encodeFunctionData({
            abi: tickBitAbi,
            functionName: "bet",
            args: [timestamps, blocks[predictionBlockIndex].height],
        });

        const transactionRequest = {
            to: contractAddress as `0x${string}`,
            data,
        };

        await walletClient.sendTransaction({
            ...transactionRequest,
            account: wallet.address as `0x${string}`,
        });

        setPredictions((predictions) => {
            return [
                ...predictions,
                {
                    Timestamps: timestamps,
                    Addr: wallet.address as `0x${string}`,
                },
            ];
        });
    };

    if (isPreviousBlock) {
        return <></>;
    }
    return (
        <div className="relative w-[90%] md:w-[50%] lg:w-[40%] mx-auto bg-white rounded-xl mt-8 p-6 flex flex-col font-semibold ">
            <h3 className="relative text-sm text-[#817A90] text-center -z-10">
                The selected block confirmation is in:
            </h3>
            <div className="flex justify-between max-w-[100%] gap-3 mt-4">
                <div className="flex flex-col max-w-[30%] bg-[#F7F6F9] p-4 rounded-xl">
                    <label className="text-[#817A90]">Hours</label>
                    <input
                        className="bg-[#F7F6F9] text-lg outline-none text-[#433C53]"
                        type="number"
                        value={hours1.toString(10).padStart(2, "0")}
                        onChange={(e) => handleTimeChange(e, "hours", "from")}
                        placeholder="00"
                    />
                </div>
                <div className="flex flex-col max-w-[30%] bg-[#F7F6F9] p-4 rounded-xl">
                    <label className="text-[#817A90]">Minutes</label>
                    <input
                        className="bg-[#F7F6F9] outline-none text-lg text-[#433C53]"
                        type="number"
                        value={minutes1.toString(10).padStart(2, "0")}
                        onChange={(e) => handleTimeChange(e, "minutes", "from")}
                        placeholder="05"
                    />
                </div>
                <div className="flex flex-col max-w-[30%] bg-[#F7F6F9] p-4 outline-none rounded-xl">
                    <label className="text-[#817A90]">Seconds</label>
                    <input
                        className="bg-[#F7F6F9] outline-none text-lg text-[#433C53]"
                        type="number"
                        value={seconds1.toString(10).padStart(2, "0")}
                        onChange={(e) => handleTimeChange(e, "seconds", "from")}
                        placeholder="35"
                    />
                </div>
            </div>

            <h3 className="text-sm text-[#817A90] mt-4 text-center font-semibold">
                to
            </h3>
            <div className="flex justify-between max-w-[100%] gap-3 mt-4">
                <div className="flex flex-col max-w-[30%] bg-[#F7F6F9] p-4 rounded-xl">
                    <label className="text-[#817A90]">Hours</label>
                    <input
                        className="bg-[#F7F6F9] text-lg outline-none text-[#433C53]"
                        type="number"
                        value={hours2.toString(10).padStart(2, "0")}
                        onChange={(e) => handleTimeChange(e, "hours", "to")}
                        placeholder="00"
                    />
                </div>
                <div className="flex flex-col max-w-[30%] bg-[#F7F6F9] p-4 rounded-xl">
                    <label className="text-[#817A90]">Minutes</label>
                    <input
                        className="bg-[#F7F6F9] outline-none text-lg text-[#433C53]"
                        type="number"
                        value={minutes2.toString(10).padStart(2, "0")}
                        onChange={(e) => handleTimeChange(e, "minutes", "to")}
                        placeholder="05"
                    />
                </div>
                <div className="flex flex-col max-w-[30%] bg-[#F7F6F9] p-4 outline-none rounded-xl">
                    <label className="text-[#817A90]">Seconds</label>
                    <input
                        className="bg-[#F7F6F9] outline-none text-lg text-[#433C53]"
                        type="number"
                        value={seconds2.toString(10).padStart(2, "0")}
                        onChange={(e) => handleTimeChange(e, "seconds", "to")}
                        placeholder="35"
                    />
                </div>
            </div>

            <div className="flex justify-between mt-4">
                <div className="flex gap-2">
                    <div className="text-[#817A90] font-normal">
                        Total Stake
                    </div>
                    <TooltipProvider>
                        <Tooltip>
                            <TooltipTrigger>
                                <svg
                                    width="16"
                                    height="16"
                                    viewBox="0 0 16 16"
                                    fill="none"
                                    xmlns="http://www.w3.org/2000/svg"
                                >
                                    <mask
                                        id="mask0_67_171"
                                        style={{
                                            background: "mask-type:alpha",
                                        }}
                                        maskUnits="userSpaceOnUse"
                                        x="0"
                                        y="0"
                                        width="16"
                                        height="16"
                                    >
                                        <rect
                                            width="16"
                                            height="16"
                                            fill="#D9D9D9"
                                        />
                                    </mask>
                                    <g mask="url(#mask0_67_171)">
                                        <path
                                            d="M7.33319 11.3336H8.66652V7.33362H7.33319V11.3336ZM7.99986 6.00028C8.18875 6.00028 8.34708 5.9364 8.47486 5.80862C8.60264 5.68084 8.66652 5.52251 8.66652 5.33362C8.66652 5.14473 8.60264 4.9864 8.47486 4.85862C8.34708 4.73084 8.18875 4.66695 7.99986 4.66695C7.81097 4.66695 7.65264 4.73084 7.52486 4.85862C7.39708 4.9864 7.33319 5.14473 7.33319 5.33362C7.33319 5.52251 7.39708 5.68084 7.52486 5.80862C7.65264 5.9364 7.81097 6.00028 7.99986 6.00028ZM7.99986 14.667C7.07763 14.667 6.21097 14.492 5.39986 14.142C4.58875 13.792 3.88319 13.317 3.28319 12.717C2.68319 12.117 2.20819 11.4114 1.85819 10.6003C1.50819 9.78917 1.33319 8.92251 1.33319 8.00028C1.33319 7.07806 1.50819 6.2114 1.85819 5.40028C2.20819 4.58917 2.68319 3.88362 3.28319 3.28362C3.88319 2.68362 4.58875 2.20862 5.39986 1.85862C6.21097 1.50862 7.07763 1.33362 7.99986 1.33362C8.92208 1.33362 9.78875 1.50862 10.5999 1.85862C11.411 2.20862 12.1165 2.68362 12.7165 3.28362C13.3165 3.88362 13.7915 4.58917 14.1415 5.40028C14.4915 6.2114 14.6665 7.07806 14.6665 8.00028C14.6665 8.92251 14.4915 9.78917 14.1415 10.6003C13.7915 11.4114 13.3165 12.117 12.7165 12.717C12.1165 13.317 11.411 13.792 10.5999 14.142C9.78875 14.492 8.92208 14.667 7.99986 14.667ZM7.99986 13.3336C9.48875 13.3336 10.7499 12.817 11.7832 11.7836C12.8165 10.7503 13.3332 9.48917 13.3332 8.00028C13.3332 6.5114 12.8165 5.25028 11.7832 4.21695C10.7499 3.18362 9.48875 2.66695 7.99986 2.66695C6.51097 2.66695 5.24986 3.18362 4.21652 4.21695C3.18319 5.25028 2.66652 6.5114 2.66652 8.00028C2.66652 9.48917 3.18319 10.7503 4.21652 11.7836C5.24986 12.817 6.51097 13.3336 7.99986 13.3336Z"
                                            fill="#817A90"
                                        />
                                    </g>
                                </svg>
                            </TooltipTrigger>
                            <TooltipContent>
                                <span className="text-[#433C53]">
                                    For every second.
                                </span>
                            </TooltipContent>
                        </Tooltip>
                    </TooltipProvider>
                </div>
                <div className="text-[#433C53]">
                    {totalStake} {tokenInfo?.symbol || "USDC"}
                </div>
            </div>

            <div className="flex justify-between mt-2">
                <div className="text-[#817A90] font-normal">Potential Win</div>
                <div className="bg-gradient-to-r from-purple-600 via-pink-500 to-orange-400 text-transparent bg-clip-text animate-gradient">
                    {totalStake + winAmount} {tokenInfo?.symbol || "USDC"}
                </div>
            </div>

            <button
                className="bg-gradient-to-r from-purple-600 via-pink-500 to-orange-400 text-white font-semibold mt-4 p-4 rounded-xl"
                onClick={submitPrediction}
            >
                Submit Prediction
            </button>
        </div>
    );
};

export default Predict;
