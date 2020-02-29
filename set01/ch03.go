package set01

import (
	"bufio"
	"encoding/hex"
	"io"
	"os"
)

func findXorKey(b []byte, s scorer) (byte, error) {
	var best struct {
		key   byte
		score float64
	}
	for i := 0; i < 256; i++ {
		ch := byte(i)
		xord := singleXor(b, ch)
		if sc := s.score(xord); sc > best.score {
			best.score = sc
			best.key = ch
		}
	}
	return best.key, nil
}

func decodeHex(b []byte) ([]byte, error) {
	b2 := make([]byte, hex.DecodedLen(len(b)))
	n, err := hex.Decode(b2, b)
	if err != nil {
		return nil, err
	}
	return b2[:n], nil
}

func singleXor(b []byte, x byte) []byte {
	r := make([]byte, len(b))
	for i := range b {
		r[i] = b[i] ^ x
	}
	return r
}

type scorer interface {
	score([]byte) float64
}

type frequencyScorer struct {
	counts map[byte]int
	total  int
	scores map[byte]float64
}

func newFrequencyScorer() (*frequencyScorer, error) {
	f, err := os.Open("./testdata/pride_and_prejudice.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rr := bufio.NewReader(f)
	fs := &frequencyScorer{
		counts: make(map[byte]int),
		scores: make(map[byte]float64),
	}
	for {
		ch, err := rr.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		fs.counts[ch]++
		fs.total++
	}
	for ch, n := range fs.counts {
		fs.scores[ch] = float64(n) / float64(fs.total)
	}
	return fs, nil
}

func (fs *frequencyScorer) score(b []byte) float64 {
	var s float64
	for _, ch := range b {
		s += fs.scores[ch]
	}
	return s
}
