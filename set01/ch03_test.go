package set01

import "testing"

func TestFindXorKey(t *testing.T) {
	h := []byte(`1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736`)
	b, err := decodeHex(h)
	if err != nil {
		t.Fatalf("decodeHex(%q): got unexpected error %v", h, err)
	}
	fs, err := newFrequencyScorer()
	if err != nil {
		t.Fatalf("newFrequencyScorer(): got unexpected error %v", err)
	}
	key, err := findXorKey(b, fs)
	if err != nil {
		t.Errorf("findXorKey(%q): got expected error %v", b, err)
	}
	if got, want := key, byte('X'); got != want {
		t.Errorf("findXorKey(%q) = %c, want %c", b, got, want)
	}
	t.Logf("Key: %c; Decrypted string: %q", key, singleXor(b, key))
}
