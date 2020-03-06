package cryptopals

func maxBlockDupeCount(b []byte, blockSize int) int {
	m := make(map[string]int)
	var max int
	for i := 0; i < len(b); i += blockSize {
		s := string(b[i : i+blockSize])
		m[s]++
		if m[s] > 1 && m[s] > max {
			max = m[s]
		}
	}
	return max
}
