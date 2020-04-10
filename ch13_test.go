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

func TestProfileFor(t *testing.T) {
	testCases := []struct {
		email, want string
	}{
		{"foo@bar.com", "email=foo%40bar.com&role=user&uid=10"},
		{"foo@bar.com&role=admin", "email=foo%40bar.com%26role%3Dadmin&role=user&uid=10"},
	}
	for _, tc := range testCases {
		if got, want := profileFor(tc.email), tc.want; got != want {
			t.Errorf("profileFor(%+v) = %q; want %q", tc.email, got, want)
		}
	}
}
