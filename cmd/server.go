package main

import (
	"log"

	"github.com/crema-labs/sxg-go/cmd/utils"
	"github.com/crema-labs/sxg-go/internal/server"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

func main() {
	env := utils.Getenv()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to create logger")
	}

	cp := make(map[string]ethereum.ChainParams)
	for chain, params := range env.ChainParams {
		ethClient, err := ethereum.NewClient(logger, params.EthClientUrl)
		if err != nil {
			panic(err)
		}
		cp[chain] = ethereum.ChainParams{
			TBAddress: common.HexToAddress(params.TBAddress),
			EthClient: ethClient,
		}

	}

	tbClient := ethereum.NewTBClient(cp, nil, logger.With(zap.String("sub-service", "spvclient")))
	// Create a new server
	srv := server.NewServer(env.SuccinctPivKey, tbClient, logger)
	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
