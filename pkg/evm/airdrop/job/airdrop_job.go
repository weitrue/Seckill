package job

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/362228416/go-eth/cs/eth2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/weitrue/Seckill/pkg/evm/airdrop/global"
	"github.com/weitrue/Seckill/pkg/evm/airdrop/model"
)

func validAddress(address string) bool {
	return len(address) == 42 && strings.Index(address, "0x") == 0
}

func getOpts(priv string) *bind.TransactOpts {
	opts2, _ := bind.NewKeyedTransactorWithChainID(eth2.HexToECSSA(priv), big.NewInt(global.ChainId))
	return opts2
}

func RunAirdropJob() {
	for {
		var err error
		var airdrop model.Airdrop
		err = global.DB.Debug().Where("state=?", 0).First(&airdrop).Error
		if err != nil {
			log.Println(err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		log.Println("待空投记录", airdrop)
		if !validAddress(airdrop.ContractAddress) {
			log.Println("合约地址无效", airdrop.ContractAddress)
			airdrop.State = 4
			err = global.DB.Debug().Model(&airdrop).Update("state", airdrop.State).Error
			continue
		}
		if !validAddress(airdrop.ToAddress) {
			log.Println("空头地址无效", airdrop.ToAddress)
			airdrop.State = 4
			err = global.DB.Debug().Model(&airdrop).Update("state", airdrop.State).Error
			continue
		}

		account := global.GetAccount(airdrop.ContractAddress)
		if account == nil {
			log.Println("找不到合约地址的私钥", airdrop.ToAddress)
			airdrop.State = 4
			err = global.DB.Debug().Model(&airdrop).Update("state", airdrop.State).Error
			continue
		}

		// 获取钱包地址余额
		ethBalance, err := global.EthClient.BalanceAt(context.Background(), common.HexToAddress(account.Address), nil)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if ethBalance.Cmp(new(big.Int).SetInt64(11e15)) < 0 {
			log.Println("手续费不足，暂时不处理", account.Address)
			continue
		}

		// 创建已部署的合约地址实例
		contract, err := NewNftContract(common.HexToAddress(airdrop.ContractAddress), global.EthClient)
		tokenId, _ := new(big.Int).SetString(airdrop.TokenId, 10)
		owner, err := contract.OwnerOf(nil, tokenId)
		if err != nil {
			log.Println("查询owner失败", err.Error())
			continue
		}
		if !strings.EqualFold(owner.Hex(), account.Address) {
			log.Println("当前地址不持有此nft")
			airdrop.State = 4
			err = global.DB.Debug().Model(&airdrop).Update("state", airdrop.State).Error
			continue
		}

		var lastAirdrop model.Airdrop
		err = global.DB.Debug().Where("from_address=?", account.Address).Order("id desc").First(&lastAirdrop).Error
		if err != nil && err.Error() != "record not found" {
			log.Println(err)
			continue
		}
		nonce, _ := global.EthClient.NonceAt(context.Background(), common.HexToAddress(account.Address), nil)
		if lastAirdrop.ID > 0 {
			if lastAirdrop.Nonce >= nonce {
				log.Println("交易未确认，等待交易确认", "上笔Nonce", lastAirdrop.Nonce, "当前Nonce", nonce)
				time.Sleep(2 * time.Second)
				continue
			}
		}

		opts := getOpts(account.PrivateKey)
		opts.Nonce = new(big.Int).SetUint64(nonce)
		gasPrice, err := global.EthClient.SuggestGasPrice(context.Background())
		if err != nil {
			log.Println("获取gasPrice失败", err)
			continue
		}
		// gasPrice 浮动 = gasPrice * GasPriceRate / 100
		if global.GasPriceRate > 0 {
			gasPrice = gasPrice.Mul(gasPrice, new(big.Int).SetUint64(global.GasPriceRate))
			gasPrice = gasPrice.Div(gasPrice, new(big.Int).SetUint64(100))
		}
		opts.GasPrice = gasPrice
		airdrop.State = 1
		err = global.DB.Debug().Model(&airdrop).Update("state", airdrop.State).Error
		if err != nil {
			log.Println("锁定记录失败", err)
			continue
		}
		tx, err := contract.SafeTransferFrom(opts, common.HexToAddress(account.Address), common.HexToAddress(airdrop.ToAddress), tokenId)
		if err != nil {
			log.Println(err)
			airdrop.State = 0
			global.DB.Debug().Model(&airdrop).Update("state", airdrop.State)
			continue
		}
		// 交易转为raw
		w := new(bytes.Buffer)
		err = tx.EncodeRLP(w)
		if err == nil {
			raw := hex.EncodeToString(w.Bytes())
			airdrop.Raw = raw
		}
		airdrop.State = 2
		airdrop.Txid = tx.Hash().Hex()
		airdrop.Nonce = nonce
		airdrop.FromAddress = account.Address
		airdrop.Version++
		err = global.DB.Debug().Model(&airdrop).Where("id=? and version=?", airdrop.ID, airdrop.Version-1).Updates(airdrop).Error
		if err != nil {
			log.Println(err)
		}

		err = waitTransactionConfirm(tx.Hash())
		if err != nil {
			log.Println(err)
			continue
		}

	}
}

func waitTransactionConfirm(txid common.Hash) error {
	for i := 0; i < 100; i++ {
		log.Println("TransactionByHash", txid)
		receipt, err := global.EthClient.TransactionReceipt(context.Background(), txid)
		if err != nil || receipt.BlockNumber == nil {
			log.Println("交易未确认", err)
			time.Sleep(5 * time.Second)
		} else {
			log.Println("交易已确认", receipt)
			return nil
		}
	}
	return errors.New("获取交易失败")
}
