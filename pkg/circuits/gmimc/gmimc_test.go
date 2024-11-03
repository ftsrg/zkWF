package gmimc_test

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	"github.com/ftsrg/zkWF/pkg/circuits/gmimc"
)

type EncryptDecryptTest struct {
	Input []frontend.Variable
	Key   []frontend.Variable
}

func (circuit EncryptDecryptTest) Define(api frontend.API) error {
	numRounds := gmimc.GetGMiMCRounds(len(circuit.Input))
	encrypted := gmimc.Encrypt(api, circuit.Input, circuit.Key, numRounds)
	api.Println("Encrypted:", encrypted)
	decrypted := gmimc.Decrypt(api, encrypted, circuit.Key, numRounds)
	api.Println("Decrypted:", decrypted)

	for i := 0; i < len(circuit.Input); i++ {
		api.AssertIsEqual(circuit.Input[i], decrypted[i])
	}

	return nil
}

func TestEncryptDecrypt(t *testing.T) {
	assert := test.NewAssert(t)

	var cubicCircuit EncryptDecryptTest = EncryptDecryptTest{
		Input: make([]frontend.Variable, 4),
		Key:   make([]frontend.Variable, 2),
	}

	var circuit EncryptDecryptTest = EncryptDecryptTest{
		Input: make([]frontend.Variable, 4),
		Key:   make([]frontend.Variable, 2),
	}
	circuit.Input = []frontend.Variable{1, 2, 3, 4}
	circuit.Key = []frontend.Variable{4, 5}

	assert.ProverSucceeded(&cubicCircuit, &circuit, test.WithCurves(ecc.BN254))
}
