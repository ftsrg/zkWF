package powersoftau

import (
	"fmt"
	"log"
	"math"
	"os"

	kzg_bn254 "github.com/consensys/gnark-crypto/ecc/bn254/kzg"
	"github.com/consensys/gnark-crypto/kzg"
	gnark_ptau "github.com/mdehoog/gnark-ptau"
)

func GetPowerOfTauParams(n uint64) (kzg.SRS, kzg.SRS, error) {
	sqrtn := math.Log2(float64(n))
	//sqrtn = 16
	fmt.Printf("sqrtn: %v\n", sqrtn)

	filename := fmt.Sprintf("powersOfTau28_hez_final_%v.ptau", sqrtn)
	_, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error checking existance of %s: %v\n"+
			"Refer to doc.go for instructions on how to download the file.",
			filename, err)
		return nil, nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening %s: %v", filename, err)
	}

	srs, err := gnark_ptau.ToSRS(file)
	if err != nil {
		log.Fatalf("error converting to SRS: %v", err)
	}
	var canonicalSRS kzg.SRS
	canonicalSRS = srs

	newSRS := &kzg_bn254.SRS{Vk: srs.Vk}
	/*
		// instead of using ToLagrangeG1 we can directly do a fft on the powers of alpha
		// since we know the randomness in test.
		pAlpha := make([]fr_bn254.Element, size)
		pAlpha[0].SetUint64(1)
		pAlpha[1].SetBigInt(tau)
		for i := 2; i < len(pAlpha); i++ {
			pAlpha[i].Mul(&pAlpha[i-1], &pAlpha[1])
		}
		// do a fft on this.
		d := fft_bn254.NewDomain(size)
		d.FFTInverse(pAlpha, fft_bn254.DIF)
		fft_bn254.BitReverse(pAlpha)

		// bath scalar mul
		_, _, g1gen, _ := bn254.Generators()
		newSRS.Pk.G1 = bn254.BatchScalarMultiplicationG1(&g1gen, pAlpha)*/
	newSRS.Pk.G1, err = kzg_bn254.ToLagrangeG1(srs.Pk.G1[:n])
	if err != nil {
		log.Fatalf("error converting to Lagrange G1: %v", err)
	}
	return canonicalSRS, newSRS, nil
}
