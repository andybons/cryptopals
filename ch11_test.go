package cryptopals

import (
	"crypto/aes"
	"testing"
)

func TestDetectCipherMode(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ciphertext := make([]byte, 3*aes.BlockSize)
		for k := range ciphertext {
			ciphertext[k] = 'A'
		}
		b, m, err := encryptUsingRandomCipherMode(ciphertext)
		if err != nil {
			t.Fatalf("encryptUsingRandomCipherMode(%q): got unexpected error: %v", ciphertext, err)
		}
		if got, want := detectCipherMode(b), m; got != want {
			t.Errorf("detectCipherMode = %v; want %v", got, want)
			break
		}
	}
}
