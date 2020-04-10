package cryptopals

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func parseURLQuery(q string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	pairs := strings.Split(q, "&")
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

func profileFor(email string) string {
	return encodeURLQuery(map[string]interface{}{
		"email": email,
		"uid":   10,
		"role":  "user",
	})
}

func encodeURLQuery(m map[string]interface{}) string {
	var sb strings.Builder
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		if sb.Len() > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(url.QueryEscape(k))
		sb.WriteByte('=')
		sb.WriteString(url.QueryEscape(fmt.Sprint(m[k])))
	}
	return sb.String()
}
