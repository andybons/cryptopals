package cryptopals

import (
	"bytes"
	"testing"
)

func TestEncryptAESinCBC(t *testing.T) {
	ciphertext := []byte(`started from the bottom now we here`)
	key := []byte(`YELLOW SUBMARINE`)
	b, err := encryptAESinECB(ciphertext, key)
	if err != nil {
		t.Fatalf("encryptAESinECB(%q, %q): got expected error %v", ciphertext, key, err)
	}
	if err := decryptAESinECB(b, key); err != nil {
		t.Errorf("decryptAESinECB: got unexpected error %v", err)
	}
	if got, want := b[:len(ciphertext)], ciphertext; !bytes.Equal(got, want) {
		t.Errorf("decryptAESinECB = %q; want %q", got, want)
	}
}
