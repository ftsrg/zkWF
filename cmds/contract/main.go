package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	_ "embed"
)

//go:embed vk.bin
var vk []byte

type ModelContract struct {
	contractapi.Contract
}

// State represents the current hash and encrypted state
type State struct {
	Hash  string   `json:"hash"`
	State []string `json:"state"`
}

func (mc *ModelContract) Init(ctx contractapi.TransactionContextInterface, initialHash string, initialState []string) error {
	state := State{
		Hash:  initialHash,
		State: initialState,
	}
	stateJSON, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal initial state: %v", err)
	}
	return ctx.GetStub().PutState("modelState", stateJSON)
}

func (mc *ModelContract) GetCurrentHash(ctx contractapi.TransactionContextInterface) (string, error) {
	state, err := mc.getState(ctx)
	if err != nil {
		return "", err
	}
	return state.Hash, nil
}

func (mc *ModelContract) GetCurrentState(ctx contractapi.TransactionContextInterface) ([]string, error) {
	state, err := mc.getState(ctx)
	if err != nil {
		return nil, err
	}
	return state.State, nil
}

func (mc *ModelContract) Update(ctx contractapi.TransactionContextInterface, proof []byte, publicInputs []byte) error {
	state, err := mc.getState(ctx)
	if err != nil {
		return err
	}

	encryptedStateLen := len(state.State)
	if len(publicInputs) < 7+encryptedStateLen+2 {
		return fmt.Errorf("insufficient public inputs")
	}

	verified := Verify(proof, publicInputs)
	if !verified {
		return fmt.Errorf("verification failed")
	}

	stateJSON, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal updated state: %v", err)
	}

	if err := ctx.GetStub().PutState("modelState", stateJSON); err != nil {
		return fmt.Errorf("failed to save updated state: %v", err)
	}

	// Handle withdraw
	amount := publicInputs[len(publicInputs)-1]
	if amount > 0 {
		// Hyperledger Fabric does not handle native currency transfers,
		// so you may need to log this or integrate with external payment systems.
		fmt.Printf("Withdraw amount: %d\n", amount)
	}

	return nil
}

func (mc *ModelContract) getState(ctx contractapi.TransactionContextInterface) (*State, error) {
	stateJSON, err := ctx.GetStub().GetState("modelState")
	if err != nil {
		return nil, fmt.Errorf("failed to read state from world state: %v", err)
	}
	if stateJSON == nil {
		return nil, fmt.Errorf("state does not exist")
	}

	var state State
	if err := json.Unmarshal(stateJSON, &state); err != nil {
		return nil, fmt.Errorf("failed to unmarshal state JSON: %v", err)
	}
	return &state, nil
}

// Verify is a placeholder for the actual verification function you need to implement.
func Verify(proof []byte, witnessBytes []byte) bool {
	var byteRader io.Reader = bytes.NewReader(proof)
	var byteRader2 io.Reader = bytes.NewReader(witnessBytes)
	var byteRader3 io.Reader = bytes.NewReader(vk)

	proofObj := plonk.NewProof(ecc.BN254)

	_, err := proofObj.ReadFrom(byteRader)
	if err != nil {
		fmt.Printf("Error reading proof: %v", err)
		return false
	}

	witnessObj, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		fmt.Printf("Error creating witness: %v", err)
		return false
	}
	witnessObj.ReadFrom(byteRader2)

	vkObj := plonk.NewVerifyingKey(ecc.BN254)
	_, err = vkObj.ReadFrom(byteRader3)
	if err != nil {
		fmt.Printf("Error reading verification key: %v", err)
		return false
	}

	err = plonk.Verify(proofObj, vkObj, witnessObj)
	if err != nil {
		fmt.Printf("Error verifying proof: %v", err)
		return false
	}

	// Add your implementation here.
	return true
}

func main() {
	chaincode, err := contractapi.NewChaincode(&ModelContract{})
	if err != nil {
		fmt.Printf("Error creating ModelContract chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting ModelContract chaincode: %v", err)
	}
}
