package ethereum

import (
	"context"
	"crypto/ecdsa"

	"math/big"

	"github.com/crema-labs/sxg-go/pkg/ethereum/bindings/typings/ERC20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var (
	maxApproval = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
)

type EthClient interface {
	CallOpts() *bind.CallOpts
	GetTransactOpts(privKey *ecdsa.PrivateKey) (*bind.TransactOpts, error)
	GetCurrentBlock() (uint64, error)
	GetProvider() *ethclient.Client
	GetERC20Balance(tokenAddr common.Address, address common.Address) (*big.Int, error)
	GetDecimals(tokenAddr common.Address) (uint8, error)
	ApproveERC20(privKey *ecdsa.PrivateKey, amount *big.Int, tokenAddr common.Address, toAddr common.Address) (string, error)
	Allowance(tokenAddr common.Address, spender common.Address, owner common.Address) (*big.Int, error)
	TransferERC20(privKey *ecdsa.PrivateKey, amount *big.Int, tokenAddr common.Address, toAddr common.Address) (string, error)
	TransferEth(privKey *ecdsa.PrivateKey, amount *big.Int, toAddr common.Address) (string, error)
	WaitMined(ctx context.Context, tx *types.Transaction) (string, error)
	ChainID() *big.Int
}

type client struct {
	logger   *zap.Logger
	url      string
	provider *ethclient.Client
	chainID  *big.Int
}

func NewClient(logger *zap.Logger, url string) (EthClient, error) {
	provider, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	childLogger := logger.With(zap.String("service", "ethClient"))
	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &client{
		logger:   childLogger,
		url:      url,
		provider: provider,
		chainID:  chainID,
	}, nil
}

func (client *client) WaitMined(ctx context.Context, tx *types.Transaction) (string, error) {
	receipt, err := bind.WaitMined(ctx, client.provider, tx)
	if err != nil {
		return "", err
	}
	return receipt.TxHash.Hex(), nil
}

func (client *client) GetTransactOpts(privKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	transactor, err := bind.NewKeyedTransactorWithChainID(privKey, client.chainID)
	if err != nil {
		return nil, err
	}
	return transactor, nil
}

func (client *client) GetCurrentBlock() (uint64, error) {
	return client.provider.BlockNumber(context.Background())
}

func (client *client) GetProvider() *ethclient.Client {
	return client.provider
}

func (client *client) GetERC20Balance(tokenAddr common.Address, userAddr common.Address) (*big.Int, error) {
	instance, err := ERC20.NewERC20(tokenAddr, client.provider)
	if err != nil {
		return big.NewInt(0), err
	}
	return instance.BalanceOf(client.CallOpts(), userAddr)
}
func (client *client) GetDecimals(tokenAddr common.Address) (uint8, error) {
	instance, err := ERC20.NewERC20(tokenAddr, client.provider)
	if err != nil {
		return 0, err
	}
	return instance.Decimals(client.CallOpts())
}

func (client *client) ApproveERC20(privKey *ecdsa.PrivateKey, amount *big.Int, tokenAddr common.Address, toAddr common.Address) (string, error) {
	instance, err := ERC20.NewERC20(tokenAddr, client.provider)
	if err != nil {
		return "", err
	}
	transactor, err := client.GetTransactOpts(privKey)
	if err != nil {
		return "", err
	}
	tx, err := instance.Approve(transactor, toAddr, amount)
	if err != nil {
		return "", err
	}
	client.logger.Debug("approve erc20",
		zap.String("amount", amount.String()),
		zap.String("token address", tokenAddr.Hex()),
		zap.String("to address", toAddr.Hex()),
		zap.String("txHash", tx.Hash().Hex()))
	receipt, err := bind.WaitMined(context.Background(), client.provider, tx)
	if err != nil {
		return "", err
	}
	return receipt.TxHash.Hex(), nil
}
func (client *client) TransferERC20(privKey *ecdsa.PrivateKey, amount *big.Int, tokenAddr common.Address, toAddr common.Address) (string, error) {
	instance, err := ERC20.NewERC20(tokenAddr, client.provider)
	if err != nil {
		return "", err
	}
	transactor, err := client.GetTransactOpts(privKey)
	if err != nil {
		return "", err
	}
	tx, err := instance.Transfer(transactor, toAddr, amount)
	if err != nil {
		return "", err
	}
	client.logger.Debug("Transfer erc20",
		zap.String("amount", amount.String()),
		zap.String("token address", tokenAddr.Hex()),
		zap.String("to address", toAddr.Hex()),
		zap.String("txHash", tx.Hash().Hex()))
	receipt, err := bind.WaitMined(context.Background(), client.provider, tx)
	if err != nil {
		return "", err
	}
	return receipt.TxHash.Hex(), nil
}
func (client *client) TransferEth(privKey *ecdsa.PrivateKey, amount *big.Int, toAddr common.Address) (string, error) {
	fromAddress := crypto.PubkeyToAddress(privKey.PublicKey)
	nonce, err := client.provider.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	gasLimit := uint64(21000)
	gasPrice, err := client.provider.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(nonce, toAddr, amount, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(client.chainID), privKey)
	if err != nil {
		return "", err
	}
	err = client.provider.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	client.logger.Debug("Transfer eth",
		zap.String("amount", amount.String()),
		zap.String("to address", toAddr.Hex()),
		zap.String("txHash", signedTx.Hash().Hex()))
	receipt, err := bind.WaitMined(context.Background(), client.provider, signedTx)
	if err != nil {
		return "", err
	}
	return receipt.TxHash.Hex(), nil
}

func (client *client) Allowance(tokenAddr common.Address, spender common.Address, owner common.Address) (*big.Int, error) {
	instance, err := ERC20.NewERC20(tokenAddr, client.provider)
	if err != nil {
		return nil, err
	}
	return instance.Allowance(client.CallOpts(), owner, spender)
}

func (client *client) CallOpts() *bind.CallOpts {
	return &bind.CallOpts{
		Pending: true,
	}
}

func (client *client) ChainID() *big.Int {
	return client.chainID
}
