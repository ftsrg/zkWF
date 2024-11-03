package zkp

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/statechecker"
	"github.com/ftsrg/zkWF/pkg/crypto/keys"
)

func (zkwf *ZkWFProgram) ComputeWitness(inputPath, witnessFullPath, publicWitnessPath string) error {
	inputs, err := loadInputs(inputPath)
	if err != nil {
		return fmt.Errorf("error loading inputs: %w", err)
	}

	var w statechecker.Circuit

	w.State_curr = inputs.State_curr.toVariableState()
	w.State_new = inputs.State_new.toVariableState()

	w.HashCurr, _ = big.NewInt(0).SetString(inputs.HashCurr, 10)
	w.HashNew, _ = big.NewInt(0).SetString(inputs.HashNew, 10)
	pubKey, err := keys.HexToPublicKey(inputs.PublicKey)
	if err != nil {
		return fmt.Errorf("error parsing public key: %w", err)
	}
	w.PublicKey.Assign(twistededwards.ID(ecc.BN254), pubKey.Bytes()[:32])

	w.Encrypted = make([]frontend.Variable, len(inputs.Encrypted))
	for i, e := range inputs.Encrypted {
		w.Encrypted[i] = e
	}

	w.Key = make([]frontend.Variable, len(inputs.Key))
	for i, k := range inputs.Key {
		w.Key[i] = k
	}

	sigBytes, err := hex.DecodeString(inputs.Signature)
	if err != nil {
		return fmt.Errorf("error decoding signature: %w", err)
	}
	w.Signature.Assign(twistededwards.ID(ecc.BN254), sigBytes)

	w.Model = zkwf.Model

	log.Println("Creating witness")
	witnessFull, err := frontend.NewWitness(&w, ecc.BN254.ScalarField())

	if err != nil {
		log.Fatalln("Failed to create witness: ", err)
	}

	log.Println("Writing witness to file")
	file, err := os.Create(witnessFullPath)
	if err != nil {
		return fmt.Errorf("error creating witness file: %w", err)
	}

	_, err = witnessFull.WriteTo(file)
	if err != nil {
		return fmt.Errorf("error writing witness to file: %w", err)
	}

	log.Println("Public witness")
	witnessPublic, err := frontend.NewWitness(&w, ecc.BN254.ScalarField(), frontend.PublicOnly())
	if err != nil {
		log.Fatalln("Failed to create public witness: ", err)
	}

	log.Println("Writing public witness to file")

	file, err = os.Create(publicWitnessPath)
	if err != nil {
		return fmt.Errorf("error creating public witness file: %w", err)
	}

	_, err = witnessPublic.WriteTo(file)
	if err != nil {
		return fmt.Errorf("error writing public witness to file: %w", err)
	}

	vector := witnessPublic.Vector().(fr.Vector)

	// Print public inputs
	// Public inputs: HashCurr, HashNew, PublicKey, Signature, Encrypted
	fmt.Println("Public inputs:", vector)
	return nil
}
