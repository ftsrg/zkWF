package zkp

import (
	"fmt"

	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/scs"
)

func (zkwf *ZkWFProgram) Compile(output string) error {

	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, zkwf.Circuit)
	if err != nil {
		return fmt.Errorf("circuit compilation error: %w", err)
	}

	file, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error creating circuit file: %w", err)
	}
	defer file.Close()

	_, err = ccs.WriteTo(file)
	if err != nil {
		return fmt.Errorf("error writing circuit to file: %w", err)
	}

	return nil
}
