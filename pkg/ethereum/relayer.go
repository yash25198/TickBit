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
	ethClient    EthClient
	signer       *ecdsa.PrivateKey
	tickbitAddr  common.Address
	logger       *zap.Logger
	tickContract TickBit.TickBit
}

type TBClient interface {
	LastResgisteredBlock() uint64
	VerifyBlock(ctx context.Context, blockNumber *big.Int, header []byte, proof []byte) error
	IsRegistered(ctx context.Context, blockNumber *big.Int) (bool, error)
	GetBets(ctx context.Context, bn *big.Int) ([]TickBit.TickBitBetPlaced, error)
	GetPoolValue(ctx context.Context, bn *big.Int) (Pool, error)
}

type Pool struct {
	AcruedAmount *big.Int
	SettledAt    *big.Int
}

func NewTBClient(ethClient EthClient, signer *ecdsa.PrivateKey, addr common.Address, logger *zap.Logger) TBClient {
	tickContract, err := TickBit.NewTickBit(addr, ethClient.GetProvider())
	if err != nil {
		panic(err)
	}
	return &tickbitClient{
		ethClient:    ethClient,
		signer:       signer,
		tickbitAddr:  addr,
		tickContract: *tickContract,
		logger:       logger,
	}
}

func (c *tickbitClient) LastResgisteredBlock() uint64 {
	latestBlock, err := c.tickContract.LatestBlock(c.ethClient.CallOpts())
	if err != nil {
		c.logger.Error(err.Error())
		return 0
	}
	return latestBlock.Uint64()
}
func (c *tickbitClient) GetAddress() common.Address {
	return c.tickbitAddr
}
func (c *tickbitClient) VerifyBlock(ctx context.Context, blockNumber *big.Int, header []byte, proof []byte) error {
	isActive, err := c.IsRegistered(ctx, blockNumber)
	if err != nil {
		return err
	}
	if isActive {
		c.logger.Info("block already registered", zap.Uint64("blockNumber", blockNumber.Uint64()))
		return nil
	}

	opts, err := c.ethClient.GetTransactOpts(c.signer)
	if err != nil {
		return err
	}
	tx, err := c.tickContract.VerifyAndSettleBlock(opts, blockNumber, header, proof)
	if err != nil {
		return err
	}
	txHash, err := c.ethClient.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	c.logger.Info("Block verified", zap.String("txHash", txHash), zap.Uint64("blockNumber", blockNumber.Uint64()))
	return nil
}

func (c *tickbitClient) IsRegistered(ctx context.Context, blockNumber *big.Int) (bool, error) {
	header, err := c.tickContract.VerifiedBlocks(c.ethClient.CallOpts(), blockNumber)
	if err != nil {
		return false, err
	}

	fmt.Println(len(header.MerkleRootHash), header.MerkleRootHash, header)
	return false, err
}

func (c *tickbitClient) GetBets(ctx context.Context, bn *big.Int) ([]TickBit.TickBitBetPlaced, error) {
	evt, err := c.tickContract.FilterBetPlaced(nil, nil, []*big.Int{bn})
	if err != nil {
		return nil, err
	}

	var evnts []TickBit.TickBitBetPlaced

	for evt.Next() {
		evnts = append(evnts, *evt.Event)
	}

	return evnts, nil
}

func (c *tickbitClient) GetPoolValue(ctx context.Context, bn *big.Int) (Pool, error) {
	return c.tickContract.Pools(c.ethClient.CallOpts(), bn)
}
