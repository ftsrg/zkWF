package zkp

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/statechecker"
	"github.com/ftsrg/zkWF/pkg/model"
)

type ZkWFProgram struct {
	Model            *model.BPMNGraph
	Circuit          *statechecker.Circuit
	constraintSystem constraint.ConstraintSystem
}

type State struct {
	States    []big.Int
	Variables map[string]big.Int
	Messages  map[string]big.Int
	Balances  map[string]string
	Radomness string
}

type Inputs struct {
	State_curr State
	State_new  State
	HashCurr   string
	HashNew    string
	PublicKey  string
	Signature  string
	Encrypted  []big.Int
	Key        []big.Int
}

func NewZkWFProgram(modelPath string) (*ZkWFProgram, error) {
	graph, err := loadFile(modelPath)
	if err != nil {
		return nil, fmt.Errorf("error loading model graph: %w", err)
	}

	executables := graph.GetExecutableNodes()
	log.Printf("Executable nodes: %v\n", len(executables))

	var circuit statechecker.Circuit
	circuit.Model = graph
	circuit.State_curr.States = make([]frontend.Variable, len(executables))
	circuit.State_new.States = make([]frontend.Variable, len(executables))
	circuit.State_curr.Variables = make(map[string]frontend.Variable, len(graph.Variables))
	circuit.State_new.Variables = make(map[string]frontend.Variable, len(graph.Variables))
	circuit.State_curr.Messages = make(map[string]frontend.Variable, graph.MessageCount)
	circuit.State_new.Messages = make(map[string]frontend.Variable, graph.MessageCount)
	circuit.State_curr.Balances = make(map[string]frontend.Variable, graph.ParticpnatCount)
	circuit.State_new.Balances = make(map[string]frontend.Variable, graph.ParticpnatCount)
	ecryptionLen := 2 + len(graph.Variables) + graph.MessageCount + graph.ParticpnatCount
	log.Println("Encryption length: ", ecryptionLen)
	circuit.Encrypted = make([]frontend.Variable, ecryptionLen)
	circuit.Key = make([]frontend.Variable, ecryptionLen/2)

	for _, key := range graph.MessageMap {
		circuit.State_curr.Messages[key] = 0
		circuit.State_new.Messages[key] = 0
	}

	for _, key := range graph.Variables {
		circuit.State_curr.Variables[key] = 0
		circuit.State_new.Variables[key] = 0
	}

	for i := 0; i < graph.ParticpnatCount; i++ {
		circuit.State_curr.Balances[fmt.Sprintf("p%d", i)] = 0
		circuit.State_new.Balances[fmt.Sprintf("p%d", i)] = 0
	}

	return &ZkWFProgram{
		Model:   graph,
		Circuit: &circuit,
	}, nil
}

func (zkwf *ZkWFProgram) LoadCompiled(path string) error {
	// Load the compiled circuit from the given path
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening circuit file: %w", err)
	}

	var css constraint.ConstraintSystem
	_, err = css.ReadFrom(file)
	if err != nil {
		return fmt.Errorf("error reading circuit from file: %w", err)
	}
	zkwf.constraintSystem = css

	return nil
}

func (state State) toVariableState() statechecker.State {
	var w statechecker.State
	w.States = make([]frontend.Variable, len(state.States))
	w.Variables = make(map[string]frontend.Variable, len(state.Variables))
	w.Messages = make(map[string]frontend.Variable, len(state.Messages))
	w.Balances = make(map[string]frontend.Variable, len(state.Balances))

	for i, bigInt := range state.States {
		w.States[i] = bigInt
	}

	for key, bigInt := range state.Variables {
		w.Variables[key] = bigInt
	}

	for key, bigInt := range state.Messages {
		w.Messages[key] = bigInt
	}

	for key, bigInt := range state.Balances {
		w.Balances[key], _ = big.NewInt(0).SetString(bigInt, 10)
	}

	w.Radomness, _ = big.NewInt(0).SetString(state.Radomness, 10)

	return w
}
