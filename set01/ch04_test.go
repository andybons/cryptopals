package set01

import (
	"bufio"
	"encoding/hex"
	"os"
	"testing"
)

func TestFindEncryptedLine(t *testing.T) {
	f, err := os.Open("./testdata/ch04.txt")
	if err != nil {
		t.Fatalf("os.Open: got unexpected error: %v", err)
	}
	defer f.Close()

	scorer, err := newFrequencyScorer()
	if err != nil {
		t.Fatalf("newFrequencyScorer(): got unexpected error %v", err)
	}
	scanner := bufio.NewScanner(f)
	var best struct {
		key   byte
		score float64
		line  []byte
		bytes []byte
	}
	for scanner.Scan() {
		l := scanner.Bytes()
		b := make([]byte, hex.DecodedLen(len(l)))
		n, err := hex.Decode(b, l)
		if err != nil {
			t.Fatalf("hex.Decode: got unexpected error %v", err)
		}
		b = b[:n]
		key, err := findXorKey(b, scorer)
		if err != nil {
			t.Errorf("findXorKey(%q): got expected error %v", b, err)
		}
		if sc := scorer.score(singleXor(b, key)); sc > best.score {
			best.key = key
			best.score = sc
			best.line = l
			best.bytes = b
		}
	}
	if err := scanner.Err(); err != nil {
		t.Errorf("scanner.Err(): got unexpected error: %v", err)
	}
	t.Logf("line: %s; Key: %c; Decrypted string: %q", best.line[:8], best.key, singleXor(best.bytes, best.key))
}
