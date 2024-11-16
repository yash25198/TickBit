package utils

import (
	"encoding/json"
	"os"
)

type ChainParams struct {
	TBAddress    string `json:"TB_ADDRESS"`
	EthClientUrl string `json:"ETH_CLIENT_URL"`
}

type Env struct {
	SuccinctPivKey string                 `json:"SP1_PRIVATE_KEY"`
	PrivKeyHex     string                 `json:"PRIV_KEY_HEX"`
	BtcIndexerUrl  string                 `json:"BTC_INDEXER_URL"`
	DiscordWebhook string                 `json:"DISCORD_WEBHOOK"`
	DefaultChain   string                 `json:"DEFAULT_CHAIN"`
	ChainParams    map[string]ChainParams `json:"CHAIN_PARAMS"`
}

func Getenv() Env {
	//read config.json
	config, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	env := Env{}
	//parse config.json
	if err := json.Unmarshal(config, &env); err != nil {
		panic(err)
	}
	return env
}
