package main

import (
	"context"
	"os"
	"time"

	"github.com/crema-labs/sxg-go/cmd/utils"
	"github.com/crema-labs/sxg-go/internal/handler"
	"github.com/crema-labs/sxg-go/pkg/bitcoin"
	btcrelay "github.com/crema-labs/sxg-go/pkg/btc-relay"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"github.com/crema-labs/sxg-go/pkg/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/maps"
)

func GetAlertLogger(discordHookUrl, service string) *zap.Logger {
	logconfig := zap.NewProductionEncoderConfig()
	logconfig.EncodeTime = zapcore.ISO8601TimeEncoder // Customize time format if needed

	// Core for console output
	consoleEncoder := zapcore.NewConsoleEncoder(logconfig)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)

	// Core for Discord, only for ErrorLevel and above
	discordHook := logger.NewDiscordHook(discordHookUrl)
	discordEncoder := zapcore.NewJSONEncoder(logconfig)
	discordCore := zapcore.NewCore(discordEncoder, zapcore.AddSync(discordHook), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel // Only log errors and above
	}))

	// Combine cores
	logger := zap.New(zapcore.NewTee(consoleCore, discordCore)).With(zap.String("service", service))
	defer logger.Sync() // Flushes buffer, if any

	return logger
}

func main() {
	env := utils.Getenv()
	os.Setenv("RUST_LOG", "info")
	os.Setenv("SP1_PROVER", "network")
	logger := GetAlertLogger(env.DiscordWebhook, "BTCRelay")
	privKey, err := crypto.HexToECDSA(env.PrivKeyHex)
	if err != nil {
		panic(err)
	}
	btcIndexer := bitcoin.NewElectrsIndexerClient(logger.With(zap.String("sub-service", "btcindexer")), env.BtcIndexerUrl, 10*time.Second)
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

	sp1handler := handler.HandleProofRequest{
		PrivKey: env.SuccinctPivKey,
		Logger:  logger.With(zap.String("sub-service", "handler")),
	}

	tbClient := ethereum.NewTBClient(cp, privKey, logger.With(zap.String("sub-service", "spvclient")))
	relay := btcrelay.NewBTCRelay(context.Background(), 2*time.Second, btcIndexer, tbClient, 0, env.DefaultChain, maps.Keys(cp), &sp1handler, logger)
	relay.Start()
}
