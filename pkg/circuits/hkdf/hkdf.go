package hkdf

import (
	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/hmac"
)

func Hkdf(api frontend.API, salt, ikm, info []frontend.Variable, l int) []frontend.Variable {
	prk := extract(api, salt, ikm)

	return expand(api, prk, info, l)
}

func hash(api frontend.API, key, data []frontend.Variable) frontend.Variable {
	return hmac.Hmac(api, key, data)
}

func extract(api frontend.API, salt, ikm []frontend.Variable) frontend.Variable {
	return hash(api, salt, ikm)
}

func expand(api frontend.API, key frontend.Variable, info []frontend.Variable, length int) []frontend.Variable {
	prk := make([]frontend.Variable, 1)
	prk[0] = key

	okm := make([]frontend.Variable, length)
	var t frontend.Variable = 0

	for i := 0; i < length; i++ {
		hashInput := make([]frontend.Variable, len(info)+1+1)
		copy(hashInput[:], info)
		hashInput[len(info)] = t
		hashInput[len(info)+1] = frontend.Variable(i)
		t = hash(api, prk, hashInput)
		okm[i] = t
	}

	return okm
}
