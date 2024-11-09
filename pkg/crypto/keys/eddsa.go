package keys

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	bn254_eddsa "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark-crypto/signature"
)

const (
	sizeFr = fr.Bytes
)

type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

func GenerateKeyPair(keyPath string) error {
	// Generate a new key pair
	privateKey, err := bn254_eddsa.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	privateKeyHex := hex.EncodeToString(privateKey.Bytes())
	publicKey := privateKey.Public()

	publicKeyHex := hex.EncodeToString(publicKey.Bytes())

	scalar := GetPrivateKeyScaler(privateKey)
	fmt.Println("Private key scalar:", scalar.String())

	// Save the key pair to a file
	keyPair := KeyPair{
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
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

func HexToPrivateKey(hexKey string) (*bn254_eddsa.PrivateKey, error) {
	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	privKey, err := bn254_eddsa.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create private key: %w", err)
	}

	privKey.SetBytes(keyBytes)

	return privKey, nil
}

func HexToPublicKey(hexKey string) (*bn254_eddsa.PublicKey, error) {
	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %w", err)
	}

	privateKey, err := bn254_eddsa.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create public key: %w", err)
	}
	pubKey := privateKey.PublicKey

	pubKey.SetBytes(keyBytes)

	return &pubKey, nil
}

func LoadKeyPair(keyPath string) (signature.Signer, error) {
	jsonBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key pair: %w", err)
	}

	var keyPair KeyPair
	err = json.Unmarshal(jsonBytes, &keyPair)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal key pair: %w", err)
	}

	privateKey, err := HexToPrivateKey(keyPair.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	return privateKey, nil
}

func GetPrivateKeyScaler(privateKey *bn254_eddsa.PrivateKey) *big.Int {
	bytes := privateKey.Bytes()

	var buf [sizeFr]byte
	subtle.ConstantTimeCopy(1, buf[:], bytes[sizeFr:2*sizeFr])
	bScalar := new(big.Int)
	bScalar.SetBytes(buf[:])

	return bScalar
}

func DiffieHellmanStep(privateKey *bn254_eddsa.PrivateKey, point twistededwards.PointAffine) *twistededwards.PointAffine {
	scalar := GetPrivateKeyScaler(privateKey)
	return point.ScalarMultiplication(&point, scalar)
}
