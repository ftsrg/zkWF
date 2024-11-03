package main

import (
	"log"

	"github.com/ftsrg/zkWF/pkg/zkp"
	"github.com/spf13/cobra"
)

var setupCommand = &cobra.Command{
	Use:   "setup <circuit>",
	Short: "Setup a zero-knowledge circuit",
	Args:  cobra.ExactArgs(1),
	Run:   setupCMDExecute,
}

func setupCMDExecute(cmd *cobra.Command, args []string) {
	circuitPath := args[0]
	pkPath, _ := cmd.Flags().GetString("pk")
	vkPath, _ := cmd.Flags().GetString("vk")
	contract, _ := cmd.Flags().GetString("contract")

	err := zkp.Setup(circuitPath, vkPath, pkPath, contract)
	if err != nil {
		log.Fatalln("Failed to setup zkWF program: ", err)
	}
}
