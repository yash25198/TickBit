package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/crema-labs/sxg-go/internal/handler"
	"github.com/crema-labs/sxg-go/pkg/bitcoin"
	btcrelay "github.com/crema-labs/sxg-go/pkg/btc-relay"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"github.com/crema-labs/sxg-go/pkg/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Env struct {
	SuccinctPivKey string
	PrivKeyHex     string
	BtcIndexerUrl  string
	EthClientUrl   string
	DiscordWebhook string
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
		PrivKeyHex:     LookupEnv("PRIV_KEY"),
		BtcIndexerUrl:  LookupEnv("BTC_INDEXER_URL"),
		EthClientUrl:   LookupEnv("ETH_CLIENT_URL"),
		DiscordWebhook: LookupEnv("DISCORD_WEBHOOK"),
		TBAddress:      common.HexToAddress(LookupEnv("TB_ADDRESS")),
		SuccinctPivKey: LookupEnv("SP1_PRIVATE_KEY"),
	}
}

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
	env := Getenv()
	logger := GetAlertLogger(env.DiscordWebhook, "BTCRelay")
	privKey, err := crypto.HexToECDSA(env.PrivKeyHex)
	if err != nil {
		panic(err)
	}
	btcIndexer := bitcoin.NewElectrsIndexerClient(logger.With(zap.String("sub-service", "btcindexer")), env.BtcIndexerUrl, 10*time.Second)
	ethClient, err := ethereum.NewClient(logger, env.EthClientUrl)
	if err != nil {
		panic(err)
	}

	sp1handler := handler.HandleProofRequest{
		PrivKey: env.SuccinctPivKey,
		Logger:  logger.With(zap.String("sub-service", "handler")),
	}

	tbClient := ethereum.NewTBClient(ethClient, privKey, env.TBAddress, logger.With(zap.String("sub-service", "spvclient")))
	relay := btcrelay.NewBTCRelay(context.Background(), 2*time.Second, btcIndexer, tbClient, 0, &sp1handler, logger)
	relay.Start()
}
