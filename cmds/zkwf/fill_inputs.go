package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var fillInputsCommand = &cobra.Command{
	Use:   "fill-inputs <input-file> <keys-file>",
	Short: "Fill inputs for the zkWF circuit",
	Long: `Fill inputs for the zkWF circuit. It reads the input and keys from the given files and fills the input file with the appropriate values as follows:
	- Generates a randomness for both states
	- Fills in the Hash fields
	- Signs the hash with the given keys
	- Fills in the encryption fields
	`,
	Args: cobra.ExactArgs(2),
	Run:  fillInputsCmdFunc,
}

func init() {
	rootCMD.AddCommand(fillInputsCommand)
}

func fillInputsCmdFunc(cmd *cobra.Command, args []string) {
	inputFile := args[0]
	keysFile := args[1]

	err := zkp.FillInputs(inputFile, keysFile)
	if err != nil {
		log.Fatalln("Failed to fill inputs:", err)
	}
}
