package hkdf

import (
	"math/big"

	"github.com/ftsrg/zkWF/pkg/crypto/hmac"
)

func Hkdf(salt, ikm, info []*big.Int, l int) []*big.Int {
	prk := extract(salt, ikm)

	return expand(prk, info, l)
}

func hash(key, data []*big.Int) *big.Int {
	return hmac.Hmac(key, data)
}

func extract(salt, ikm []*big.Int) *big.Int {
	return hash(salt, ikm)
}

func expand(key *big.Int, info []*big.Int, length int) []*big.Int {
	prk := make([]*big.Int, 1)
	prk[0] = key

	okm := make([]*big.Int, length)
	var t *big.Int = big.NewInt(0)

	for i := 0; i < length; i++ {
		hashInput := make([]*big.Int, len(info)+1+1)
		copy(hashInput[:], info)
		hashInput[len(info)] = t
		hashInput[len(info)+1] = big.NewInt(int64(i))
		t = hash(prk, hashInput)
		okm[i] = t
	}

	return okm
}
