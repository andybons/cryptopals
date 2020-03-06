package cryptopals

import (
	"bytes"
	"testing"
)

func TestPKCS7Padding(t *testing.T) {
	in := []byte(`YELLOW SUBMARINE`)
	want := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	if got := padPKCS7(in, 10); !bytes.Equal(got, want) {
		t.Errorf("padPKCS7(%q, 10) = %q; want %q", in, got, want)
	}
}
