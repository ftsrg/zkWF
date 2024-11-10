package web3

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/web3"
	"github.com/spf13/cobra"
)

var deployEcdhCommand = &cobra.Command{
	Use:   "deploy-ecdh",
	Args:  cobra.ExactArgs(0),
	Short: "Deploy the ECDH contract with predefined public keys",
	RunE:  deployEcdhCommandFunc,
}

func init() {
	deployEcdhCommand.Flags().StringSliceP("public-keys", "p", []string{}, "Public keys to be deployed")
	Web3Command.AddCommand(deployEcdhCommand)
}

func deployEcdhCommandFunc(cmd *cobra.Command, args []string) error {

	publicKeys, _ := cmd.Flags().GetStringSlice("public-keys")
	nodeUrl, keyPath, chainId := getFlags()

	if len(publicKeys) == 0 {
		log.Fatalln("No public keys provided")
	}

	address, err := web3.DeployEcdhContract(nodeUrl, keyPath, chainId, publicKeys)
	if err != nil {
		return err
	}

	cmd.Printf("Contract deployed! Address: %s\n", address)
	return nil
}
