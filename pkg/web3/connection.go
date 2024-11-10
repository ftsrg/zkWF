package web3

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ftsrg/zkWF/pkg/crypto/keys"
)

func CreateConnection(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth node: %w", err)
	}
	return client, nil
}

func CreateAuth(client *ethclient.Client, chainId *big.Int, keyFile string) (*bind.TransactOpts, error) {
	privateKey, err := keys.LoadEthKeyPair(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}
