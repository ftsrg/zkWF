package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var proveCommand = &cobra.Command{
	Use:   "prove <circuit> <prover-key> <full-witness>",
	Short: "Prove a statement using a given circuit and witness",
	Args:  cobra.ExactArgs(3),
	Run:   proveCMDExecute,
}

func init() {
	proveCommand.PersistentFlags().StringP("output", "o", "proof.bin", "Output file for the proof")
	rootCMD.AddCommand(proveCommand)
}

func proveCMDExecute(cmd *cobra.Command, args []string) {
	circuitPath := args[0]
	pkPath := args[1]
	fullWitnessPath := args[2]
	proofPath, _ := cmd.Flags().GetString("output")

	err := zkp.Prove(circuitPath, pkPath, fullWitnessPath, proofPath)
	if err != nil {
		log.Fatalln("Failed to prove statement: ", err)
	}
}
