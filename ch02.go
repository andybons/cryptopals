package cryptopals

import (
	"encoding/hex"
	"errors"
)

func fixedXorHex(h1, h2 []byte) ([]byte, error) {
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
	b, err := fixedXor(b1, b2)
	if err != nil {
		return nil, err
	}
	r := make([]byte, hex.EncodedLen(len(b)))
	n := hex.Encode(r, b)
	return r[:n], nil
}

func fixedXor(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, errors.New("mismatched lengths")
	}
	b := make([]byte, len(b1))
	for i := range b {
		b[i] = b1[i] ^ b2[i]
	}
	return b, nil
}
