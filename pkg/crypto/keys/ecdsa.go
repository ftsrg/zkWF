package keys

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

type EthKeyPair struct {
	PrivateKey string
	PublicKey  string
}

func GenerateKeyEthPair(keyPath string) error {
	// Generate a new key pair
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Save the key pair to a file
	keyPair := EthKeyPair{
		PrivateKey: privateKeyHex,
		PublicKey:  address.String(),
	}

	jsonBytes, err := json.Marshal(keyPair)
	if err != nil {
		return fmt.Errorf("failed to marshal key pair: %w", err)
	}

	err = os.WriteFile(keyPath, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to save key pair: %w", err)
	}

	return nil
}

func HexToEthPrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	privKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	return privKey, nil
}

func LoadEthKeyPair(keyPath string) (*ecdsa.PrivateKey, error) {
	jsonBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key pair: %w", err)
	}

	var keyPair EthKeyPair
	err = json.Unmarshal(jsonBytes, &keyPair)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal key pair: %w", err)
	}

	return HexToEthPrivateKey(keyPair.PrivateKey)
}
