package cryptopals

import (
	"bytes"
	"testing"
)

func TestAttackAESinECB(t *testing.T) {
	b, err := decryptAESinCBCOracle()
	if err != nil {
		t.Errorf("decryptAESinCBCOracle(): got unexpected error: %v", err)
	}
	if got, want := b, mysteryBytes(); !bytes.Equal(got, want) {
		t.Errorf("decrypted mystery bytes = %q\nwant %q", got, want)
	}
	t.Logf("Got %s", b)
}
