package gmimc

import (
	"log"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/iden3/go-iden3-crypto/ff"
	"github.com/iden3/go-iden3-crypto/keccak256"
	"github.com/iden3/go-iden3-crypto/utils"
)

// SEED defines the seed used to constants
const SEED = "mimc"
const r2lg3 = 321
const r2lg7 = 181

var constants = generateConstantsData()

type constantsData struct {
	seedHash *big.Int
	iv       *big.Int
	nRounds  int
	cts      []*ff.Element
}

// Source: https://github.com/iden3/go-iden3-crypto/blob/fb1d25298f545c79bdf8173f2e83bc808e15df61/mimc7/mimc7.go#L25C1-L36C2
func generateConstantsData() constantsData {
	var consts constantsData

	consts.seedHash = new(big.Int).SetBytes(keccak256.Hash([]byte(SEED)))
	c := new(big.Int).SetBytes(keccak256.Hash([]byte(SEED + "_iv")))
	consts.iv = new(big.Int).Mod(c, ecc.BN254.ScalarField())

	consts.nRounds = 800
	cts := getConstants(SEED, consts.nRounds)
	consts.cts = cts
	return consts
}

func getConstants(seed string, nRounds int) []*ff.Element {
	cts := make([]*ff.Element, nRounds)
	cts[0] = ff.NewElement()
	c := new(big.Int).SetBytes(keccak256.Hash([]byte(seed)))
	for i := 1; i < nRounds; i++ {
		c = new(big.Int).SetBytes(keccak256.Hash(c.Bytes()))

		n := new(big.Int).Mod(c, ecc.BN254.ScalarField())
		cts[i] = ff.NewElement().SetBigInt(n)
	}
	return cts
}

func EncryptBig(input []*big.Int, key []*big.Int, nRounds int) []*big.Int {
	if !utils.CheckBigIntArrayInField(input) || !utils.CheckBigIntArrayInField(key) {
		log.Fatalln("inputs values not inside Finite Field")
	}
	// Copy input to state
	state := make([]*big.Int, len(input))
	for i, v := range input {
		state[i] = new(big.Int).Set(v)
	}

	for i := 0; i < nRounds; i++ {
		constr := new(big.Int).SetInt64(0)
		constants.cts[i].ToBigIntRegular(constr)
		// F
		temp := new(big.Int).Add(new(big.Int).Add(state[0], key[i%len(key)]), constr)
		temp = new(big.Int).Mod(temp, ecc.BN254.ScalarField())
		temp.Exp(temp, big.NewInt(7), ecc.BN254.ScalarField())
		// Add F
		for j := 1; j < len(state); j++ {
			state[j].Add(state[j], temp)
			state[j].Mod(state[j], ecc.BN254.ScalarField())
		}
		// Feistel swap
		if i < (nRounds - 1) {
			new_state := make([]*big.Int, len(state))
			new_state[0] = state[len(state)-1]
			for j := 1; j < len(state); j++ {
				new_state[j] = state[j-1]
			}
			state = new_state
		}

	}
	return state
}

func DecryptBig(input []*big.Int, key []*big.Int, nRounds int) []*big.Int {

	// Copy input to state
	state := make([]*big.Int, len(input))
	for i, v := range input {
		state[i] = new(big.Int).Set(v)
	}

	for i := 0; i < nRounds; i++ {
		constr := new(big.Int).SetInt64(0)
		constants.cts[nRounds-i-1].ToBigIntRegular(constr)
		// F
		temp := new(big.Int).Add(new(big.Int).Add(state[0], key[i%len(key)]), constr)
		temp = new(big.Int).Mod(temp, ecc.BN254.ScalarField())
		temp.Exp(temp, big.NewInt(7), ecc.BN254.ScalarField())
		// Add F
		for j := 1; j < len(state); j++ {
			state[j].Sub(state[j], temp)
			state[j].Mod(state[j], ecc.BN254.ScalarField())
		}
		// Feistel swap
		if i < (nRounds - 1) {
			new_state := make([]*big.Int, len(state))
			new_state[len(new_state)-1] = state[0]
			for j := 0; j < len(state)-1; j++ {
				new_state[j] = state[j+1]
			}
			state = new_state
		}
	}
	return state
}

func GetGMiMCRounds(branches int) int {
	r1 := 2 * (1 + branches + branches*branches)
	r2 := r2lg7 + 2*branches // ceil(2 * log(7, branches) + 2 * branches)
	if r2 > r1 {
		return r2
	}
	return r1
}
