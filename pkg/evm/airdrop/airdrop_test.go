package airdrop

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/weitrue/Seckill/pkg/evm/airdrop/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"

	"github.com/weitrue/Seckill/pkg/evm/airdrop/job"
	"github.com/weitrue/Seckill/pkg/evm/airdrop/model"
)

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123123@tcp(127.0.0.1:3306)/chain_ethw?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                // default size for string fields
		DisableDatetimePrecision:  true,                                                                               // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                               // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                               // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                              // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)
	db.AutoMigrate(model.Airdrop{})

	SetDatabase(db)

	//client, err := ethclient.Dial("https://mainnet.ethereumpow.org")		// ethw mainnet
	client, err := ethclient.Dial("http://localhost:8545") // heco testnet
	if err != nil {
		panic(err)
	}
	SetEthClient(client)
	SetGasPriceRate(120) // 提高20%的gasPrice

	SetAccount(global.Account{
		ContractAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
		Address:         "0x35253468619BcBB9585CC09aec5e5f7820914Fca",
		PrivateKey:      "91e31d94755f15ee1a30f80e2227105d4f43d27e984e162e9e63f8f4b2516837",
	})

}

func TestInit(t *testing.T) {

}

func TestAddAirdrop(t *testing.T) {
	err := AddAirdrop(&model.Airdrop{
		ContractAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
		TokenId:         "0",
		ToAddress:       "0x9EFBcc05704BC1F3BE3cD40EC066623Bb886Cc46",
	})
	if err != nil {
		panic(err)
	}
}

func TestRunAirdropJob(t *testing.T) {
	TestAddAirdrop(t)
	job.RunAirdropJob()
}
