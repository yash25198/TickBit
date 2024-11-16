import { useBitcoin } from "../../providers/bitcoin";
import Card from "./Card";
import useEmblaCarousel from "embla-carousel-react";
import { useDotButton } from "./DotButton";
import loadingAnimation from "../../assets/loading.json";
import Lottie from "react-lottie";
import { useEffect } from "react";

const Carousel = () => {
    const { blocks, loading, setPredictionBlockIndex } = useBitcoin();

    const [emblaRef, emblaApi] = useEmblaCarousel({
        loop: false,
        align: "center",
        containScroll: false,
        startIndex: 10,
    });

    const { selectedIndex, onDotButtonClick } = useDotButton(emblaApi);

    useEffect(() => {
        if (!emblaApi) return;
        if (selectedIndex === 10) return;
        setTimeout(() => onDotButtonClick(10), 1000);
    }, []);

    useEffect(() => {
        setPredictionBlockIndex(selectedIndex);
    }, [selectedIndex, setPredictionBlockIndex]);

    if (loading) {
        return (
            <div className="embla__slide">
                <Lottie
                    options={{
                        loop: true,
                        autoplay: true,
                        animationData: loadingAnimation,
                    }}
                    height={160}
                    width={160}
                />
            </div>
        );
    }

    return (
        <div className="relative flex flex-col gap-8 mt-10">
            <h1 className="text-4xl font-bold text-center">
                Predict The Next Bitcoin Block
            </h1>
            <div
                className="embla relative mt-10"
                ref={emblaRef}
            >
                <div className="embla__container items-center">
                    {blocks.map((block, index) => {
                        if (!block) return <></>;

                        return (
                            <Card
                                key={index}
                                block={block!}
                                onClick={() => {
                                    onDotButtonClick(index);
                                }}
                                isActive={
                                    index === emblaApi?.selectedScrollSnap()
                                }
                                isPreviousCardActive={
                                    index - 1 === emblaApi?.selectedScrollSnap()
                                }
                            />
                        );
                    })}
                    {/* {
                        <Lottie
                            options={{
                                loop: true,
                                autoplay: true,
                                animationData: loadingAnimation,
                            }}
                            height={160}
                            width={160}
                        />
                    } */}
                </div>
            </div>
        </div>
    );
};

export default Carousel;
