package hkdf_test

import (
	"math/big"
	"testing"

	"github.com/ftsrg/zkWF/pkg/crypto/hkdf"
)

func TestHkdf(t *testing.T) {
	key, _ := new(big.Int).SetString("14670954021912277657185193395461625777238371161394276953777323017688456414673", 10)
	r1, _ := new(big.Int).SetString("10623130405679805330576346668288554504439363553729204031500032016204369167485", 10)
	r2, _ := new(big.Int).SetString("4585745623398727590722715144656109933609484040934873213979218735388288847795", 10)
	salt := []*big.Int{big.NewInt(0)}
	ikm := []*big.Int{key}

	info := []*big.Int{r1, r2}

	res := hkdf.Hkdf(salt, ikm, info, 2)
	expected1, _ := new(big.Int).SetString("17366330016152514483841990308612469611643178212099168530126231105881584360633", 10)
	if res[0].Cmp(expected1) != 0 {
		t.Errorf("Expected %v, got %v", expected1, res[0])
	}

}
