package cryptopals

import (
	"bytes"
	"testing"
)

func TestFixedXorHex(t *testing.T) {
	b1 := []byte(`1c0111001f010100061a024b53535009181c`)
	b2 := []byte(`686974207468652062756c6c277320657965`)
	got, err := fixedXorHex(b1, b2)
	if err != nil {
		t.Errorf("fixedXorHex(): got unexpected error: %v", err)
	}
	want := []byte(`746865206b696420646f6e277420706c6179`)
	if !bytes.Equal(got, want) {
		t.Errorf("fixedXorHex(%q, %q) = %q, want %q", b1, b2, got, want)
	}
}
