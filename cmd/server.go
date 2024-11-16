package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crema-labs/sxg-go/internal/server"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type Env struct {
	SuccinctPivKey string
	EthClientUrl   string
	TBAddress      common.Address
}

func LookupEnv(k string) string {
	v, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Errorf("environment variable %s is not set", k))
	}
	return v
}

func Getenv() Env {
	return Env{
		EthClientUrl:   LookupEnv("ETH_CLIENT_URL"),
		TBAddress:      common.HexToAddress(LookupEnv("TB_ADDRESS")),
		SuccinctPivKey: LookupEnv("SP1_PRIVATE_KEY"),
	}
}

func main() {
	env := Getenv()
	os.Setenv("RUST_LOG", "info")
	os.Setenv("SP1_PROVER", "network")
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to create logger")
	}

	ethClient, err := ethereum.NewClient(logger, env.EthClientUrl)
	if err != nil {
		panic(err)
	}

	tbClient := ethereum.NewTBClient(ethClient, nil, env.TBAddress, logger.With(zap.String("sub-service", "spvclient")))
	// Create a new server
	srv := server.NewServer(env.SuccinctPivKey, tbClient, logger)
	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
