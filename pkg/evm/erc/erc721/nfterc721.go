package erc721

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	ERC721 = "erc721"
)

type NftErc721 struct {
	client   *ethclient.Client
	endpoint string
}

func NewNftErc721(endpoint string) (*NftErc721, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	return &NftErc721{
		client:   client,
		endpoint: endpoint,
	}, nil
}

func (n *NftErc721) GetItemOwner(address string, tokenId string) (string, error) {
	addr := common.HexToAddress(address)
	instance, err := NewErc721Caller(addr, n.client)
	if err != nil {
		return "", err
	}

	token := new(big.Int)
	token.SetString(tokenId, 10)
	ownerOf, err := instance.OwnerOf(&bind.CallOpts{}, token)
	if err != nil {
		return "", err
	}

	return ownerOf.Hex(), nil
}
