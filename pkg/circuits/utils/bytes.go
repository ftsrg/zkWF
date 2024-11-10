package utils

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/bits"
)

func DecomposeToBytes(api frontend.API, element frontend.Variable) []frontend.Variable {
	pows := []int{1, 2, 4, 8, 16, 32, 64, 128}
	// BN254 field elements fit into 32 bytes (256 bits).
	numBytes := 32
	bytes := make([]frontend.Variable, numBytes)
	bits := bits.ToBinary(api, element)
	for i := 0; i < numBytes; i++ {
		bytes[i] = 0
		for j := 0; j < 8; j++ {
			if i*8+j < 254 {
				bytes[i] = api.Add(bytes[i], api.Mul(bits[i*8+j], pows[j]))
			}
		}
	}

	return bytes
}

// CompressToFieldElement concatenates the given states (a number from 0 to 10) into a single field element.
// Since the states are 4 bits long, and a field element is 254 bits long, we can fit 63 states into a single field element.
func CompressToFieldElement(api frontend.API, states []frontend.Variable) frontend.Variable {
	if len(states) > 63 {
		panic("Too many states to compress into a single field element.")
	}

	// Initialize the compressed field element
	var compressed frontend.Variable = 0
	// Compress the states into a single field element
	for i := 0; i < len(states); i++ {
		// Shift the compressed field element by 4 bits to make space for the next state
		compressed = api.Mul(compressed, 16)
		// Add the next state to the compressed field element
		compressed = api.Add(compressed, states[i])
	}

	return compressed
}

func CompressToFieldElementBig(states []int64) frontend.Variable {
	if len(states) > 63 {
		panic("Too many states to compress into a single field element.")
	}
	// ecc.BN254.ScalarField()
	// Initialize the compressed field element
	var compressed *big.Int = big.NewInt(0)
	// Compress the states into a single field element
	for i := 1; i < len(states); i++ {
		// Shift the compressed field element by 4 bits to make space for the next state
		compressed = new(big.Int).Mul(compressed, big.NewInt(16))
		//compressed = new(big.Int).Mod(compressed, ecc.BN254.ScalarField())
		// Add the next state to the compressed field element
		//compressed = api.Add(compressed, states[i])
		compressed = new(big.Int).Add(compressed, big.NewInt(states[i]))
		//compressed = new(big.Int).Mod(compressed, ecc.BN254.ScalarField())
	}

	return compressed
}
