package gmimc

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/iden3/go-iden3-crypto/ff"
	"github.com/iden3/go-iden3-crypto/keccak256"
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

/*
	Pseudo code

def encrypt(p, k, num_rounds, rc):

	state = list(p)
	ks = list(k)

	for i in range(0, num_rounds):
	    # F
	    temp = state[0] + ks[i % t] + rc[i]
	    temp = temp^3
	    print(i, temp)
	    # Add F
	    for j in range(1, t):
	        state[j] = state[j] + temp
	    # Feistel swap
	    if i < (num_rounds - 1):
	        state = state[1:] + state[:1]

	return state
*/

func Encrypt(api frontend.API, input []frontend.Variable, keyState []frontend.Variable, nRounds int) []frontend.Variable {
	state := make([]frontend.Variable, len(input))
	copy(state, input)

	for i := 0; i < nRounds; i++ {
		// F
		temp := state[0]
		t := api.Add(api.Add(temp, keyState[i%len(keyState)]), constants.cts[i])

		t2 := api.Mul(t, t)
		t4 := api.Mul(t2, t2)
		r := api.Mul(api.Mul(t4, t2), t)
		// Add F
		for j := 1; j < len(state); j++ {
			state[j] = api.Add(state[j], r)
		}
		// Feistel swap
		if i < (nRounds - 1) {
			new_state := make([]frontend.Variable, len(state))
			new_state[0] = state[len(state)-1]
			for j := 1; j < len(state); j++ {
				new_state[j] = state[j-1]
			}
			state = new_state
		}
	}

	return state
}

/*
		Pseudo code
		def decrypt(c, k, num_rounds, rc):

	    state = list(c)
	    ks = list(reversed(list(k)))
	    rc = list(reversed(rc))

	    for i in range(0, num_rounds):
	        # F
	        temp = state[0] + ks[i % t] + rc[i]
	        temp = temp^3
	        print(num_rounds - i, temp)
	        # Add F
	        for j in range(1, t):
	            state[j] = state[j] - temp
	        # Feistel swap
	        if i < (num_rounds - 1):
	            state = state[t-1:] + state[:t-1]

	    return state
*/

func Decrypt(api frontend.API, input []frontend.Variable, keyState []frontend.Variable, nRounds int) []frontend.Variable {
	state := make([]frontend.Variable, len(input))
	copy(state, input)

	for i := 0; i < nRounds; i++ {
		// F
		temp := state[0]
		t := api.Add(api.Add(temp, keyState[i%len(keyState)]), constants.cts[nRounds-i-1])
		t2 := api.Mul(t, t)
		t4 := api.Mul(t2, t2)
		r := api.Mul(api.Mul(t4, t2), t)
		// Add F
		for j := 1; j < len(state); j++ {
			state[j] = api.Sub(state[j], r)
		}
		// Feistel swap
		if i < (nRounds - 1) {
			new_state := make([]frontend.Variable, len(state))
			new_state[len(new_state)-1] = state[0]
			for j := 0; j < len(state)-1; j++ {
				new_state[j] = state[j+1]
			}
			state = new_state
		}
	}

	return state
}

/*
	Pseudo code

def get_gmimc_rounds(t, d, p):

	r1 = 2 * (1 + t + t^2)
	r2 = ceil(2 * log(p, d) + 2 * t)
	if r2 > r1:
	    return r2
	return r1
*/
func GetGMiMCRounds(branches int) int {
	r1 := 2 * (1 + branches + branches*branches)
	r2 := r2lg7 + 2*branches // ceil(2 * log(7, branches) + 2 * branches)
	if r2 > r1 {
		return r2
	}
	return r1
}
