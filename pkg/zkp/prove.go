package zkp

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	plonk_bn254 "github.com/consensys/gnark/backend/plonk/bn254"
	"github.com/consensys/gnark/backend/witness"
)

func Prove(r1csPath, pkPath, witnessPath, proofPath string) error {
	ccs, err := loadProgram(r1csPath)
	if err != nil {
		return fmt.Errorf("error loading circuit: %w", err)
	}

	pk, err := loadProvingKey(pkPath)
	if err != nil {
		return fmt.Errorf("error loading proving key: %w", err)
	}

	witnessFull, err := LoadWitness(witnessPath)
	if err != nil {
		return fmt.Errorf("error loading witness: %w", err)
	}

	proof, err := plonk.Prove(ccs, pk, witnessFull)
	if err != nil {
		return fmt.Errorf("error proving: %w", err)
	}

	proofFile, err := os.Create(proofPath)
	if err != nil {
		return fmt.Errorf("error creating proof file: %w", err)
	}
	defer proofFile.Close()

	_, err = proof.WriteTo(proofFile)
	if err != nil {
		return fmt.Errorf("error writing proof to file: %w", err)
	}

	bn254proof := proof.(*plonk_bn254.Proof)
	fmt.Println("Soldity Proof:", hex.EncodeToString(bn254proof.MarshalSolidity()))

	return nil
}

func loadProvingKey(path string) (plonk.ProvingKey, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening proving key file: %w", err)
	}
	defer file.Close()

	var pk plonk.ProvingKey = plonk.NewProvingKey(ecc.BN254)

	_, err = plonk.ProvingKey.ReadFrom(pk, file)
	if err != nil {
		return nil, fmt.Errorf("error reading proving key file: %w", err)
	}

	return pk, nil
}

func LoadWitness(path string) (witness.Witness, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening witness file: %w", err)
	}
	defer file.Close()

	wts, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		return nil, fmt.Errorf("error creating witness: %w", err)
	}

	_, err = wts.ReadFrom(file)
	if err != nil {
		return nil, fmt.Errorf("error reading witness file: %w", err)
	}

	return wts, nil
}
