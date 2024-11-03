package mimc

import (
	"math/big"
	"strings"

	"github.com/consensys/gnark-crypto/ecc"
)

var prime = ecc.BN254.ScalarField()

func MiMC7(nRounds int, xIn, k *big.Int) *big.Int {
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

	var t *big.Int = new(big.Int)
	t2 := make([]*big.Int, nRounds)
	t4 := make([]*big.Int, nRounds)
	t6 := make([]*big.Int, nRounds)
	t7 := make([]*big.Int, nRounds-1)
	var out *big.Int = new(big.Int)

	for i := 0; i < nRounds; i++ {
		if i == 0 {

			t = big.NewInt(0).Add(k, xIn)
			t.Mod(t, prime)
		} else {

			t.Add(k, t7[i-1])
			t.Mod(t, prime)
		}
		t2[i] = new(big.Int)
		t2[i].Mul(t, t)
		t2[i].Mod(t2[i], prime)
		t2[i].Mul(t2[i], t2[i])
		t4[i] = new(big.Int)
		t4[i].Mod(t4[i], prime)
		t6[i] = new(big.Int)
		t6[i].Mul(t4[i], t2[i])
		t6[i].Mod(t6[i], prime)
		if i < nRounds-1 {
			t7[i] = new(big.Int)
			t7[i].Mul(t6[i], t)
			t7[i].Mod(t7[i], prime)
		} else {
			//out = api.Add(api.Mul(t6[i], t), k)
			out.Mul(t6[i], t)
			out.Mod(out, prime)
			out.Add(out, k)
			out.Mod(out, prime)
		}
	}
	return out
}

func MultiMiMC7(nRounds int, in []*big.Int, k *big.Int) *big.Int {
	var out *big.Int = new(big.Int)

	nInputs := len(in)
	r := make([]*big.Int, len(in)+1)

	r[0] = k
	for i := 0; i < nInputs; i++ {
		mims := MiMC7(nRounds, in[i], r[i])
		r[i+1] = big.NewInt(0).Add(r[i], in[i])
		r[i+1].Mod(r[i+1], prime)
		//r[i+1] = r[i].Add(r[i], in[i], mims)
		r[i+1].Add(r[i+1], mims)
		r[i+1].Mod(r[i+1], prime)
	}
	out = r[nInputs]
	return out
}

func MiMC5(nRounds int, xIn, k *big.Int) *big.Int {

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

	var t *big.Int = new(big.Int)

	t2 := make([]*big.Int, nRounds)
	t4 := make([]*big.Int, nRounds)
	t5 := make([]*big.Int, nRounds)

	for i := 0; i < nRounds; i++ {
		if i == 0 {
			t = new(big.Int)
			t.Add(xIn, k)
			t.Mod(t, prime)
		} else {
			t = new(big.Int).Add(k, t5[i-1])
			t.Mod(t, prime)
			t.Add(t, c[i])
			t.Mod(t, prime)
		}
		//t2[i] = api.Mul(t, t)
		//t4[i] = api.Mul(t2[i], t2[i])
		t2[i] = new(big.Int)
		t2[i].Mul(t, t)
		t2[i].Mod(t2[i], prime)

		t4[i] = new(big.Int)
		t4[i].Mul(t2[i], t2[i])
		t4[i].Mod(t4[i], prime)

		/*if i < nRounds-1 {
			t5[i] = api.Mul(t4[i], t)
		} else {
			out = api.Add(api.Mul(t4[i], t), k)
		}*/
		t5[i] = new(big.Int)
		if i < nRounds-1 {
			t5[i].Mul(t4[i], t)
			t5[i].Mod(t5[i], prime)
		} else {
			t5[i].Mul(t4[i], t)
			t5[i].Mod(t5[i], prime)
			t5[i].Add(t5[i], k)
			t5[i].Mod(t5[i], prime)
		}
	}
	return t5[nRounds-1]
}

func MultiMiMC5(nRounds int, in []*big.Int, k *big.Int) *big.Int {
	var out *big.Int

	nInputs := len(in)
	r := make([]*big.Int, len(in)+1)

	r[0] = k
	for i := 0; i < nInputs; i++ {
		mims := MiMC5(nRounds, in[i], r[i])
		//r[i+1] = api.Add(r[i], in[i], mims)
		r[i+1] = big.NewInt(0).Add(r[i], in[i])
		r[i+1].Mod(r[i+1], prime)
		r[i+1].Add(r[i+1], mims)
		r[i+1].Mod(r[i+1], prime)
	}
	out = r[nInputs]
	return out
}
