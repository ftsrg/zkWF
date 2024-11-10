package web3

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ftsrg/zkWF/pkg/contracts/model"
	"github.com/ftsrg/zkWF/pkg/crypto/gmimc"
	"github.com/ftsrg/zkWF/pkg/crypto/hkdf"
	"github.com/ftsrg/zkWF/pkg/zkp"
)

func GetState(url, keyPath string, chainID *big.Int, contractAddress string, encKeyPrev, r1, r2 *big.Int) (string, error) {
	client, err := CreateConnection(url)
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}

	contract, err := model.NewModel(common.HexToAddress(contractAddress), client)
	if err != nil {
		return "", fmt.Errorf("failed to instantiate contract: %w", err)
	}

	encrypted, err := contract.GetCurrentState(&bind.CallOpts{})
	if err != nil {
		return "", fmt.Errorf("failed to get current state: %w", err)
	}

	salt := []*big.Int{big.NewInt(0)}
	ikm := []*big.Int{encKeyPrev}

	info := []*big.Int{r1, r2}

	res := hkdf.Hkdf(salt, ikm, info, 2)

	decrypted := gmimc.DecryptBig(encrypted, res, gmimc.GetGMiMCRounds(len(encrypted)))
	for _, v := range decrypted {
		fmt.Println("decrypted:", v)
	}

	decompressed := zkp.Decompress(decrypted[0])
	for _, v := range decompressed {
		fmt.Println("decompressed:", v)
	}

	return "", nil
}
