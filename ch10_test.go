package cryptopals

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncryptAESinECB(t *testing.T) {
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

func TestEncryptAESinCBC(t *testing.T) {
	f, err := os.Open("./testdata/ch10.txt")
	if err != nil {
		t.Fatalf("os.Open: got unexpected error: %v", err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, f))
	if err != nil {
		t.Fatalf("ioutil.ReadAll: got unexpected error: %v", err)
	}
	key := []byte(`YELLOW SUBMARINE`)
	b, err = decryptAESinCBC(b, key)
	if err != nil {
		t.Errorf("decryptAESinCBC: got unexpected error %v", err)
	}
	if got, want := b[:2*aes.BlockSize], []byte(`I'm back and I'm ringin' the bel`); !bytes.Equal(got, want) {
		t.Errorf("decryptAESinCBC = %q; want %q", got, want)
	}
}
