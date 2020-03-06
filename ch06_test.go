package cryptopals

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	b1 := []byte(`this is a test`)
	b2 := []byte(`wokka wokka!!!`)
	if got, want := hammingDist(b1, b2), 37; got != want {
		t.Errorf("hammingDist(%q, %q) = %d, want %d", b1, b2, got, want)
	}
}

func TestCrackRepeatingKeyXor(t *testing.T) {
	f, err := os.Open("./testdata/ch06.txt")
	if err != nil {
		t.Fatalf("os.Open: got unexpected error: %v", err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, f))
	if err != nil {
		t.Fatalf("ioutil.ReadAll: got unexpected error: %v", err)
	}
	got, err := crackRepeatingKeyXor(b)
	if err != nil {
		t.Fatalf("crackRepeatingKeyXor: got unexpected error: %v", err)
	}
	if want := []byte(`Terminator X: Bring the noise`); !bytes.Equal(got, want) {
		t.Errorf("crackRepeatingKeyXor: key = %q; want %q", got, want)
	}
	t.Logf("Key: %q; Text: %q", got, repeatingKeyXor(b, got))
}
