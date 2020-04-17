package cryptopals

import (
	"crypto/aes"
	"errors"
)

func decryptAESinECB(b, key []byte) ([]byte, error) {
	if len(b)%aes.BlockSize != 0 {
		return nil, errors.New("length of ciphertext must be multiple of aes.BlockSize")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	d := make([]byte, len(b))
	for i := 0; i < len(b); i += aes.BlockSize {
		block.Decrypt(d[i:], b[i:i+aes.BlockSize])
	}
	for d[len(d)-1] == '\x04' {
		d = d[:len(d)-1]
	}
	return d, nil
}
