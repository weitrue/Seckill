package model

import (
	"gorm.io/gorm"
)

type Airdrop struct {
	gorm.Model
	ContractAddress string `gorm:"comment:合约地址"`
	TokenId         string `gorm:"comment:TokenId"`
	ToAddress       string `gorm:"comment:接收空投地址"`
	State           int    `gorm:"comment:状态：0-待转出 1-正在转出 2-已转出 3-转出失败 4-跳过转出"`
	Txid            string `gorm:"comment:交易ID"`
	FromAddress     string `gorm:"comment:转出地址"`
	Nonce           uint64 `gorm:"comment:Nonce"`
	Raw             string `gorm:"type:text;comment:原始交易"`
	Version         int
}

func (a *Airdrop) TableName() string {
	return "e_airdrop"
}
