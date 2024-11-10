package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var compileCommand = &cobra.Command{
	Use:   "compile <model>",
	Short: "Compile a BPMN file into a zero-knowledge circuit",
	Args:  cobra.ExactArgs(1),
	Run:   compileCMDExecute,
}

func init() {
	compileCommand.PersistentFlags().StringP("output", "o", "circuit.r1cs", "Output file for the compiled circuit")
	rootCMD.AddCommand(compileCommand)
}

func compileCMDExecute(cmd *cobra.Command, args []string) {
	modelPath := args[0]
	outputFlag, _ := cmd.Flags().GetString("output")

	zkwf, err := zkp.NewZkWFProgram(modelPath)
	if err != nil {
		log.Fatalln("Failed to create zkWF program:", err)
	}

	err = zkwf.Compile(outputFlag)
	if err != nil {
		log.Fatalln("Failed to compile zkWF program:", err)
	}
}
