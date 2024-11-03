package main

import (
	"container/list"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ftsrg/zkWF/pkg/model"
)

var rootCMD = &cobra.Command{
	Use:   "zkwf",
	Short: "zkWF is a zero-knowledge workflow system",
}

func init() {
	compileCommand.PersistentFlags().StringP("output", "o", "circuit.r1cs", "Output file for the compiled circuit")
	rootCMD.AddCommand(compileCommand)

	witnessCommand.PersistentFlags().StringP("public", "p", "public.wtns", "Output file for the public witness")
	witnessCommand.PersistentFlags().StringP("full", "f", "full.wtns", "Output file for the full witness")
	rootCMD.AddCommand(witnessCommand)

	setupCommand.PersistentFlags().StringP("pk", "p", "pk.bin", "Output file for the proving key")
	setupCommand.PersistentFlags().StringP("vk", "v", "vk.bin", "Output file for the verification key")
	setupCommand.PersistentFlags().StringP("contract", "c", "contract.sol", "Output file for the Solidity contract")
	rootCMD.AddCommand(setupCommand)

	proveCommand.PersistentFlags().StringP("output", "o", "proof.bin", "Output file for the proof")
	rootCMD.AddCommand(proveCommand)
}

// BFS traverses the BPMN graph starting from a StartEvent using Breadth-First Search
func BFS(graph *model.BPMNGraph, startNodeID string) {
	// Initialize a queue (FIFO) for BFS and a visited map
	queue := list.New()
	visited := make(map[string]bool)

	// Enqueue the start node
	queue.PushBack(startNodeID)
	visited[startNodeID] = true

	// Traverse the graph
	for queue.Len() > 0 {
		// Dequeue the next node
		element := queue.Front()
		nodeID := element.Value.(string)
		queue.Remove(element)

		// Get the node from the graph
		node := graph.Nodes[nodeID]
		fmt.Printf("Visited Node: %s (Type: %s, Name: %s)\n", node.ID, node.Type, node.Name)
		next := node.GetNextNodes()
		if len(next) > 0 {
			fmt.Print("\tNext nodes: ")
			for _, n := range next {
				fmt.Print(n.ID, " ")
			}
			fmt.Println()
		}

		perv := node.GetPairs()
		if len(perv) > 0 {
			fmt.Print("\tPair nodes: ")
			for _, n := range perv {
				fmt.Print(n.ID, " ")
			}
			fmt.Println()
		}

		// Enqueue all outgoing edges' target nodes (if not visited)
		for _, edge := range node.Outgoing {
			targetID := edge.TargetRef.ID
			if !visited[targetID] {
				queue.PushBack(targetID)
				visited[targetID] = true
			}
		}
	}
}

