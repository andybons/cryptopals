package cryptopals

import "testing"

func TestEncryptionOracle(t *testing.T) {
	encryptionOracle([]byte(`hello, encryption!`))
}
