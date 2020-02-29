package set01

import (
	"bufio"
	"crypto/aes"
	"os"
	"testing"
)

func TestDetectAESinECB(t *testing.T) {
	f, err := os.Open("./testdata/ch08.txt")
	if err != nil {
		t.Fatalf("os.Open: got unexpected error: %v", err)
	}
	defer f.Close()
	var best struct {
		line      []byte
		blockDups int
		lineNum   int
	}
	var lineNum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineNum++
		b := scanner.Bytes()
		dupeCount := maxBlockDupeCount(b, aes.BlockSize)
		if dupeCount == 0 {
			continue
		}
		t.Logf("Line: %d; block dupes: %d; %q", lineNum, dupeCount, b)
		if dupeCount > best.blockDups {
			best.blockDups = dupeCount
			best.line = make([]byte, len(b))
			copy(best.line, b)
			best.lineNum = lineNum
		}
	}
	t.Logf("Line: %d; Block dupes: %d; %q", best.lineNum, best.blockDups, best.line)
}
