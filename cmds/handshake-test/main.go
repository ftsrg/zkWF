package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ftsrg/zkWF/pkg/crypto/keys"
	"github.com/ftsrg/zkWF/pkg/web3"
)

const (
	participants = 3
)

func main() {
	for i := 0; i < participants; i++ {
		err := keys.GenerateKeyEthPair(fmt.Sprintf("key%d.json", i))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Key pairs generated successfully!")

	// Load the key pair
	privateKeys := make([]*eddsa.PrivateKey, participants)
	for i := 0; i < participants; i++ {
		keyPair, err := keys.LoadKeyPair(fmt.Sprintf("key%d.json", i))
		if err != nil {
			panic(err)
		}

		privateKey, err := eddsa.GenerateKey(rand.Reader)
		if err != nil {
			panic(err)
		}

		privateKey.SetBytes(keyPair.Bytes())
		privateKeys[i] = privateKey
	}
	pubKeys := make([]*eddsa.PublicKey, participants)
	for i := 0; i < 3; i++ {
		pubKeys[i] = &privateKeys[i].PublicKey
	}

	pubKeyStrs := make([]string, participants)
	for i := 0; i < participants; i++ {
		pubKeyStrs[i] = hex.EncodeToString(privateKeys[i].PublicKey.Bytes())
	}

	contractAddress, err := web3.DeployEcdhContract("http://localhost:8545", "eth_key.json", big.NewInt(31337), pubKeyStrs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Contract deployed successfully at address:", contractAddress)

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}

	ethPrivKeys := []string{"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d", "47c99abed3324a2707c28affff1267e45918ec8c3f20b8aa892e8b065d2942dd"}

	var wg sync.WaitGroup
	wg.Add(participants)
	for i, privKey := range ethPrivKeys {
		go func(i int, privKey string) {
			defer wg.Done()
			privateKey, err := crypto.HexToECDSA(privKey)
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

			secret, err := web3.PerformHandshake(client, contractAddress, auth, privateKeys[i], pubKeys)
			if err != nil {
				panic(err)
			}
			secretBig := new(big.Int)
			secretBig.SetString(hex.EncodeToString(secret), 16)

			fmt.Printf("Secret %d: %s\n", i, secretBig.String())
		}(i, privKey)
	}

	wg.Wait()
}
