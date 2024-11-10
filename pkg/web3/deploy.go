package web3

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ftsrg/zkWF/pkg/contracts/ecdh"
	"github.com/ftsrg/zkWF/pkg/contracts/model"
)

// DeployEcdhContract deploys the ECDH contract with predefined public keys
func DeployEcdhContract(url, keyPath string, chainID *big.Int, publicKeys []string) (string, error) {

	// Connect to the local Ethereum node
	client, err := CreateConnection(url)
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}

	auth, err := CreateAuth(client, chainID, keyPath)
	if err != nil {
		return "", fmt.Errorf("failed to create auth: %w", err)
	}

	// Convert public keys to byte arrays
	var keys [][]byte
	for _, pk := range publicKeys {
		pkBytes, err := hex.DecodeString(pk)
		if err != nil {
			return "", err
		}
		keys = append(keys, pkBytes)
	}

	address, tx, _, err := ecdh.DeployEcdh(auth, client, keys)
	if err != nil {
		return "", err
	}

	log.Printf("Contract deployed! Address: %s, Transaction Hash: %s\n", address.Hex(), tx.Hash().Hex())
	return address.Hex(), nil
}

func DeployModelContract(url, keyPath string, chainID *big.Int, initialHash *big.Int, initialState []*big.Int) (string, error) {
	// Connect to the local Ethereum node
	client, err := CreateConnection(url)
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}

	auth, err := CreateAuth(client, chainID, keyPath)
	if err != nil {
		return "", fmt.Errorf("failed to create auth: %w", err)
	}

	cotractBin, err := compileContract()
	if err != nil {
		return "", fmt.Errorf("failed to compile contract: %w", err)
	}

	abi, err := model.ModelMetaData.GetAbi()
	if err != nil {
		return "", fmt.Errorf("failed to get contract abi: %w", err)
	}

	address, tx, _, err := bind.DeployContract(auth, *abi, cotractBin, client, initialHash, initialState)
	if err != nil {
		return "", err
	}

	log.Printf("Contract deployed! Address: %s, Transaction Hash: %s\n", address.Hex(), tx.Hash().Hex())
	return address.Hex(), nil
}
