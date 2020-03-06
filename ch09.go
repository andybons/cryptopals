package cryptopals

func padPKCS7(b []byte, blockLen int) []byte {
	var n int
	for n < len(b) {
		n += blockLen
	}
	pb := make([]byte, n)
	for i := range pb {
		if i < len(b) {
			pb[i] = b[i]
		} else {
			pb[i] = '\x04'
		}
	}
	return pb
}
