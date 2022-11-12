package eip

import (
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
)

func ToCheckSumAddress(address string) (string, error) {
	if address == "" {
		return "", errors.New("empty address")
	}

	if strings.HasPrefix(address, "0x") {
		address = address[2:]
	}
	
	bytes, err := hex.DecodeString(address)
	if err != nil {
		return "", err
	}

	hash := calculateKeccak256([]byte(strings.ToLower(address)))
	result := "0x"
	for i, b := range bytes {
		result += checksumByte(b>>4, hash[i]>>4)
		result += checksumByte(b&0xF, hash[i]&0xF)
	}

	return result, nil
}

func checksumByte(addr byte, hash byte) string {
	result := strconv.FormatUint(uint64(addr), 16)
	if hash >= 8 {
		return strings.ToUpper(result)
	} else {
		return result
	}
}

func calculateKeccak256(addr []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(addr)
	return hash.Sum(nil)
}
