package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var witnessCommand = &cobra.Command{
	Use:   "witness <model> <input> <key>",
	Short: "Generate a witness for a given input",
	Args:  cobra.ExactArgs(3),
	Run:   witnessCMDExecute,
}

func init() {
	witnessCommand.PersistentFlags().StringP("public", "p", "public.wtns", "Output file for the public witness")
	witnessCommand.PersistentFlags().StringP("full", "f", "full.wtns", "Output file for the full witness")
	rootCMD.AddCommand(witnessCommand)
}

func witnessCMDExecute(cmd *cobra.Command, args []string) {
	modelPath := args[0]
	inputPath := args[1]
	keyPath := args[2]
	fullWitnessPath, _ := cmd.Flags().GetString("full")
	publicWitnessPath, _ := cmd.Flags().GetString("public")

	zkwf, err := zkp.NewZkWFProgram(modelPath)
	if err != nil {
		log.Fatalln("Failed to create zkWF program: ", err)
	}

	err = zkwf.ComputeWitness(inputPath, keyPath, fullWitnessPath, publicWitnessPath)
	if err != nil {
		log.Fatalln("Failed to compute witness: ", err)
	}
}
