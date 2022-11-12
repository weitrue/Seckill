package airdrop

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/weitrue/Seckill/pkg/evm/airdrop/global"
	"github.com/weitrue/Seckill/pkg/evm/airdrop/model"
	"gorm.io/gorm"
)

// SetDatabase 设置数据库
func SetDatabase(_db *gorm.DB) {
	global.DB = _db
}

// SetGasPriceRate 设置gasPrice百分比
func SetGasPriceRate(rate uint64) {
	global.GasPriceRate = rate
}

// SetEthClient 设置ETH客户端
func SetEthClient(client *ethclient.Client) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	global.ChainId = chainId.Int64()
	global.EthClient = client
}

// AddAirdrop 空投
func AddAirdrop(m *model.Airdrop) error {
	return global.DB.Create(m).Error
}

// SetAccount 设置账号
func SetAccount(_accounts ...global.Account) {
	global.SetAccount(_accounts...)
}
