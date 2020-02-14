package set01

import (
	"encoding/base64"
)

func hexToBase64(src []byte) ([]byte, error) {
	b, err := decodeHex(src)
	if err != nil {
		return nil, err
	}
	enc := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(enc, b)
	return enc, nil
}
