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
	runeCount map[rune]int
	runeTotal int
	runeScore map[rune]float64
}

func newFrequencyScorer() (*frequencyScorer, error) {
	f, err := os.Open("./testdata/pride_and_prejudice.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rr := bufio.NewReader(f)
	fs := &frequencyScorer{
		runeCount: make(map[rune]int),
		runeScore: make(map[rune]float64),
	}
	for {
		ch, _, err := rr.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		fs.runeCount[ch]++
		fs.runeTotal++
	}
	for ch, n := range fs.runeCount {
		fs.runeScore[ch] = float64(n) / float64(fs.runeTotal)
	}
	return fs, nil
}

func (fs *frequencyScorer) score(b []byte) float64 {
	var s float64
	for _, ch := range b {
		s += fs.runeScore[rune(ch)]
	}
	return s
}