func main() {
	rootCMD.Execute()
	/*
	   //xmlBytes, _ := os.ReadFile("/home/balazs/Letöltések/diagram(2).bpmn")
	   xmlBytes, _ := os.ReadFile("/home/balazs/Code/Kotlin Projects/zkWF/models/leasing-payment/leasing-payment.bpmn")

	   var definitions bpmn.Definitions
	   err := xml.Unmarshal(xmlBytes, &definitions)

	   	if err != nil {
	   		fmt.Println("Error parsing XML:", err)
	   		os.Exit(1)
	   	}

	   // Build the graph from the parsed BPMN definitions
	   graph := model.BuildGraph(&definitions)
	   startNodes := graph.GetStartNodes()

	   	for _, node := range startNodes {
	   		fmt.Printf("Start Node: %s (Type: %s)\n", node.ID, node.Type)
	   	}

	   BFS(graph, startNodes[0].ID)

	   executables := graph.GetExecutableNodes()

	   // Print the executable nodes
	   fmt.Println("Executable Nodes:")

	   	for i, node := range executables {
	   		fmt.Printf("%d. Node: %s (Type: %s, Name: %s)\n", i, node.ID, node.Type, node.Name)
	   		// Print outgoing edges for each node
	   		for _, edge := range node.Outgoing {
	   			fmt.Printf("\tEdge: %s -> %s\n", edge.SourceRef.ID, edge.TargetRef.ID)
	   		}

	   }

	   var circuit statechecker.Circuit
	   circuit.Model = graph
	   circuit.State_curr.States = make([]frontend.Variable, len(executables))
	   circuit.State_new.States = make([]frontend.Variable, len(executables))
	   circuit.State_curr.Variables = make(map[string]frontend.Variable, 0)
	   circuit.State_new.Variables = make(map[string]frontend.Variable, 0)
	   x := []frontend.Variable{0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	   	if len(x) != len(executables) {
	   		log.Fatalln("Invalid state length")
	   	}

	   y := []frontend.Variable{0, 0, 0, 10, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	   	if len(y) != len(executables) {
	   		log.Fatalln("Invalid state length")
	   	}

	   var w statechecker.Circuit
	   var state1, state2 statechecker.State

	   	state1 = statechecker.State{
	   		States:    make([]frontend.Variable, len(executables)),
	   		Variables: make(map[string]frontend.Variable, 1),
	   		Radomness: 0,
	   	}

	   copy(state1.States, x)
	   state1.Variables["a"] = 0

	   w.State_curr = state1

	   	state2 = statechecker.State{
	   		States:    make([]frontend.Variable, len(executables)),
	   		Variables: make(map[string]frontend.Variable, 1),
	   		Radomness: 1,
	   	}

	   copy(state2.States, y)
	   state2.Variables["a"] = 2

	   w.State_new = state2

	   w.HashCurr, _ = big.NewInt(0).SetString("10623130405679805330576346668288554504439363553729204031500032016204369167485", 10)
	   w.HashNew, _ = big.NewInt(0).SetString("17735585319247439581380290770726239898495846102644983442830724852587043806858", 10)
	   w.Model = graph

	   fmt.Printf("%v ==> %v?\n\n", x, y)

	   log.Println("Compiling circuit")
	   ccs, err := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, &circuit)

	   	if err != nil {
	   		fmt.Println("circuit compilation error")
	   	}

	   sizeSystem := ccs.GetNbPublicVariables() + ccs.GetNbConstraints()
	   fmt.Println("Size of the system: ", sizeSystem)

	   nextPowerTwo := ecc.NextPowerOfTwo(uint64(sizeSystem))
	   fmt.Println("Next power of two: ", nextPowerTwo)

	   log.Println("Generating powers of tau params")
	   srs, srsLagrange, err := powersoftau.GetPowerOfTauParams(nextPowerTwo)

	   	if err != nil {
	   		log.Fatalln("Failed to get powers of tau params: ", err)
	   	}

	   log.Println("Creating witness")
	   witnessFull, err := frontend.NewWitness(&w, ecc.BN254.ScalarField())

	   	if err != nil {
	   		log.Fatalln("Failed to create witness: ", err)
	   	}

	   log.Println("Public witness")
	   witnessPublic, err := frontend.NewWitness(&w, ecc.BN254.ScalarField(), frontend.PublicOnly())

	   	if err != nil {
	   		log.Fatalln("Failed to create public witness: ", err)
	   	}

	   log.Println("Setup Plonk")
	   pk, vk, err := plonk.Setup(ccs, srs, srsLagrange)
	   //_, err := plonk.Setup(r1cs, kate, &publicWitness)

	   	if err != nil {
	   		log.Fatalln("Failed to setup Plonk: ", err)
	   	}

	   log.Println("Plonk Prove")

	   proof, err := plonk.Prove(ccs, pk, witnessFull)

	   	if err != nil {
	   		log.Fatalln("Failed to prove: ", err)
	   	}

	   log.Println("Plonk Verify")
	   err = plonk.Verify(proof, vk, witnessPublic)

	   	if err != nil {
	   		fmt.Printf("Invalid proof")
	   	} else {

	   		fmt.Printf("Valid proof")
	   	}
	*/
}
