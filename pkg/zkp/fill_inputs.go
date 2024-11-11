package zkp

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/hash"
	"github.com/ftsrg/zkWF/pkg/crypto/gmimc"
	"github.com/ftsrg/zkWF/pkg/crypto/hkdf"
	"github.com/ftsrg/zkWF/pkg/crypto/keys"
	"github.com/ftsrg/zkWF/pkg/crypto/mimc"
)

func FillInputs(inputFilePath, keysPath string) error {
	privateKey, err := keys.LoadKeyPair(keysPath)
	if err != nil {
		return fmt.Errorf("failed to load key pair: %w", err)
	}

	input, err := loadInputs(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to load inputs: %w", err)
	}

	if input.State_curr.Radomness == "0" {
		input.State_curr.Radomness = generateSalt().String()
		fmt.Println("Randomness:", input.State_curr.Radomness)
	}

	if input.State_new.Radomness == "0" {
		input.State_new.Radomness = generateSalt().String()
		fmt.Println("Randomness:", input.State_new.Radomness)
	}

	compressedStateCurr := compressState(input.State_curr)
	compressedStateNew := compressState(input.State_new)
	bigZero := new(big.Int).SetInt64(0)

	input.HashCurr = mimc.MultiMiMC5(91, compressedStateCurr, bigZero).String()
	log.Println("HashCurr:", input.HashCurr)
	input.HashNew = mimc.MultiMiMC5(91, compressedStateNew, bigZero).String()
	log.Println("HashNew:", input.HashNew)
	//input.PublicKey = hex.EncodeToString(privateKey.Public().Bytes())
	//log.Println("PublicKey:", input.PublicKey)

	hFunc := hash.MIMC_BN254.New()

	var hashBig *big.Int
	hashBig, _ = big.NewInt(0).SetString(input.HashNew, 10)

	signature, err := privateKey.Sign(hashBig.Bytes(), hFunc)

	if err != nil {
		return fmt.Errorf("failed to sign: %w", err)
	}
	log.Println("Signature:", hex.EncodeToString(signature))
	input.Signature = hex.EncodeToString(signature)

	salt := []*big.Int{big.NewInt(0)}
	ikm := []*big.Int{&input.Key[0]}

	info := []*big.Int{compressedStateCurr[1], compressedStateNew[1]}

	res := hkdf.Hkdf(salt, ikm, info, 2)

	encrypted := gmimc.EncryptBig(compressedStateNew, res, gmimc.GetGMiMCRounds(len(compressedStateCurr)))
	for i, enc := range encrypted {
		input.Encrypted[i] = *enc
	}

	jsonBytes, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal inputs: %w", err)
	}

	err = os.WriteFile(inputFilePath, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write inputs: %w", err)
	}

	return nil
}

func generateSalt() *big.Int {
	// 32 random bytes
	random32 := make([]byte, 31)
	rand.Read(random32)
	//random32[31] &= 0x03 // 254 bits max

	salt := new(big.Int)
	salt.SetBytes(random32)

	return salt
}

func compressState(state State) []*big.Int {
	length := len(state.Variables) + len(state.Messages) + len(state.Balances) + 2

	uncompressedState := make([]*big.Int, len(state.States))

	for i, v := range state.States {
		uncompressedState[i] = big.NewInt(v.Int64())
	}

	if len(uncompressedState) > 63 {
		panic("Too many states to compress into a single field element.")
	}

	// Initialize the compressed field element
	var compressed *big.Int = big.NewInt(0)
	// Compress the states into a single field element
	for i := 0; i < len(uncompressedState); i++ {
		// Shift the compressed field element by 4 bits to make space for the next state
		compressed = new(big.Int).Mul(compressed, big.NewInt(16))
		// Add the next state to the compressed field element
		compressed = new(big.Int).Add(compressed, uncompressedState[i])
	}

	compressedState := make([]*big.Int, length)
	compressedState[0] = compressed
	compressedState[1], _ = new(big.Int).SetString(state.Radomness, 10)
	i := 2
	for _, v := range state.Variables {
		compressedState[i] = v
		i++
	}

	for _, m := range state.Messages {
		compressedState[i], _ = new(big.Int).SetString(m, 10)
		i++
	}

	for _, b := range state.Balances {
		compressedState[i], _ = new(big.Int).SetString(b, 10)
		i++
	}

	return compressedState
}

func Decompress(state *big.Int) []uint64 {
	// Initialize the decompressed state array
	var decompressed, result []uint64
	// Decompress the states from the field element

	for state.Cmp(big.NewInt(16)) >= 0 || state.Cmp(big.NewInt(0)) > 0 {
		// Extract the next state from the compressed field element
		decompressedBig := new(big.Int).Mod(state, big.NewInt(16))
		decompressed = append(decompressed, decompressedBig.Uint64())
		state = new(big.Int).Div(state, big.NewInt(16))
	}

	// Reverse the decompressed state array to get the original order of states
	for i := len(decompressed) - 1; i >= 0; i-- {
		result = append(result, decompressed[i])
	}

	return result
}
