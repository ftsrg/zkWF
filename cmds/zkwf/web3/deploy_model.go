package web3

import (
	"math/big"

	"github.com/ftsrg/zkWF/pkg/web3"
	"github.com/spf13/cobra"
)

var deployModelCommand = &cobra.Command{
	Use:   "deploy-model",
	Args:  cobra.MinimumNArgs(3),
	Short: "Deploy the State manager contract to Ethereum",
	RunE:  deployModelCommandFunc,
}

func init() {

	Web3Command.AddCommand(deployModelCommand)
}

func deployModelCommandFunc(cmd *cobra.Command, args []string) error {
	nodeUrl, keyPath, chainId := getFlags()
	var initialHash *big.Int = new(big.Int)
	var initialState []*big.Int = make([]*big.Int, len(args)-1)

	initialHash.SetString(args[0], 10)
	for i := 1; i < len(args); i++ {
		initialState[i-1] = new(big.Int)
		initialState[i-1].SetString(args[i], 10)
	}

	_, err := web3.DeployModelContract(nodeUrl, keyPath, chainId, initialHash, initialState)
	if err != nil {
		return err
	}

	return nil
}
