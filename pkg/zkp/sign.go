package zkp

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/hash"
	"github.com/ftsrg/zkWF/pkg/crypto/keys"
)

func SignHash(keysPath, inputPath string) error {
	privateKey, err := keys.LoadKeyPair(keysPath)
	if err != nil {
		return fmt.Errorf("failed to load key pair: %w", err)
	}

	input, err := loadInputs(inputPath)
	if err != nil {
		return fmt.Errorf("failed to load inputs: %w", err)
	}

	hFunc := hash.MIMC_BN254.New()

	var hashBig *big.Int
	hashBig, _ = big.NewInt(0).SetString(input.HashNew, 10)

	signature, err := privateKey.Sign(hashBig.Bytes(), hFunc)

	if err != nil {
		return fmt.Errorf("failed to sign: %w", err)
	}

	fmt.Println("Signature:", hex.EncodeToString(signature))

	return nil
}
