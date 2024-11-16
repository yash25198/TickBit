import { ChevronRight } from "lucide-react";
import React from "react";
import { Block } from "../../providers/bitcoin";

const LeftArrow = () => {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            className="h-4 w-4"
        >
            <path
                d="M15 19l-7-7 7-7"
                stroke="#433C53"
            />
        </svg>
    );
};

const RightArrow = () => {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            className="h-4 w-4"
        >
            <path
                d="M9 5l7 7-7 7"
                stroke="#433C53"
            />
        </svg>
    );
};

const CubeIcon = () => (
    <svg
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        className="h-4 w-4 text-purple-500"
    >
        <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" />
        <path d="m3.3 7 8.7 5 8.7-5" />
        <path d="M12 22V12" />
    </svg>
);

const HashIcon = () => (
    <svg
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        className="h-4 w-4 text-pink-500"
    >
        <line
            x1="4"
            y1="9"
            x2="20"
            y2="9"
        />
        <line
            x1="4"
            y1="15"
            x2="20"
            y2="15"
        />
        <line
            x1="10"
            y1="3"
            x2="8"
            y2="21"
        />
        <line
            x1="16"
            y1="3"
            x2="14"
            y2="21"
        />
    </svg>
);

const ClockIcon = () => (
    <svg
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        className="h-4 w-4 text-orange-500"
    >
        <circle
            cx="12"
            cy="12"
            r="10"
        />
        <path d="M12 6v6l4 2" />
    </svg>
);

const truncateHash = (hash: string) => {
    return hash.slice(0, 6) + "..." + hash.slice(hash.length - 6, hash.length);
};

const Card = ({
    block,
    isActive,
    isPreviousCardActive,
    onClick,
}: {
    block: Block;
    isActive: boolean;
    isPreviousCardActive: boolean;
    onClick: () => void;
}) => {
    return (
        <div className="embla__slide relative">
            {isActive && (
                <div className="absolute top-0 bg-gradient-to-r from-purple-600 via-pink-500 to-orange-400 text-transparent bg-clip-text animate-gradient left-1/2 -translate-x-1/2 text-sm">
                    Selected
                </div>
            )}
            {!isActive && block.isFirstPreviousBlock && (
                <div className="absolute top-0 ml-10 text-[#817A90] flex text-sm items-center gap-2">
                    Previous Blocks
                    <svg
                        width="12"
                        height="8"
                        viewBox="0 0 12 8"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            d="M8.76669 4.66667H1.33335C1.14446 4.66667 0.986131 4.60278 0.858354 4.475C0.730576 4.34723 0.666687 4.18889 0.666687 4C0.666687 3.81111 0.730576 3.65278 0.858354 3.525C0.986131 3.39723 1.14446 3.33334 1.33335 3.33334H8.76669L6.86669 1.43334C6.73335 1.3 6.66947 1.14445 6.67502 0.966669C6.68058 0.788892 6.74446 0.633336 6.86669 0.500003C7.00002 0.366669 7.15835 0.297225 7.34169 0.291669C7.52502 0.286114 7.68335 0.350003 7.81669 0.483336L10.8667 3.53334C10.9334 3.6 10.9806 3.67223 11.0084 3.75C11.0361 3.82778 11.05 3.91111 11.05 4C11.05 4.08889 11.0361 4.17223 11.0084 4.25C10.9806 4.32778 10.9334 4.4 10.8667 4.46667L7.81669 7.51667C7.68335 7.65 7.52502 7.71389 7.34169 7.70834C7.15835 7.70278 7.00002 7.63334 6.86669 7.5C6.74446 7.36667 6.68058 7.21111 6.67502 7.03334C6.66947 6.85556 6.73335 6.7 6.86669 6.56667L8.76669 4.66667Z"
                            fill="#817A90"
                        />
                    </svg>
                </div>
            )}
            <div
                className={`flex relative p-[2.5px] overflow-ellipsis rounded-[1.15rem] mt-8 ${
                    block.isFirstPreviousBlock && "ml-4"
                } ${
                    isActive
                        ? "bg-gradient-to-r from-purple-600 via-pink-500 to-orange-400"
                        : "bg-gray-50"
                }`}
                onClick={() => onClick()}
            >
                <div
                    className={`embla__slide__inner relative flex bg-white py-6 rounded-2xl items-center min-w-[16rem] ${
                        isActive
                            ? "justify-between"
                            : "justify-center scale-x-90 hover:cursor-pointer"
                    } select-none `}
                >
                    {isActive && (
                        <button
                            className="ml-3"
                            // onClick={onLeftClick}
                        >
                            <LeftArrow />
                        </button>
                    )}
                    <div className="flex flex-col items-center gap-2">
                        <div className="flex items-center gap-2 font-medium mb-2">
                            <CubeIcon />
                            {block.height}
                        </div>
                        <div className="flex items-center gap-2">
                            <HashIcon />
                            {(block.id && truncateHash(block.id)) ||
                                "000000...??????"}
                        </div>
                        <div className="flex items-center gap-2 mt-2">
                            <ClockIcon />
                            {block.timestamp
                                ? new Date(
                                      block.timestamp * 1000
                                  ).toLocaleString()
                                : "dd/mm/yy, --:--:--"}
                        </div>
                    </div>
                    {isActive && (
                        <button
                            className="mr-3"
                            // onClick={onRightClick}
                        >
                            <RightArrow />
                        </button>
                    )}
                </div>
            </div>
            {block.isFirstPreviousBlock && (
                <div
                    className={`absolute ${!isActive && "left-[4%]"} ${
                        isPreviousCardActive && "left-[6%]"
                    } h-[calc(100%-2rem)] bottom-0 my-auto w-[1px] border-l-[1px] border-dotted border-gray-500`}
                ></div>
            )}
        </div>
    );
};

export default Card;
