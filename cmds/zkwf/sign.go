package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var signCommand = &cobra.Command{
	Use:   "sign <keys> <input>",
	Short: "Sign a given input",
	Args:  cobra.ExactArgs(2),
	Run:   signCMDExecute,
}

func init() {
	rootCMD.AddCommand(signCommand)
}

func signCMDExecute(cmd *cobra.Command, args []string) {
	keysPath := args[0]
	inputPath := args[1]

	err := zkp.SignHash(keysPath, inputPath)
	if err != nil {
		log.Fatalln("Failed to sign: ", err)
	}
}
