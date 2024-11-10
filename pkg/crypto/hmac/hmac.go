package hmac

import (
	"math/big"

	"github.com/ftsrg/zkWF/pkg/crypto/mimc"
)

const BLOCK_SIZE = 64

func Hmac(key, message []*big.Int) *big.Int {
	innerHash := innerHash(key, message)

	outerHash := outerHash(key, innerHash)

	return outerHash
}

func innerHash(key, message []*big.Int) *big.Int {
	var keyBytes []byte
	for _, k := range key {
		keyFieldDecomposed := reverseBytes(k.Bytes()) //utils.DecomposeToBytes(api, k)
		keyBytes = append(keyBytes, keyFieldDecomposed...)
	}

	expandedKey := make([]byte, BLOCK_SIZE)
	copy(expandedKey[:], keyBytes)

	for i := len(keyBytes); i < BLOCK_SIZE; i++ {
		expandedKey[i] = 0x36
	}

	for i, k := range keyBytes { // XOR with just the key, since 0x36 is already in the expanded key
		expandedKey[i] = bitwizeXor(k, 0x36)
	}

	innerMessage := make([]*big.Int, len(message)+len(expandedKey))
	/*copy(innerMessage[:], expandedKey)
	copy(innerMessage[len(expandedKey):], message)*/
	for i := 0; i < len(expandedKey); i++ {
		expandedKeyBig := new(big.Int)
		bytes := []byte{expandedKey[i]}
		expandedKeyBig.SetBytes(bytes)

		innerMessage[i] = expandedKeyBig
	}
	for i := 0; i < len(message); i++ {
		innerMessage[i+len(expandedKey)] = message[i]
	}

	result := mimc.MultiMiMC5(91, innerMessage, big.NewInt(0))

	return result
}

func bitwizeXor(a, b byte) byte {
	return a ^ b
}

func reverseBytes(input []byte) []byte {
	var result []byte = make([]byte, len(input))

	for i := range input {
		result[i] = input[len(input)-1-i]
	}

	return result
}

func outerHash(key []*big.Int, innerHash *big.Int) *big.Int {

	var keyBytes []byte
	for _, k := range key {
		keyFieldDecomposed := reverseBytes(k.Bytes()) //utils.DecomposeToBytes(api, k)
		keyBytes = append(keyBytes, keyFieldDecomposed...)
	}

	expandedKey := make([]byte, BLOCK_SIZE)
	copy(expandedKey[:], keyBytes)

	for i := len(keyBytes); i < BLOCK_SIZE; i++ {
		expandedKey[i] = 0x5c
	}

	for i, k := range keyBytes {
		expandedKey[i] = bitwizeXor(k, 0x5c)
	}

	outerMessage := make([]*big.Int, len(expandedKey)+1)
	//copy(outerMessage[:], expandedKey)
	for i, kByte := range expandedKey {
		expandedKeyBig := new(big.Int)
		bytes := []byte{kByte}
		expandedKeyBig.SetBytes(bytes)

		outerMessage[i] = expandedKeyBig
	}

	outerMessage[len(expandedKey)] = innerHash

	result := mimc.MultiMiMC5(91, outerMessage, big.NewInt(0))

	return result
}
