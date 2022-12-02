package snaili

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	Contract = "0xAeef2b7D89b609b2890e10413D13177fC9301Ee6"
)

type ContractSnail struct {
	Client       *ethclient.Client
	endpoint     string
	SnailAddress string
}

func NewContractSnail(endpoint string) (*ContractSnail, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	return &ContractSnail{
		Client:       client,
		endpoint:     endpoint,
		SnailAddress: Contract,
	}, nil
}

func (c *ContractSnail) GetOwner(tokenId string) (string, error) {
	snail, err := NewSnaili(common.HexToAddress(c.SnailAddress), c.Client)
	if err != nil {
		return "", err
	}

	tokenIdStr, _ := new(big.Int).SetString(tokenId, 10)
	ownerOf, err := snail.OwnerOf(nil, tokenIdStr)
	if err != nil {
		return "", err
	}
	return ownerOf.Hex(), nil
}

func (c *ContractSnail) PublicMint(tokenId string) (string, error) {
	//bind.NewKeyedTransactorWithChainID()
	// TODO implement me
	panic("implement me")
}

func (c *ContractSnail) TokenUrl(tokenId string) (string, error) {
	snail, err := NewSnaili(common.HexToAddress(c.SnailAddress), c.Client)
	if err != nil {
		return "", err
	}

	tokenIdStr, _ := new(big.Int).SetString(tokenId, 10)
	tokenURI, err := snail.TokenURI(nil, tokenIdStr)
	if err != nil {
		return "", err
	}

	return tokenURI, nil
}
