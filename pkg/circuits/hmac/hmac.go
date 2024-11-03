package hmac

import (
	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/mimc"
	"github.com/ftsrg/zkWF/pkg/circuits/utils"
)

const BLOCK_SIZE = 64

func Hmac(api frontend.API, key, message []frontend.Variable) frontend.Variable {
	innerHash := innerHash(api, key, message)

	outerHash := outerHash(api, key, innerHash)

	return outerHash
}

func innerHash(api frontend.API, key, message []frontend.Variable) frontend.Variable {
	keyBytes := make([]frontend.Variable, len(key)*32) // 32 bytes per field element
	for i, k := range key {
		keyFieldDecomposed := utils.DecomposeToBytes(api, k)

		copy(keyBytes[i*32:], keyFieldDecomposed)
	}

	expandedKey := make([]frontend.Variable, BLOCK_SIZE)
	copy(expandedKey[:], keyBytes)

	for i := len(keyBytes); i < BLOCK_SIZE; i++ {
		expandedKey[i] = frontend.Variable(0x36)
	}

	for i, k := range keyBytes { // XOR with just the key, since 0x36 is already in the expanded key
		expandedKey[i] = bitwizeXor(api, k, frontend.Variable(0x36))
	}

	innerMessage := make([]frontend.Variable, len(message)+len(expandedKey))
	/*copy(innerMessage[:], expandedKey)
	copy(innerMessage[len(expandedKey):], message)*/
	for i := 0; i < len(expandedKey); i++ {
		innerMessage[i] = expandedKey[i]
	}
	for i := 0; i < len(message); i++ {
		innerMessage[i+len(expandedKey)] = message[i]
	}

	result := mimc.MultiMiMC5(api, 91, innerMessage, 0)

	return result
}

func bitwizeXor(api frontend.API, a, b frontend.Variable) frontend.Variable {
	bitsA := api.ToBinary(a, 8)
	bitsB := api.ToBinary(b, 8)
	xoredBits := make([]frontend.Variable, len(bitsA))
	for i := range bitsA {
		xoredBits[i] = api.Xor(bitsA[i], bitsB[i])
	}
	return api.FromBinary(xoredBits...)
}

func outerHash(api frontend.API, key []frontend.Variable, innerHash frontend.Variable) frontend.Variable {

	keyBytes := make([]frontend.Variable, len(key)*32) // 32 bytes per field element
	for i, k := range key {
		keyFieldDecomposed := utils.DecomposeToBytes(api, k)
		copy(keyBytes[i*32:], keyFieldDecomposed)
	}

	expKey := make([]frontend.Variable, BLOCK_SIZE)
	copy(expKey[:], keyBytes)

	for i := len(keyBytes); i < BLOCK_SIZE; i++ {
		expKey[i] = frontend.Variable(0x5c)
	}

	for i, k := range keyBytes {
		expKey[i] = bitwizeXor(api, k, frontend.Variable(0x5c))
	}

	outerMessage := make([]frontend.Variable, len(expKey)+1)
	copy(outerMessage[:], expKey)
	outerMessage[len(expKey)] = innerHash

	result := mimc.MultiMiMC5(api, 91, outerMessage, 0)

	return result
}
