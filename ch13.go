package cryptopals

import (
	"net/url"
	"strings"
)

func parseURLQuery(q string) (map[string]string, error) {
	m := make(map[string]string)
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

func encodeURLQuery(m map[string]string) string {
	var sb strings.Builder
	for k, v := range m {
		if sb.Len() > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(url.QueryEscape(k))
		sb.WriteByte('=')
		sb.WriteString(url.QueryEscape(v))
	}
	return sb.String()
}
