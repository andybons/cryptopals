package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"testing"
)

func TestParseURLQuery(t *testing.T) {
	in := []byte(`foo=bar&baz=qux&zap=zazzle`)
	m, err := parseURLQuery(in)
	if err != nil {
		t.Fatalf("parseURLQuery(%q): got unexpected error: %v", in, err)
	}
	if got, want := len(m), 3; got != want {
		t.Errorf("len(parseURLQuery(%q) = %d; want %d", in, got, want)
	}
	for k, v := range map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	} {
		if got, want := m[k], v; got != want {
			t.Errorf("m[%q] = %q; want %q", k, got, want)
		}
	}
}

func TestProfileFor(t *testing.T) {
	testCases := []struct {
		email, want []byte
	}{
		{[]byte("foo@bar.com"), []byte("email=foo%40bar.com&uid=10&role=user")},
		{[]byte("foo@bar.com&role=admin"), []byte("email=foo%40bar.com%26role%3Dadmin&uid=10&role=user")},
	}
	for _, tc := range testCases {
		if got, want := profileFor(tc.email), tc.want; !bytes.Equal(got, want) {
			t.Errorf("profileFor(%+v) = %q; want %q", tc.email, got, want)
		}
	}
}

func TestEncryptDecryptUserProfile(t *testing.T) {
	key := make([]byte, aes.BlockSize)
	if _, err := rand.Read(key); err != nil {
		t.Fatalf("rand.Read: got unexpected error: %v", err)
	}

	b, err := profileOracle([]byte(`andybo@gmail.com`), key)
	if err != nil {
		t.Errorf("profileOracle: got unexpected error: %v", err)
	}

	// Encrypted bytes will equal:
	// email=andybo%40gmail.com&uid=10&role=user*******
	// ----------------++++++++++++++++----------------

	// Replace last block.
	enc, err := encryptAESinECB([]byte(`role=admin`), key)
	if err != nil {
		t.Errorf("encryptAESinECB: got unexpected error: %v", err)
	}

	copy(b[aes.BlockSize*2:aes.BlockSize*3], enc)

	dec, err := decryptAESinECB(b, key)
	if err != nil {
		t.Errorf("decryptAESinECB: got unexpected error: %v", err)
	}
	m, err := parseURLQuery(dec)
	if err != nil {
		t.Errorf("parseURLQuery(%q): got unexpected error: %v", b, err)
	}
	if got, want := m["role"], "admin"; got != want {
		t.Errorf(`m["role"] = %q; want %q`, got, want)
	}
	t.Logf("%+v", m)
}
