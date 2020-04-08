package cryptopals

import (
	"crypto/aes"
	"crypto/rand"
	"math/big"
)

func encryptUsingRandomCipherMode(b []byte) ([]byte, cipherMode, error) {
	pfx := make([]byte, randInt(5, 11))
	if _, err := rand.Read(pfx); err != nil {
		return nil, cipherModeUnknown, err
	}
	sfx := make([]byte, randInt(5, 11))
	if _, err := rand.Read(sfx); err != nil {
		return nil, cipherModeUnknown, err
	}

	padded := append(pfx, b...)
	padded = append(padded, sfx...)

	key, err := randAESKey()
	if err != nil {
		return nil, cipherModeUnknown, err
	}

	var enc []byte
	mode := randCipherMode()
	switch mode {
	case cipherModeECB:
		enc, err = encryptAESinECB(padded, key)
		if err != nil {
			return nil, mode, err
		}
	case cipherModeCBC:
		iv := make([]byte, aes.BlockSize)
		if _, err := rand.Read(iv); err != nil {
			return nil, mode, err
		}
		enc, err = encryptAESinCBC(padded, key, iv)
		if err != nil {
			return nil, mode, err
		}
	}
	return enc, mode, nil
}

func detectCipherMode(b []byte) cipherMode {
	if maxBlockDupeCount(b, aes.BlockSize) > 0 {
		return cipherModeECB
	}
	return cipherModeCBC
}

type cipherMode int

func (m cipherMode) String() string {
	switch m {
	case cipherModeECB:
		return "ECB"
	case cipherModeCBC:
		return "CBC"
	case cipherModeUnknown:
		return "Unknown"
	default:
		panic("cipher mode not present in String() function")
	}
}

const (
	cipherModeUnknown cipherMode = iota
	cipherModeECB
	cipherModeCBC
)

func randCipherMode() cipherMode {
	return cipherMode(randInt(1, 3))
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
