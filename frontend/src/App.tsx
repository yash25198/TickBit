import React from "react";
import Header from "./components/Header";
import PrivyProvider from "./providers/privy";
import { BitcoinProvider } from "./providers/bitcoin";
import Carousel from "./components/Carousel/Carousel";
import Predict from "./components/Predict";
import Predictions, { Prediction } from "./components/Predictions";

function App() {
    return (
        <PrivyProvider>
            <BitcoinProvider>
                <div className="bg-[#F7F6F9] flex flex-col">
                    <Header />
                    <Carousel />
                    <Predict />
                    <Predictions />
                </div>
            </BitcoinProvider>
        </PrivyProvider>
    );
}

export default App;
