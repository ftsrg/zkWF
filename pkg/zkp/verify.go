package zkp

import (
	"fmt"
	"log"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
)

func VerifyProof(proofPath, verificationKeyPath, witnessPath string) error {

	proof := plonk.NewProof(ecc.BN254)
	file, err := os.Open(proofPath)
	if err != nil {
		return fmt.Errorf("failed to open proof file: %w", err)
	}

	_, err = proof.ReadFrom(file)
	if err != nil {
		return fmt.Errorf("failed to read proof from file: %w", err)
	}

	vk, err := loadVerificationKey(verificationKeyPath)
	if err != nil {
		return fmt.Errorf("failed to load verification key: %w", err)
	}

	witness, err := loadWitness(witnessPath)
	if err != nil {
		return fmt.Errorf("failed to load witness: %w", err)
	}

	err = plonk.Verify(proof, vk, witness)
	if err != nil {
		return fmt.Errorf("verification failed: %w", err)
	}

	log.Println("Verification successful!")

	return nil
}

func loadVerificationKey(path string) (plonk.VerifyingKey, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening proving key file: %w", err)
	}
	defer file.Close()

	var vk plonk.VerifyingKey = plonk.NewVerifyingKey(ecc.BN254)

	_, err = vk.ReadFrom(file)
	if err != nil {
		return nil, fmt.Errorf("error reading proving key file: %w", err)
	}

	return vk, nil
}
