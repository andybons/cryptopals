package cryptopals

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"math/big"
)

func encryptionOracle(b []byte) ([]byte, error) {
	pfx := make([]byte, randInt(5, 11))
	if _, err := rand.Read(pfx); err != nil {
		return nil, err
	}
	sfx := make([]byte, randInt(5, 10))
	if _, err := rand.Read(sfx); err != nil {
		return nil, err
	}

	padded := append(pfx, b...)
	padded = append(padded, sfx...)

	fmt.Printf("%q\n", padded)
	fmt.Println(randCipherMode())
	return nil, nil
}

type cipherMode int

const (
	ecbMode cipherMode = iota
	cbcMode
)

func randCipherMode() cipherMode {
	return cipherMode(randInt(0, 2))
}

func randInt(min, max int64) int {
	n, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err.Error())
	}
	return int(min + n.Int64())
}

func randAESKey() ([]byte, error) {
	k := make([]byte, aes.BlockSize)
	if _, err := rand.Read(k); err != nil {
		return nil, err
	}
	return k, nil
}
