package web3

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ftsrg/zkWF/pkg/contracts/ecdh"
)

// DeployContract deploys the ECDH contract with predefined public keys
func DeployContract(publicKeys []string) (string, error) {

	// Connect to the local Ethereum node
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}

	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

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

	fmt.Printf("Contract deployed! Address: %s, Transaction Hash: %s\n", address.Hex(), tx.Hash().Hex())
	return address.Hex(), nil
}
