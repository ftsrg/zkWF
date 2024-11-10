package web3

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ftsrg/zkWF/pkg/contracts/ecdh"
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
