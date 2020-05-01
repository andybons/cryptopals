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
		{[]byte("foo@bar.com"), []byte("email=foo@bar.com&uid=10&role=user")},
		{[]byte("foo@bar.com&role=admin"), []byte("email=foo@bar.comroleadmin&uid=10&role=user")},
	}
	for _, tc := range testCases {
		if got, want := profileFor(tc.email), tc.want; !bytes.Equal(got, want) {
			t.Errorf("profileFor(%q) = %q; want %q", tc.email, got, want)
		}
	}
}

func TestEncryptDecryptUserProfile(t *testing.T) {
	key := make([]byte, aes.BlockSize)
	if _, err := rand.Read(key); err != nil {
		t.Fatalf("rand.Read: got unexpected error: %v", err)
	}

	// email=XXXXXXXXXXXXX&uid=10&role=user____________
	// ----------------++++++++++++++++----------------
	b, err := profileOracle([]byte(`XXXXXXXXXXXXX`), key)
	if err != nil {
		t.Errorf("profileOracle: got unexpected error: %v", err)
	}

	// Construct admin block via oracle.
	adminBlock := padPKCS7([]byte(`admin`), aes.BlockSize)

	// email=XXXXXXXXXXadmin___________&uid=10&role=user_______________
	// ----------------++++++++++++++++----------------++++++++++++++++
	ab, err := profileOracle(append([]byte("XXXXXXXXXX"), adminBlock...), key)
	if err != nil {
		t.Errorf("profileOracle: got unexpected error: %v", err)
	}

	// Replace last block.
	copy(b[aes.BlockSize*2:aes.BlockSize*3], ab[aes.BlockSize*1:aes.BlockSize*2])

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
