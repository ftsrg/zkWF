package main

import (
	"fmt"
	"log"

	"github.com/ftsrg/zkWF/pkg/crypto/keys"
	"github.com/spf13/cobra"
)

var generateKeyCommand = &cobra.Command{
	Use:   "generate-key",
	Short: "Generate a new eddsa key pair",
	RunE:  generateKeyCommandFunc,
}

func init() {
	generateKeyCommand.PersistentFlags().StringP("output", "o", "key.json", "Output file path")
	rootCMD.AddCommand(generateKeyCommand)
}

func generateKeyCommandFunc(cmd *cobra.Command, args []string) error {
	outputFile, _ := cmd.Flags().GetString("output")

	err := keys.GenerateKeyPair(outputFile)
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	log.Println("Key pair generated and saved to", outputFile)
	return nil
}
