package mimc

import (
	"math/big"
	"strings"

	"github.com/consensys/gnark/frontend"
)

func MiMC7(api frontend.API, nRounds int, xIn, k frontend.Variable) frontend.Variable {
	c := make([]*big.Int, 0)
	ss := strings.Split(cStr, ",")
	for _, v := range ss {
		v = strings.TrimSpace(v)
		if v != "" {
			a, _ := big.NewInt(0).SetString(v, 10)
			c = append(c, a)
		}
	}
	if len(c) != 91 {
		panic("parse c error")
	}

	var t frontend.Variable
	t2 := make([]frontend.Variable, nRounds)
	t4 := make([]frontend.Variable, nRounds)
	t6 := make([]frontend.Variable, nRounds)
	t7 := make([]frontend.Variable, nRounds-1)
	var out frontend.Variable

	for i := 0; i < nRounds; i++ {
		if i == 0 {
			t = api.Add(k, xIn)
		} else {
			t = api.Add(k, t7[i-1], c[i])
		}
		t2[i] = api.Mul(t, t)
		t4[i] = api.Mul(t2[i], t2[i])
		t6[i] = api.Mul(t4[i], t2[i])
		if i < nRounds-1 {
			t7[i] = api.Mul(t6[i], t)
		} else {
			out = api.Add(api.Mul(t6[i], t), k)
		}
	}
	return out
}

func MultiMiMC7(api frontend.API, nRounds int, in []frontend.Variable, k frontend.Variable) frontend.Variable {
	var out frontend.Variable

	nInputs := len(in)
	r := make([]frontend.Variable, len(in)+1)

	r[0] = k
	for i := 0; i < nInputs; i++ {
		api.Println("i", i, in[i], r[i])
		mims := MiMC7(api, nRounds, in[i], r[i])
		api.Println("mims", mims)
		r[i+1] = api.Add(r[i], in[i], mims)
	}
	out = r[nInputs]
	return out
}

func MiMC5(api frontend.API, nRounds int, xIn, k frontend.Variable) frontend.Variable {

	c := make([]*big.Int, 0)
	ss := strings.Split(cStr, ",")
	for _, v := range ss {
		v = strings.TrimSpace(v)
		if v != "" {
			a, _ := big.NewInt(0).SetString(v, 10)
			c = append(c, a)
		}
	}
	if len(c) != 91 {
		panic("parse c error")
	}

	var t frontend.Variable
	t2 := make([]frontend.Variable, nRounds)
	t4 := make([]frontend.Variable, nRounds)
	t5 := make([]frontend.Variable, nRounds)

	for i := 0; i < nRounds; i++ {
		if i == 0 {
			t = api.Add(k, xIn)
		} else {
			t = api.Add(k, t5[i-1], c[i])
		}
		t2[i] = api.Mul(t, t)
		t4[i] = api.Mul(t2[i], t2[i])
		if i < nRounds-1 {
			t5[i] = api.Mul(t4[i], t)
		} else {
			t5[i] = api.Add(api.Mul(t4[i], t), k)
		}
		//t5[i] = api.Mul(t4[i], t)
	}
	return t5[nRounds-1]
}

func MultiMiMC5(api frontend.API, nRounds int, in []frontend.Variable, k frontend.Variable) frontend.Variable {
	var out frontend.Variable

	nInputs := len(in)
	r := make([]frontend.Variable, len(in)+1)

	r[0] = k
	for i := 0; i < nInputs; i++ {
		mims := MiMC5(api, nRounds, in[i], r[i])
		r[i+1] = api.Add(r[i], in[i], mims)
	}
	out = r[nInputs]
	return out
}
