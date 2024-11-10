package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify <proof> <vk> <public-witness>",
	Short: "Verify a proof",
	Long:  `Verify a proof using the provided verification key (vk).`,
	Args:  cobra.ExactArgs(3),
	Run:   verifyCommand,
}

func init() {
	rootCMD.AddCommand(verifyCmd)
}

func verifyCommand(cmd *cobra.Command, args []string) {
	proofFile := args[0]
	vkFile := args[1]
	witness := args[2]

	err := zkp.VerifyProof(proofFile, vkFile, witness)
	if err != nil {
		log.Fatalln("Failed to verify proof:", err)
	}
}
