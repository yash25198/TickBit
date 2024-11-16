package ethereum

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"fmt"

	"github.com/crema-labs/sxg-go/pkg/ethereum/bindings/typings/TickBit"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const MaxQueryBlockRange = 1000

type tickbitClient struct {
	signer        *ecdsa.PrivateKey
	logger        *zap.Logger
	tickContracts map[string]TickBit.TickBit
	chainParams   map[string]ChainParams
}

type TBClient interface {
	LastResgisteredBlock(chain string) (uint64, error)
	VerifyBlock(ctx context.Context, chain string, blockNumber *big.Int, header []byte, proof []byte) error
	IsRegistered(ctx context.Context, chain string, blockNumber *big.Int) (bool, error)
	GetBets(ctx context.Context, chain string, bn *big.Int) ([]TickBit.TickBitBetPlaced, error)
	GetPoolValue(ctx context.Context, chain string, bn *big.Int) (Pool, error)
}

type Pool struct {
	AcruedAmount *big.Int
	SettledAt    *big.Int
}

type ChainParams struct {
	TBAddress common.Address
	EthClient EthClient
}

func NewTBClient(cp map[string]ChainParams, signer *ecdsa.PrivateKey, logger *zap.Logger) TBClient {

	tickContracts := make(map[string]TickBit.TickBit)
	for k, v := range cp {
		tickContract, err := TickBit.NewTickBit(v.TBAddress, v.EthClient.GetProvider())
		if err != nil {
			panic(err)
		}
		tickContracts[k] = *tickContract
	}
	return &tickbitClient{
		signer:        signer,
		tickContracts: tickContracts,
		logger:        logger,
		chainParams:   cp,
	}
}

func (c *tickbitClient) LastResgisteredBlock(chain string) (uint64, error) {

	ethClient, contract, err := c.getChainParams(chain)
	if err != nil {
		return 0, err
	}

	latestBlock, err := contract.LatestBlock(ethClient.CallOpts())
	if err != nil {
		c.logger.Error("error getting latest block", zap.Error(err))
		return 0, err
	}
	return latestBlock.Uint64(), nil
}
func (c *tickbitClient) GetAddress(chain string) common.Address {
	return c.chainParams[chain].TBAddress
}
func (c *tickbitClient) VerifyBlock(ctx context.Context, chain string, blockNumber *big.Int, header []byte, proof []byte) error {
	isActive, err := c.IsRegistered(ctx, chain, blockNumber)
	if err != nil {
		return err
	}
	if isActive {
		c.logger.Info("block already registered", zap.Uint64("blockNumber", blockNumber.Uint64()))
		return nil
	}

	ethClient, contract, err := c.getChainParams(chain)
	if err != nil {
		return err
	}

	opts, err := ethClient.GetTransactOpts(c.signer)
	if err != nil {
		return err
	}

	tx, err := contract.VerifyAndSettleBlock(opts, blockNumber, header, proof)
	if err != nil {
		return err
	}
	txHash, err := ethClient.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	c.logger.Info("Block verified", zap.String("txHash", txHash), zap.Uint64("blockNumber", blockNumber.Uint64()))
	return nil
}

func (c *tickbitClient) IsRegistered(ctx context.Context, chain string, blockNumber *big.Int) (bool, error) {

	ethClient, contract, err := c.getChainParams(chain)
	if err != nil {
		return false, err
	}

	header, err := contract.VerifiedBlocks(ethClient.CallOpts(), blockNumber)
	if err != nil {
		return false, err
	}

	fmt.Println(len(header.MerkleRootHash), header.MerkleRootHash, header)
	return false, err
}

func (c *tickbitClient) GetBets(ctx context.Context, chain string, bn *big.Int) ([]TickBit.TickBitBetPlaced, error) {
	contract, ok := c.tickContracts[chain]
	if !ok {
		return nil, fmt.Errorf("GetBets : contract not found")
	}

	evt, err := contract.FilterBetPlaced(nil, nil, []*big.Int{bn})
	if err != nil {
		return nil, err
	}

	var evnts []TickBit.TickBitBetPlaced

	for evt.Next() {
		evnts = append(evnts, *evt.Event)
	}

	return evnts, nil
}

func (c *tickbitClient) GetPoolValue(ctx context.Context, chain string, bn *big.Int) (Pool, error) {
	ethClient, contract, err := c.getChainParams(chain)
	if err != nil {
		return Pool{}, err
	}
	return contract.Pools(ethClient.CallOpts(), bn)
}

func (c *tickbitClient) getChainParams(chain string) (EthClient, TickBit.TickBit, error) {
	cp, ok := c.chainParams[chain]
	if !ok {
		return nil, TickBit.TickBit{}, fmt.Errorf("getChainParams : chain not found")
	}

	return cp.EthClient, c.tickContracts[chain], nil

}
