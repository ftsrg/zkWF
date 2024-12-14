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
	Variables map[string]*big.Int
	Messages  map[string]string
	Balances  map[string]string
	Radomness string
}

type Inputs struct {
	State_curr State
	State_new  State
	HashCurr   string
	HashNew    string
	Signature  string
	Encrypted  []big.Int
	Key        []big.Int
	Deposit    big.Int
	Withdraw   big.Int
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
	circuit.State_curr.Variables = make([]frontend.Variable, len(graph.Variables))
	circuit.State_new.Variables = make([]frontend.Variable, len(graph.Variables))
	circuit.State_curr.Messages = make([]frontend.Variable, graph.MessageCount)
	circuit.State_new.Messages = make([]frontend.Variable, graph.MessageCount)
	circuit.State_curr.Balances = make([]frontend.Variable, len(graph.Participants))
	circuit.State_new.Balances = make([]frontend.Variable, len(graph.Participants))
	ecryptionLen := 2 + len(graph.Variables) + graph.MessageCount + len(graph.Participants)
	/*circuit.Keys.PrivateKey = 0
	circuit.Keys.PublicKey.Assign(twistededwards.BN254, make([]byte, 32))*/
	log.Println("Encryption length: ", ecryptionLen)
	/*circuit.Encrypted = make([]frontend.Variable, ecryptionLen)
	circuit.Key = make([]frontend.Variable, ecryptionLen/2)*/

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
	w.Variables = make([]frontend.Variable, len(state.Variables))
	w.Messages = make([]frontend.Variable, len(state.Messages))
	w.Balances = make([]frontend.Variable, len(state.Balances))

	for i, bigInt := range state.States {
		w.States[i] = bigInt
	}

	i := 0
	for _, bigInt := range state.Variables {
		w.Variables[i] = bigInt
		i++
	}

	i = 0
	for _, bigInt := range state.Messages {
		w.Messages[i], _ = big.NewInt(0).SetString(bigInt, 10)
		i++
	}

	i = 0
	for _, bigInt := range state.Balances {
		w.Balances[i], _ = big.NewInt(0).SetString(bigInt, 10)
		i++
	}

	w.Radomness, _ = big.NewInt(0).SetString(state.Radomness, 10)

	return w
}
