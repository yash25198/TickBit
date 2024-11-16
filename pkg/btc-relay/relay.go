package btcrelay

import (
	"context"
	"encoding/hex"
	"math/big"
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
	sp1Handler *handler.HandleProofRequest,
	logger *zap.Logger) BTCRelay {

	return &btcRelay{
		ctx:                 ctx,
		pollInterval:        pollInterval,
		lastRegisteredBlock: lastRegisteredBlock,
		btcIndexer:          btcIndexer,
		tbClient:            tbClient,
		sp1Handler:          sp1Handler,
		logger:              logger,
	}

}

func (b *btcRelay) Start() {
	b.lastRegisteredBlock = b.tbClient.LastResgisteredBlock()
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
					err = b.tbClient.VerifyBlock(context.Background(), big.NewInt(int64(newBlock.BlockNumber)), newBlock.Header, proof)
					if err != nil {
						b.logger.Error("error verifying block", zap.Error(err))
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
