package proofofownership

import (
	"fmt"
	"math/big"

	ecc_twisted "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/std/signature/eddsa"
)

// ProofOfOwnership is a proof of ownership
func ProofOfOwnership(api frontend.API, publicKey eddsa.PublicKey, privKey [2]frontend.Variable) error {
	pointA := publicKey.A
	pow_2_128 := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil) // 2^128

	edCurve, err := twistededwards.NewEdCurve(api, ecc_twisted.BN254) // If you need to change the curve, change this line
	if err != nil {
		return fmt.Errorf("failed to create twisted edwards curve: %v", err)
	}

	base := twistededwards.Point{
		X: edCurve.Params().Base[0],
		Y: edCurve.Params().Base[1],
	}
	var pointB, chunk1, chunk2 twistededwards.Point // The private key is split into two parts, as it is 32 bytes long, just a few bits short of the field size. This is why we need to split it into two 128-bit chunks. The first chunk is multiplied by 2^128 to ensure that it is in the correct range. Then, the two chunks are added together to get the public key.
	chunk1 = edCurve.ScalarMul(base, privKey[0])
	chunk2 = edCurve.ScalarMul(base, privKey[1])

	chunk1 = edCurve.ScalarMul(chunk1, pow_2_128)

	pointB = edCurve.Add(chunk1, chunk2)

	edCurve.AssertIsOnCurve(pointB)

	api.Println("X coordinates are: ", pointA.X, pointB.X)
	api.AssertIsEqual(pointA.X, pointB.X)
	api.Println("Y coordinates are: ", pointA.Y, pointB.Y)
	api.AssertIsEqual(pointA.Y, pointB.Y)

	return nil
}
