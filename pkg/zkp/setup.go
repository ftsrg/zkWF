package zkp

import (
	"fmt"
	"log"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/constraint"
	"github.com/ftsrg/zkWF/pkg/powersoftau"
)

func loadProgram(r1csPath string) (constraint.ConstraintSystem, error) {
	file, err := os.Open(r1csPath)
	if err != nil {
		return nil, fmt.Errorf("error opening circuit file: %w", err)
	}
	defer file.Close()

	var ccs constraint.ConstraintSystem = plonk.NewCS(ecc.BN254)

	_, err = ccs.ReadFrom(file)
	if err != nil {
		return nil, fmt.Errorf("error reading circuit file: %w", err)
	}

	return ccs, nil
}

func Setup(r1csPath, vkPath, pkPath, contractPath string) error {
	ccs, err := loadProgram(r1csPath)
	if err != nil {
		return fmt.Errorf("error loading circuit: %w", err)
	}

	sizeSystem := ccs.GetNbPublicVariables() + ccs.GetNbConstraints()
	fmt.Println("Size of the system: ", sizeSystem)

	nextPowerTwo := ecc.NextPowerOfTwo(uint64(sizeSystem))
	fmt.Println("Next power of two: ", nextPowerTwo)

	log.Println("Generating powers of tau params")
	srs, srsLagrange, err := powersoftau.GetPowerOfTauParams(nextPowerTwo)

	if err != nil {
		return fmt.Errorf("failed to get powers of tau params: %w", err)
	}

	log.Println("Setup Plonk")
	pk, vk, err := plonk.Setup(ccs, srs, srsLagrange)

	if err != nil {
		return fmt.Errorf("failed to setup Plonk: %w", err)
	}

	log.Println("Writing verification key to file")
	vkFile, err := os.Create(vkPath)
	if err != nil {
		return fmt.Errorf("error creating verification key file: %w", err)
	}
	defer vkFile.Close()

	_, err = vk.WriteTo(vkFile)
	if err != nil {
		return fmt.Errorf("error writing verification key to file: %w", err)
	}

	log.Println("Writing proving key to file")

	pkFile, err := os.Create(pkPath)
	if err != nil {
		return fmt.Errorf("error creating proving key file: %w", err)
	}

	defer pkFile.Close()

	_, err = pk.WriteTo(pkFile)
	if err != nil {
		return fmt.Errorf("error writing proving key to file: %w", err)
	}

	log.Println("Setup completed")

	contractFile, err := os.Create(contractPath)
	if err != nil {
		return fmt.Errorf("error creating contract file: %w", err)
	}
	defer contractFile.Close()

	if err := vk.ExportSolidity(contractFile); err != nil {
		return fmt.Errorf("error exporting contract: %w", err)
	}

	return nil
}
