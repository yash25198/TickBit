package btcrelay_test

// import (
// 	"context"
// 	"time"

// 	btcrelay "github.com/crema-labs/sxg-go/btc-relay"
// 	"github.com/crema-labs/sxg-go/pkg/bitcoin"
// 	"github.com/crema-labs/sxg-go/pkg/ethereum"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	. "github.com/onsi/ginkgo/v2"

// 	// . "github.com/onsi/gomega"
// 	"go.uber.org/zap"
// )

// var _ = Describe("BtcRelay", func() {
// 	Context("BtcRelay testing", func() {
// 		It("should work ___-___", func() {
// 			logger, err := zap.NewDevelopment()
// 			if err != nil {
// 				panic(err)
// 			}
// 			privKey1, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
// 			if err != nil {
// 				panic(err)
// 			}
// 			btcIndexer := bitcoin.NewElectrsIndexerClient(logger, "http://13.251.31.137:3000", 10*time.Second)
// 			ethClient, err := ethereum.NewClient(logger, "https://rpc.tenderly.co/fork/53d9b704-a4cd-465c-9fe9-a6fb665d52d2")
// 			if err != nil {
// 				panic(err)
// 			}
// 			// header, err := bitcoin.ParseHeader("00e00020c737bd4d5f94416ff1555230fd88e2622e210e71096c040000000000000000001b5a3ccb475286e6b33a986e75574352dc0f83bfeccd9576e0eb466a764c13f2a4283d65a99c04179685cd73")
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// topts, err := ethClient.GetTransactOpts(privKey1)
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// _, _, _, err = VerifySPV.DeployVerifySPV(topts, ethClient.GetProvider(), header)
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// buf := make([]byte, 32)
// 			// // then we can call rand.Read.
// 			// _, err = rand.Read(buf)
// 			// s, err := verifyContract.BlockHeaders(ethClient.CallOpts(), [32]byte(buf))
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// fmt.Println(s)
// 			spvClient := ethereum.NewSPVClient(ethClient, privKey1, common.HexToAddress("0xA172158Bc63C8037f5eA9f6373f18d2d42A8B9b4"), logger)
// 			relay := btcrelay.NewBTCRelay(context.Background(), 812448, 72, time.Minute, btcIndexer, spvClient, logger)
// 			relay.Start()
// 			// fmt.Println(relay)
// 		})
// 	})
// })
