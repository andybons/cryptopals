package cryptopals

import (
	"encoding/hex"
	"errors"
)

func fixedXor(h1, h2 []byte) ([]byte, error) {
	if len(h1) != len(h2) {
		return nil, errors.New("mismatched lengths")
	}

	b1, err := decodeHex(h1)
	if err != nil {
		return nil, err
	}
	b2, err := decodeHex(h2)
	if err != nil {
		return nil, err
	}
	for i := range b1 {
		b1[i] = b1[i] ^ b2[i]
	}
	r := make([]byte, hex.EncodedLen(len(b1)))
	n := hex.Encode(r, b1)
	return r[:n], nil
}
