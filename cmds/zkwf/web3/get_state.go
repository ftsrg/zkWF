package web3

import (
	"log"
	"math/big"

	"github.com/ftsrg/zkWF/pkg/web3"
	"github.com/spf13/cobra"
)

var getStateCmd = &cobra.Command{
	Use:   "get-state <contract> <key> <r1> <r2>",
	Args:  cobra.ExactArgs(4),
	Short: "Get the state of a workflow",
	Run:   getStateCmdFunc,
}

func init() {
	Web3Command.AddCommand(getStateCmd)
}

func getStateCmdFunc(cmd *cobra.Command, args []string) {
	nodeUrl, keyPath, chainId := getFlags()
	contract, prevEncKeyStr, r1Str, r2Str := args[0], args[1], args[2], args[3]
	prevEncKey, _ := new(big.Int).SetString(prevEncKeyStr, 10)
	r1, _ := new(big.Int).SetString(r1Str, 10)
	r2, _ := new(big.Int).SetString(r2Str, 10)

	_, err := web3.GetState(nodeUrl, keyPath, chainId, contract, prevEncKey, r1, r2)
	if err != nil {
		log.Fatalln("Failed to get state:", err)
	}
}
