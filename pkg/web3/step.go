package web3

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/backend/plonk"
	plonk_bn254 "github.com/consensys/gnark/backend/plonk/bn254"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ftsrg/zkWF/pkg/contracts/model"
	"github.com/ftsrg/zkWF/pkg/zkp"
)

func StepModel(url, keyPath, contractAddress string, chainID *big.Int, proofPath, witnessPath string) error {
	// Connect to the local Ethereum node
	client, err := CreateConnection(url)
	if err != nil {
		return fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}

	auth, err := CreateAuth(client, chainID, keyPath)
	if err != nil {
		return fmt.Errorf("failed to create auth: %w", err)
	}

	contract, err := model.NewModel(common.HexToAddress(contractAddress), client)
	if err != nil {
		return fmt.Errorf("failed to instantiate contract: %w", err)
	}

	proof, err := readProof(proofPath)
	if err != nil {
		return fmt.Errorf("failed to read proof: %w", err)
	}

	witness, err := zkp.LoadWitness(witnessPath)
	if err != nil {
		return fmt.Errorf("failed to load witness: %w", err)
	}

	vector := witness.Vector().(fr.Vector)
	publicInputs := make([]*big.Int, vector.Len())
	for i := 0; i < vector.Len(); i++ {
		publicInputs[i] = new(big.Int)
		vector[i].BigInt(publicInputs[i])
	}

	tx, err := contract.Update(auth, proof, publicInputs)
	if err != nil {
		return fmt.Errorf("failed to call Update: %w", err)
	}

	log.Println("Transaction hash:", tx.Hash().Hex())

	return nil
}

func readProof(path string) ([]byte, error) {
	proof := plonk.NewProof(ecc.BN254)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open proof file: %w", err)
	}

	_, err = proof.ReadFrom(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read proof from file: %w", err)
	}

	proof_bn254 := proof.(*plonk_bn254.Proof)

	return proof_bn254.MarshalSolidity(), nil
}
