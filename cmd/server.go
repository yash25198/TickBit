package main

import (
	"log"
	"os"

	"github.com/crema-labs/sxg-go/internal/server"
	"go.uber.org/zap"
)

func main() {
	os.Setenv("RUST_LOG", "info")
	os.Setenv("SP1_PROVER", "network")
	priv_key, ok := os.LookupEnv("SP1_PRIVATE_KEY")
	if !ok {
		log.Fatal("SP1_PRIVATE_KEY environment variable is required")
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to create logger")
	}
	// Create a new server
	srv := server.NewServer(priv_key, logger)
	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
