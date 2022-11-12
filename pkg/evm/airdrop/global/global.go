package global

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"strings"
)

var (
	DB           *gorm.DB
	EthClient    *ethclient.Client
	ChainId      int64
	GasPriceRate = uint64(100) // gasPrice浮动
)

var (
	accounts []Account
)

type Account struct {
	ContractAddress string // 合约地址
	Address         string // 地址
	PrivateKey      string // 私钥
}

// SetAccount 设置账号
func SetAccount(_accounts ...Account) {
	accounts = _accounts
}

func GetAccount(contractAddress string) *Account {
	for _, acc := range accounts {
		if strings.EqualFold(acc.ContractAddress, contractAddress) {
			return &acc
		}
	}
	return nil
}
