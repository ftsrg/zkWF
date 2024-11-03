package web3

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ecdh_contract "github.com/ftsrg/zkWF/pkg/contracts/ecdh"
	"github.com/ftsrg/zkWF/pkg/crypto/keys"
)

// UploadIntermediateValue uploads a computed intermediate value to the smart contract
func uploadIntermediateValue(client *ethclient.Client, contractAddress string, auth *bind.TransactOpts, combo string, value []byte) error {
	// Instantiate the contract
	instance, err := ecdh_contract.NewEcdh(common.HexToAddress(contractAddress), client)
	if err != nil {
		return fmt.Errorf("failed to bind to deployed contract: %v", err)
	}
	fmt.Println("Before Nonce: ", auth.Nonce)
	// Upload the intermediate value
	tx, err := instance.UploadIntermediateValue(auth, combo, value)
	if err != nil {
		return fmt.Errorf("failed to upload intermediate value: %v", err)
	}

	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

	fmt.Println("After Nonce: ", auth.Nonce)

	fmt.Printf("Uploaded intermediate value for combo %s: tx hash: %s\n", combo, tx.Hash().Hex())
	return nil
}

// GetIntermediateValue retrieves a stored intermediate value from the smart contract
func getIntermediateValue(client *ethclient.Client, contractAddress string, combo string) ([]byte, error) {
	// Bind the deployed contract instance
	instance, err := ecdh_contract.NewEcdh(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to deployed contract: %v", err)
	}

	// Retrieve the intermediate value
	value, err := instance.GetIntermediateValue(&bind.CallOpts{}, combo)
	if err != nil {
		return nil, fmt.Errorf("failed to get intermediate value: %v", err)
	}

	return value, nil
}

// ComputeFinalSecret computes the final shared secret using retrieved intermediate values
func computeFinalSecret(priv *eddsa.PrivateKey, intermediateValues [][]byte) ([]byte, error) {
	var finalSecret []byte
	for _, value := range intermediateValues {
		// Convert value to ECDH public key
		var intermediateValue twistededwards.PointAffine
		intermediateValue.SetBytes(value)

		// Compute the final shared secret by multiplying the private key with the intermediate values
		sharedSecret := keys.DiffieHellmanStep(priv, intermediateValue)

		sharedSecretBytes := sharedSecret.Bytes()
		finalSecret = sharedSecretBytes[:]
	}

	return finalSecret, nil
}

// PerformHandshake calculates shared secrets for a participant using X25519 and interacts with the smart contract
func PerformHandshake(client *ethclient.Client, contractAddress string, auth *bind.TransactOpts, priv *eddsa.PrivateKey, publicKeys []*eddsa.PublicKey) ([]byte, error) {

	pub := priv.PublicKey
	// Find the index of the private key in the list of public keys
	index := -1
	for i, pk := range publicKeys {
		if pk.A.X.Equal(&pub.A.X) && pk.A.Y.Equal(&pub.A.Y) {
			index = i
			break
		}
	}
	if index == -1 {
		log.Fatalln("Private key not found in the list of public keys")
	}

	// Loop over public keys to compute and upload intermediate values
	for i, xPub := range publicKeys {
		if i == index {
			continue
		}
		sharedKey := keys.DiffieHellmanStep(priv, xPub.A)

		// Generate a unique combination label for the intermediate value (e.g., "P1-P2")
		combo := fmt.Sprintf("P%d-P%d", index, i)

		// Upload the intermediate value to the smart contract
		sharedKeyBytes := sharedKey.Bytes()
		err := uploadIntermediateValue(client, contractAddress, auth, combo, sharedKeyBytes[:])
		if err != nil {
			return nil, fmt.Errorf("failed to upload intermediate value for combo %s: %v", combo, err)
		}
	}

	// Wait a short time for other participants to upload their values
	time.Sleep(1 * time.Second) // Adjust based on your network latency and setup

	// Retrieve all necessary intermediate values from the smart contract
	intermediateValues := [][]byte{}
	for i := range publicKeys {
		if i == index {
			continue
		}
		/*combo := fmt.Sprintf("P%d-P%d", index, i)
		value, err := getIntermediateValue(client, contractAddress, combo)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve intermediate value for combo %s: %v", combo, err)
		}
		intermediateValues = append(intermediateValues, value)*/
		for j := range publicKeys {
			if j == index || j == i {
				continue
			}
			combo := fmt.Sprintf("P%d-P%d", i, j)
			value, err := getIntermediateValue(client, contractAddress, combo)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve intermediate value for combo %s: %v", combo, err)
			}
			intermediateValues = append(intermediateValues, value)
		}
	}

	// Compute the final shared secret using the retrieved intermediate values
	finalSecret, err := computeFinalSecret(priv, intermediateValues)
	if err != nil {
		return nil, fmt.Errorf("failed to compute final shared secret: %v", err)
	}

	fmt.Printf("Final shared secret: %x\n", finalSecret)
	return finalSecret, nil
}
