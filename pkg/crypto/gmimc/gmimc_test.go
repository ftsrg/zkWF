package gmimc

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecryptBig(t *testing.T) {
	p := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4)}
	k := []*big.Int{big.NewInt(4), big.NewInt(5)}
	numRounds := GetGMiMCRounds(len(p))
	t.Log(numRounds)

	encrypted := EncryptBig(p, k, numRounds)
	t.Log(encrypted)
	decrypted := DecryptBig(encrypted, k, numRounds)
	t.Log(decrypted)

	assert.Equal(t, p, decrypted)
}
