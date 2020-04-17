package cryptopals

import (
	"bytes"
	"fmt"
	"net/url"
	"sort"
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
	var buf bytes.Buffer
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(fmt.Sprint(m[k])))
	}
	return buf.Bytes()
}
