package cryptopals

import (
	"crypto/aes"
)

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

func decryptAESinCBC(b, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	prevBlk := make([]byte, aes.BlockSize)
	copy(prevBlk, iv)
	dec := make([]byte, len(b))
	for i := 0; i < len(b); i += aes.BlockSize {
		blk := b[i : i+aes.BlockSize]
		block.Decrypt(dec[i:], blk)
		for j := 0; j < aes.BlockSize; j++ {
			dec[i+j] ^= prevBlk[j]
		}
		prevBlk = blk
	}
	return dec, nil
}

func encryptAESinCBC(b, key, iv []byte) ([]byte, error) {
	b = padPKCS7(b, aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	prevBlk := make([]byte, aes.BlockSize)
	copy(prevBlk, iv)
	enc := make([]byte, len(b))
	for i := 0; i < len(b); i += aes.BlockSize {
		blk := b[i : i+aes.BlockSize]
		for j := 0; j < aes.BlockSize; j++ {
			blk[j] ^= prevBlk[j]
		}
		block.Encrypt(enc[i:], blk)
		prevBlk = enc[i : i+aes.BlockSize]
	}
	return enc, nil
}
