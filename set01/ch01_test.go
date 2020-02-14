package set01

import (
	"bytes"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	hex := []byte(`49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d`)
	want := []byte(`SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t`)

	got, err := hexToBase64(hex)
	if err != nil {
		t.Errorf("hexToBase64(%q): got unexpected error %v", hex, err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("hexToBase64(%q) = %q, want %q", hex, got, want)
	}
}

func BenchmarkHexToBase64(b *testing.B) {
	in := []byte(`49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hexToBase64(in)
	}
}
