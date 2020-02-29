package set01

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecryptAESinECB(t *testing.T) {
	f, err := os.Open("./testdata/ch07.txt")
	if err != nil {
		t.Fatalf("os.Open: got unexpected error %v", err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, f))
	if err != nil {
		t.Fatalf("ioutil.ReadAll: got unexpected error: %v", err)
	}

	key := []byte(`YELLOW SUBMARINE`)
	if err := decryptAESinECB(b, key); err != nil {
		t.Errorf("decryptAESinECB: got unexpected error %v", err)
	}
	if got, want := b[:aes.BlockSize*2], []byte(`I'm back and I'm ringin' the bel`); !bytes.Equal(got, want) {
		t.Errorf("first two blocks of decrypted ciphertext = %q; want %q", got, want)
	}
	t.Logf("Decrypted text: %q", b)
}
