package web3

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/web3"
	"github.com/spf13/cobra"
)

var stepCmd = &cobra.Command{
	Use:   "step <contract> <proof> <witness>",
	Short: "Execute a single step of the proof",
	Long:  `Execute a single step of the proof`,
	Args:  cobra.ExactArgs(3),
	Run:   stepCmdFunc,
}

func init() {
	Web3Command.AddCommand(stepCmd)
}

func stepCmdFunc(cmd *cobra.Command, args []string) {
	nodeUrl, keyPath, chainId := getFlags()
	contract, proof, witness := args[0], args[1], args[2]

	err := web3.StepModel(nodeUrl, keyPath, contract, chainId, proof, witness)
	if err != nil {
		log.Fatalln("Failed to step model:", err)
	}
}
