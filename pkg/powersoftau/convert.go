package powersoftau

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"

	kzg_bn254 "github.com/consensys/gnark-crypto/ecc/bn254/kzg"
	"github.com/consensys/gnark-crypto/kzg"
	gnark_ptau "github.com/mdehoog/gnark-ptau"
)

func downloadFile(power int) error {

	tau := powersoftau[power]

	file, err := os.Create(fmt.Sprintf("powersOfTau28_hez_final_%v.ptau", power))
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	resp, err := http.Get(tau.URL)
	if err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}

	return nil
}

func GetPowerOfTauParams(n uint64) (kzg.SRS, kzg.SRS, error) {
	sqrtn := math.Log2(float64(n))
	//sqrtn = 16
	fmt.Printf("sqrtn: %v\n", sqrtn)

	filename := fmt.Sprintf("powersOfTau28_hez_final_%v.ptau", sqrtn)
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		err := downloadFile(int(sqrtn))
		if err != nil {
			log.Fatalf("error downloading file: %v", err)
		}
	} else if err != nil {
		log.Fatalf("error checking if %s exists: %v", filename, err)
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
