package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
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
	b, err = decryptAESinECB(b, key)
	if err != nil {
		t.Errorf("decryptAESinECB: got unexpected error %v", err)
	}
	if got, want := b[:len(ciphertext)], ciphertext; !bytes.Equal(got, want) {
		t.Errorf("decryptAESinECB = %q; want %q", got, want)
	}
}

func TestDecryptAESinCBC(t *testing.T) {
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
	iv := make([]byte, aes.BlockSize)
	b, err = decryptAESinCBC(b, key, iv)
	if err != nil {
		t.Errorf("decryptAESinCBC: got unexpected error %v", err)
	}
	if got, want := b[:2*aes.BlockSize], []byte(`I'm back and I'm ringin' the bel`); !bytes.Equal(got, want) {
		t.Errorf("decryptAESinCBC = %q; want %q", got, want)
	}
}

func TestEncryptAESinCBC(t *testing.T) {
	ciphertext := []byte(`started from the bottom now we here`)
	key := []byte(`YELLOW SUBMARINE`)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		t.Fatalf("rand.Read(iv): got unexpected error: %v", err)
	}

	enc, err := encryptAESinCBC(ciphertext, key, iv)
	if err != nil {
		t.Fatalf("encryptAESinCBC: got unexpected error: %v", err)
	}
	b, err := decryptAESinCBC(enc, key, iv)
	if err != nil {
		t.Fatalf("decryptAESinCBC: got unexpected error: %v", err)
	}
	if got, want := b[:len(ciphertext)], ciphertext; !bytes.Equal(got, want) {
		t.Errorf("decryptAESinCBC = %q; want %q", got, want)
	}
}
