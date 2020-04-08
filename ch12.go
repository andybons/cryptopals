package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func blockSizeOfOracle(key []byte) (int, error) {
	var encLen, blockSize int
	for i := 0; i < 100; i++ {
		var buf bytes.Buffer
		for j := 0; j < i; j++ {
			buf.WriteByte('A')
		}
		b, err := encryptAESInECBOracle(buf.Bytes(), key)
		if err != nil {
			return 0, err
		}
		if encLen == 0 {
			encLen = len(b)
		}
		if encLen != len(b) {
			blockSize = len(b) - encLen
			break
		}
	}
	return blockSize, nil
}

func cipherModeOfOracle(key []byte) (cipherMode, error) {
	var buf bytes.Buffer
	for i := 0; i < 2*aes.BlockSize; i++ {
		if err := buf.WriteByte('A'); err != nil {
			return cipherModeUnknown, err
		}
	}
	return detectCipherMode(buf.Bytes()), nil
}

func decryptAESinCBCOracle() ([]byte, error) {
	key := make([]byte, aes.BlockSize)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	blockSize, err := blockSizeOfOracle(key)
	if err != nil {
		return nil, err
	}
	if blockSize != aes.BlockSize {
		return nil, fmt.Errorf("block size = %d; want %d", blockSize, aes.BlockSize)
	}

	var (
		dec      []byte // the decrypted stuff
		blockNum int
	)
	for {
		padLen := blockSize - 1 - len(dec)%blockSize

		var buf bytes.Buffer
		for i := 0; i < padLen; i++ {
			buf.WriteByte('A')
		}
		buf.Write(dec)

		blockStart := blockSize * blockNum
		m := make(map[string]byte)
		for i := 0; i < 0xff; i++ {
			ch := byte(i)
			buf.WriteByte(ch)
			b, err := encryptAESInECBOracle(buf.Bytes(), key)
			if err != nil {
				return nil, err
			}
			m[string(b[blockStart:blockStart+blockSize])] = ch
			buf.Truncate(blockSize*(blockNum+1) - 1)
		}

		buf.Truncate(padLen)
		b, err := encryptAESInECBOracle(buf.Bytes(), key)
		if err != nil {
			return nil, err
		}

		if blockStart+blockSize > len(b) {
			break
		}

		dec = append(dec, m[string(b[blockStart:blockStart+blockSize])])
		blockNum = len(dec) / blockSize
	}

	// Strip trailing padding.
	const pkcs7PaddingByte = '\x04'
	for dec[len(dec)-1] == pkcs7PaddingByte {
		dec = dec[:len(dec)-1]
	}
	return dec, nil
}

func encryptAESInECBOracle(b, key []byte) ([]byte, error) {
	return encryptAESinECB(append(b, mysteryBytes()...), key)
}

func mysteryBytes() []byte {
	const input = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	b, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		panic(err.Error())
	}
	return b
}
