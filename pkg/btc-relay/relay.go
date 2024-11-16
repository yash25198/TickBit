package btcrelay

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"time"

	"github.com/crema-labs/sxg-go/internal/handler"
	"github.com/crema-labs/sxg-go/pkg/bitcoin"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"go.uber.org/zap"
)

type btcRelay struct {
	ctx                 context.Context
	lastRegisteredBlock uint64
	pollInterval        time.Duration
	btcIndexer          bitcoin.IndexerClient
	tbClient            ethereum.TBClient
	sp1Handler          *handler.HandleProofRequest
	defaultChain        string
	chainList           []string
	logger              *zap.Logger
}

type BTCRelay interface {
	Start()
}

func NewBTCRelay(ctx context.Context,
	pollInterval time.Duration,
	btcIndexer bitcoin.IndexerClient,
	tbClient ethereum.TBClient,
	lastRegisteredBlock uint64,
	defaultChain string,
	chainList []string,
	sp1Handler *handler.HandleProofRequest,
	logger *zap.Logger) BTCRelay {

	return &btcRelay{
		ctx:                 ctx,
		pollInterval:        pollInterval,
		lastRegisteredBlock: lastRegisteredBlock,
		btcIndexer:          btcIndexer,
		tbClient:            tbClient,
		defaultChain:        defaultChain,
		chainList:           chainList,
		sp1Handler:          sp1Handler,
		logger:              logger,
	}

}

func (b *btcRelay) Start() {
	var err error
	if b.lastRegisteredBlock, err = b.tbClient.LastResgisteredBlock(b.defaultChain); err != nil {
		b.logger.Error("error getting last registered block", zap.Error(err))
		return
	}
	newBlockChan, err := b.btcIndexer.SubscribeToLatestBlocks(context.Background(), b.pollInterval)
	if err != nil {
		panic(err)
	}
	for {

		select {
		case newBlock := <-newBlockChan:
			if newBlock.BlockNumber > b.lastRegisteredBlock {
				b.lastRegisteredBlock = newBlock.BlockNumber
				go func() {
					b.logger.Info("new block found", zap.Uint64("blockNumber", newBlock.BlockNumber))

					var reqId string
					// add label
				proofRequest:
					for {
						reqId, err = b.sp1Handler.GenerateProofRequest(handler.ProofRequest{
							BlockNumber: newBlock.BlockNumber,
							BlockHash:   newBlock.BlockHash,
						})
						switch err {
						case nil:
							break proofRequest
						case handler.ErrDatNotFound:
							b.logger.Warn("data not found", zap.Uint64("blockNumber", newBlock.BlockNumber))
							time.Sleep(2 * time.Second)
							continue
						default:
							if strings.Contains(err.Error(), "proof request already exists, try getting status") {
								b.logger.Warn("proof request already exists", zap.Uint64("blockNumber", newBlock.BlockNumber))
								break proofRequest
							}
							b.logger.Error("error generating proof request", zap.Error(err))
							return
						}
					}

					b.logger.Info("proof request generated", zap.Uint64("blockNumber", newBlock.BlockNumber), zap.String("reqId", reqId))
					var proofResponse handler.StatusResponse
					for {
						proofResponse, err = b.sp1Handler.ProofStatus(newBlock.BlockNumber)
						if err != nil {
							b.logger.Warn("failed to check status")
						}

						if proofResponse.Status == handler.Ready && proofResponse.Proof != nil {
							b.logger.Info("proof ready", zap.Uint64("blockNumber", newBlock.BlockNumber))
							break
						}
						time.Sleep(2 * time.Second)
					}

					proof, err := hex.DecodeString(proofResponse.Proof.ProofBytes[2:])
					if err != nil {
						b.logger.Error("error encoding proof", zap.Error(err))
					}
					for _, chain := range b.chainList {
						err = b.tbClient.VerifyBlock(context.Background(), chain, big.NewInt(int64(newBlock.BlockNumber)), newBlock.Header, proof)
						if err != nil {
							b.logger.Error("error verifying block", zap.Error(err), zap.String("chain", chain))
						}
					}
				}()

			} else {
				b.logger.Info("block already registered 1", zap.Uint64("blockNumber", newBlock.BlockNumber))
			}
		case <-time.After(b.pollInterval):
			b.logger.Info("looking for new blocks")
		}
	}
}
