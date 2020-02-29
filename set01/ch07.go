package set01

import (
	"crypto/aes"
	"errors"
)

func decryptAESinECB(b, key []byte) error {
	if len(b)%aes.BlockSize != 0 {
		return errors.New("length of ciphertext must be multiple of aes.BlockSize")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	for i := 0; i < len(b); i += aes.BlockSize {
		block.Decrypt(b[i:], b[i:i+aes.BlockSize])
	}
	return nil
}
