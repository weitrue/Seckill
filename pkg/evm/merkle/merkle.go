package merkle

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

type Address string

// CalculateHash hashes the values of a TestContent
func (s Address) CalculateHash() ([]byte, error) {
	hx, e := hexutil.Decode(string(s))
	if e != nil {
		return []byte{}, e
	}
	h := sha3.NewLegacyKeccak256()
	if _, err := h.Write(hx); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (s Address) Equals(other Content) (bool, error) {
	return s == other.(Address), nil
}

func GetProof(address string, addresses []string) ([]string, error) {
	t, err := buildTree(addresses)
	if err != nil {
		return nil, err
	}
	path, _, err := t.GetMerklePath(Address(address))
	if err != nil {
		return []string{}, err
	}
	ret := make([]string, 0)
	for _, v := range path {
		ret = append(ret, fmt.Sprintf("0x%x", v))
	}
	return ret, nil
}

func GetRoot(addresses []string) (string, error) {
	t, err := buildTree(addresses)
	if err != nil {
		return "", err
	}

	root := t.MerkleRoot()

	return fmt.Sprintf("0x%x", root), nil
}

func buildTree(addresses []string) (*MerkleTree, error) {
	if len(addresses) == 0 {
		return nil, errors.New("no length")
	}

	if len(addresses) > 100000 {
		return nil, errors.New("size too large")
	}

	var list []Content
	for _, v := range addresses {
		list = append(list, Address(v))
	}

	t, err := NewTreeWithHashStrategy(list, sha3.NewLegacyKeccak256)
	if err != nil {
		return t, err
	}

	return t, nil
}
