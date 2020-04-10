package cryptopals

import "testing"

func TestParseURLQuery(t *testing.T) {
	in := `foo=bar&baz=qux&zap=zazzle`
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

func TestEncodeURLQuery(t *testing.T) {
	m := map[string]string{
		"email": "foo@bar.com",
		"uid":   "10",
		"role":  "user",
	}
	q := encodeURLQuery(m)
	_ = q
}
