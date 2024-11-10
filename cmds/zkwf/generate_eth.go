package main

import (
	"fmt"
	"log"

	"github.com/ftsrg/zkWF/pkg/crypto/keys"
	"github.com/spf13/cobra"
)

var generateEthKeyCommand = &cobra.Command{
	Use:   "generate-eth-key",
	Short: "Generate a new ecdsa key pair for Ethereum",
	RunE:  generateEthKeyCommandFunc,
}

func init() {
	generateEthKeyCommand.PersistentFlags().StringP("output", "o", "eth_key.json", "Output file path")
	rootCMD.AddCommand(generateEthKeyCommand)
}

func generateEthKeyCommandFunc(cmd *cobra.Command, args []string) error {
	outputFile, _ := cmd.Flags().GetString("output")

	err := keys.GenerateKeyEthPair(outputFile)
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	log.Println("Key pair generated and saved to", outputFile)
	return nil
}
