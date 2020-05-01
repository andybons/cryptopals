package cryptopals

import (
	"fmt"
	"net/url"
	"strings"
)

func parseURLQuery(q []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	pairs := strings.Split(string(q), "&")
	for _, p := range pairs {
		kv := strings.Split(p, "=")
		ke, err := url.QueryUnescape(kv[0])
		if err != nil {
			return nil, err
		}
		ve, err := url.QueryUnescape(kv[1])
		if err != nil {
			return nil, err
		}
		m[ke] = ve
	}
	return m, nil
}

func profileFor(email []byte) []byte {
	return encodeURLQuery(map[string]interface{}{
		"email": string(email),
		"uid":   10,
		"role":  "user",
	})
}

func encodeURLQuery(m map[string]interface{}) []byte {
	return []byte(fmt.Sprintf("email=%s&uid=%s&role=%s",
		stripMetas(fmt.Sprint(m["email"])),
		stripMetas(fmt.Sprint(m["uid"])),
		stripMetas(fmt.Sprint(m["role"])),
	))
}

func stripMetas(s string) string {
	s = strings.ReplaceAll(s, "&", "")
	s = strings.ReplaceAll(s, "=", "")
	return s
}

func profileOracle(email, key []byte) ([]byte, error) {
	return encryptAESinECB(profileFor(email), key)
}
