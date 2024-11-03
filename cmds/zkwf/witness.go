package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var witnessCommand = &cobra.Command{
	Use:   "witness <model> <input>",
	Short: "Generate a witness for a given input",
	Args:  cobra.ExactArgs(2),
	Run:   witnessCMDExecute,
}

func witnessCMDExecute(cmd *cobra.Command, args []string) {
	modelPath := args[0]
	inputPath := args[1]
	fullWitnessPath, _ := cmd.Flags().GetString("full")
	publicWitnessPath, _ := cmd.Flags().GetString("public")

	zkwf, err := zkp.NewZkWFProgram(modelPath)
	if err != nil {
		log.Fatalln("Failed to create zkWF program: ", err)
	}

	err = zkwf.ComputeWitness(inputPath, fullWitnessPath, publicWitnessPath)
	if err != nil {
		log.Fatalln("Failed to compute witness: ", err)
	}
}
