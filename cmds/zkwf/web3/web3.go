package web3

import (
	"math/big"

	"github.com/spf13/cobra"
)

var Web3Command = &cobra.Command{
	Use:   "web3",
	Short: "Command for interacting with the web3 API",
}

func init() {
	Web3Command.PersistentFlags().StringP("eth-node", "n", "http://127.0.0.1:8545", "Ethereum node URL")
	Web3Command.PersistentFlags().StringP("key-path", "k", "eth_key.json", "Path to the key file")
	Web3Command.PersistentFlags().StringP("chain-id", "c", "31337", "Chain ID of the Ethereum network") // default to hardhat localhost
}

func getFlags() (ethNodeURL string, keyPath string, chainID *big.Int) {
	ethNodeURL, _ = Web3Command.PersistentFlags().GetString("eth-node")
	keyPath, _ = Web3Command.PersistentFlags().GetString("key-path")
	chainIDStr, _ := Web3Command.PersistentFlags().GetString("chain-id")

	chainID = new(big.Int)
	chainID.SetString(chainIDStr, 10)

	return ethNodeURL, keyPath, chainID
}
