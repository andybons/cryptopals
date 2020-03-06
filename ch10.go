package cryptopals

import "crypto/aes"

func encryptAESinECB(b, key []byte) ([]byte, error) {
	b = padPKCS7(b, aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(b); i += aes.BlockSize {
		block.Encrypt(b[i:], b[i:i+aes.BlockSize])
	}
	return b, nil
}
