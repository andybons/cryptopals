package set01

import (
	"fmt"
	"math"
)

func hammingDist(b1, b2 []byte) int {
	// This assumes that b1 and b2 have the same length.
	var dist int
	for i := range b1 {
		for j := 0; j < 8; j++ {
			if (b1[i]^b2[i])&(1<<j) > 0 {
				dist++
			}
		}
	}
	return dist
}

func crackRepeatingKeyXor(b []byte) (key []byte, err error) {
	keysize := probableKeySize(b)
	fmt.Println("Probable key size:", keysize)
	fmt.Println("Length of cipher:", len(b))

	// Now that you probably know the KEYSIZE: break the ciphertext
	// into blocks of KEYSIZE length.
	// Now transpose the blocks: make a block that is the first byte
	// of every block, and a block that is the second byte of every
	// block, and so on.
	numBlocks := int(math.Ceil(float64(len(b)) / float64(keysize)))
	transposed := make([][]byte, keysize)
	for i := 0; i < keysize; i++ {
		for j := 0; j < numBlocks; j++ {
			idx := keysize*j + i
			if idx >= len(b) {
				break
			}
			transposed[i] = append(transposed[i], b[idx])
		}
	}

	fs, err := newFrequencyScorer()
	if err != nil {
		return nil, err
	}
	k := make([]byte, keysize)
	for i := range k {
		k[i], err = findXorKey(transposed[i], fs)
		if err != nil {
			return nil, err
		}
	}

	return k, nil
}

func probableKeySize(b []byte) int {
	minDist := math.MaxFloat64
	var keylen int
	for keysize := 2; keysize <= 40; keysize++ {
		var dist float64
		for i := 0; i < 10; i++ {
			b0 := b[keysize*i : keysize*(i+1)]
			b1 := b[keysize*(i+1) : keysize*(i+2)]
			dist += float64(hammingDist(b0, b1))
		}
		nDist := dist / float64(keysize)
		if nDist < minDist {
			minDist = nDist
			keylen = keysize
		}
	}
	fmt.Printf("min dist: %f\tkeysize: %d\n", minDist, keylen)
	return keylen
}
