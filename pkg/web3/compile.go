package web3

import (
	"github.com/lmittmann/go-solc"
)

func compileContract() ([]byte, error) {
	c := solc.New("0.8.21")

	contract, err := c.Compile("contracts", "Model", solc.WithOptimizer(&solc.Optimizer{Enabled: true, Runs: 999999}))
	if err != nil {
		return nil, err
	}

	return contract.Constructor, nil
}
