package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/weitrue/Seckill/pkg/evm/contract/snaili"
	"math/big"
)

type Config struct {
	Endpoint string `toml:"endpoint" json:"endpoint"`
	Contract string `toml:"contract" json:"contract"`
}

type Erc interface {
	// BalanceAt 账户的余额
	BalanceAt(ctx context.Context, account string) (*big.Int, error)
	// PendingBalanceAt 待处理的余额
	PendingBalanceAt(ctx context.Context, account string) (*big.Int, error)
	// BalanceAtBlock 某区块时的账户余额
	BalanceAtBlock(ctx context.Context, account string, blockNumber int64) (*big.Int, error)
	// GenerateKey 生成公私钥
	GenerateKey() (string, string, error)
	Contract
}

type Contract interface {
	GetOwner(tokenId string) (string, error)
}

func NewErc(c *Config) (Erc, error) {
	if c == nil || len(c.Endpoint) == 0 {
		return nil, errors.New("err config")
	}

	var contract Contract
	var err error
	switch c.Contract {
	case snaili.Contract:
		contract, err = snaili.NewContractSnail(c.Endpoint)
		if err != nil {
			return nil, err
		}
	default:
		contract = nil
	}

	return NewClient(c.Endpoint, contract)
}

type defaultClient struct {
	c        *ethclient.Client
	contract Contract
}

func (c *defaultClient) GetOwner(tokenId string) (string, error) {
	if c.contract == nil {
		return "", errors.New("unsupported")
	}

	return c.contract.GetOwner(tokenId)
}

func NewClient(endpoint string, contract Contract) (*defaultClient, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	return &defaultClient{
		c:        client,
		contract: contract,
	}, nil
}

func (c *defaultClient) BalanceAt(ctx context.Context, account string) (*big.Int, error) {
	balanceAt, err := c.c.BalanceAt(ctx, common.HexToAddress(account), nil)
	if err != nil {
		return nil, err
	}

	return balanceAt, nil
}

func (c *defaultClient) PendingBalanceAt(ctx context.Context, account string) (*big.Int, error) {
	balanceAt, err := c.c.PendingBalanceAt(ctx, common.HexToAddress(account))
	if err != nil {
		return nil, err
	}

	return balanceAt, nil
}

func (c *defaultClient) BalanceAtBlock(ctx context.Context, account string, blockNumber int64) (*big.Int, error) {
	balanceAt, err := c.c.BalanceAt(ctx, common.HexToAddress(account), big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}

	return balanceAt, nil
}

func (c *defaultClient) GenerateKey() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey) // FromECDSA方法将其转换为字节
	private := hexutil.Encode(privateKeyBytes[2:])  // 转换为十六进制字符串并删除"0x",即为私钥

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("failed on assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), private, err
}
