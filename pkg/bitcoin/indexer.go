package bitcoin

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/crema-labs/sxg-go/pkg/types"
	"github.com/crema-labs/sxg-go/pkg/utils"
	"go.uber.org/zap"
)

const (
	DefaultElectrsIndexerURL = "http://0.0.0.0:30000"

	DefaultRetryInterval = 5 * time.Second
)

type IndexerClient interface {
	GetTipBlockHeight(ctx context.Context) (uint64, error)

	GetTx(ctx context.Context, txid string) (types.Transaction, error)

	GetBlockHash(ctx context.Context, blockNumber uint64) (string, error)

	GetBlockNumber(ctx context.Context, blockHash string) (uint64, error)

	GetBlockHeader(ctx context.Context, blockHash string) ([]byte, error)

	SubscribeToLatestBlocks(ctx context.Context, poolTimer time.Duration) (<-chan types.BlockDetails, error)

	GetBlockOfTransaction(ctx context.Context, txHash string) (uint64, string, error)

	GetMerkleProofs(ctx context.Context, txHash string) (types.ProofResponse, error)
}

type electrsIndexerClient struct {
	logger        *zap.Logger
	url           string
	retryInterval time.Duration
}

func NewElectrsIndexerClient(logger *zap.Logger, url string, retryInterval time.Duration) IndexerClient {
	return &electrsIndexerClient{
		logger:        logger,
		url:           url,
		retryInterval: retryInterval,
	}
}

func (client *electrsIndexerClient) GetBlockHash(ctx context.Context, blockNumber uint64) (string, error) {
	endpoint, err := url.JoinPath(client.url, "block-height", strconv.FormatUint(blockNumber, 10))
	if err != nil {
		return "", err
	}
	var blockhash string
	if err := utils.Retry(client.logger, ctx, client.retryInterval, func() error {
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Decode response
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		blockhash = string(data)
		return nil
	}); err != nil {
		return "", err
	}

	return blockhash, nil
}
func (client *electrsIndexerClient) GetBlockNumber(ctx context.Context, blockHash string) (uint64, error) {
	endpoint, err := url.JoinPath(client.url, "block", blockHash, "status")
	if err != nil {
		return 0, err
	}

	// Send the request
	var status struct {
		Height uint64 `json:"height"`
	}
	if err := utils.Retry(client.logger, ctx, client.retryInterval, func() error {
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
			return fmt.Errorf("failed to decode status response: %w", err)
		}
		return err
	}); err != nil {
		return 0, err
	}

	return status.Height, nil
}
func (client *electrsIndexerClient) SubscribeToLatestBlocks(ctx context.Context, pollTimer time.Duration) (<-chan types.BlockDetails, error) {
	blockChan := make(chan types.BlockDetails)
	blockNumber, err := client.GetTipBlockHeight(ctx)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(blockChan)
		for {
			time.Sleep(pollTimer)
			latest, err := client.GetTipBlockHeight(ctx)
			if err != nil {
				break
			}
			if blockNumber > latest {
				continue
			}
			for i := blockNumber; i <= latest; i++ {
				blockHash, err := client.GetBlockHash(ctx, i)
				if err != nil {
					break
				}
				header, err := client.GetBlockHeader(ctx, blockHash)
				if err != nil {
					break
				}
				blockChan <- types.BlockDetails{
					BlockHash:   blockHash,
					Header:      header,
					BlockNumber: i,
				}
			}
			blockNumber = latest + 1

		}
	}()
	return blockChan, nil

}
func (client *electrsIndexerClient) GetTipBlockHeight(ctx context.Context) (uint64, error) {
	endpoint, err := url.JoinPath(client.url, "blocks", "tip", "height")
	if err != nil {
		return 0, err
	}

	// Send the request
	var height uint64
	if err := utils.Retry(client.logger, ctx, client.retryInterval, func() error {
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Decode response
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		height, err = strconv.ParseUint(string(data), 10, 64)
		return err
	}); err != nil {
		return 0, err
	}

	return height, nil
}
func (client *electrsIndexerClient) GetMerkleProofs(ctx context.Context, txHash string) (types.ProofResponse, error) {
	endpoint, err := url.JoinPath(client.url, "tx", txHash, "merkle-proof")
	if err != nil {
		return types.ProofResponse{}, err
	}

	// Send the request
	var proofs types.ProofResponse
	if err := utils.Retry(client.logger, ctx, client.retryInterval, func() error {
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&proofs); err != nil {
			return fmt.Errorf("failed to decode status response: %w", err)
		}
		return err
	}); err != nil {
		return types.ProofResponse{}, err
	}

	return proofs, nil
}
func (client *electrsIndexerClient) GetBlockHeader(ctx context.Context, blockHash string) ([]byte, error) {
	endpoint, err := url.JoinPath(client.url, "block", blockHash, "header")
	if err != nil {
		return nil, err
	}

	var blockHeaderRaw string
	if err := utils.Retry(client.logger, ctx, client.retryInterval, func() error {
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Decode response
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		blockHeaderRaw = string(data)
		return nil
	}); err != nil {
		return nil, err
	}

	return hex.DecodeString(blockHeaderRaw)
}
func (client *electrsIndexerClient) GetBlockOfTransaction(ctx context.Context, txHash string) (uint64, string, error) {
	endpoint, err := url.JoinPath(client.url, "tx", txHash, "status")
	if err != nil {
		return 0, "", err
	}

	// Send the request
	var status struct {
		Confirmed   bool   `json:"confirmed"`
		BLockHeight uint64 `json:"block_height"`
		BlockHash   string `json:"block_hash"`
	}
	if err := utils.Retry(client.logger, ctx, client.retryInterval, func() error {
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
			return fmt.Errorf("failed to decode status response: %w", err)
		}
		return err
	}); err != nil {
		return 0, "", err
	}

	if status.Confirmed {
		return status.BLockHeight, status.BlockHash, nil
	}
	return 0, "", nil
}

func (client *electrsIndexerClient) GetTx(ctx context.Context, txid string) (types.Transaction, error) {
	endpoint, err := url.JoinPath(client.url, "tx", txid)
	if err != nil {
		return types.Transaction{}, err
	}

	// Send the request
	var tx types.Transaction
	if err := utils.Retry(client.logger, ctx, client.retryInterval, func() error {
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Decode response
		if err := json.NewDecoder(resp.Body).Decode(&tx); err != nil {
			return fmt.Errorf("failed to decode UTXOs: %w", err)
		}
		return nil
	}); err != nil {
		return types.Transaction{}, err
	}

	return tx, nil
}
