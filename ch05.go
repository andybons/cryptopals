package cryptopals

import (
	"encoding/hex"
)

func repeatingKeyXor(b []byte, key []byte) []byte {
	enc := make([]byte, len(b))
	for i := range b {
		enc[i] = b[i] ^ key[i%len(key)]
	}
	return enc
}

func encodeHex(src []byte) []byte {
	b := make([]byte, hex.EncodedLen(len(src)))
	n := hex.Encode(b, src)
	return b[:n]
}
