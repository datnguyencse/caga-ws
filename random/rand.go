package random

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func Random(generated map[string]bool, max *big.Int) (*big.Int, error) {
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}

	// check to see if we already used all number in range
	if big.NewInt(int64(len(generated))).Cmp(max) >= 0 {
		return nil, errors.New("out of number")
	}

	if generated[num.String()] {
		return Random(generated, max)
	}

	generated[num.String()] = true
	return num, nil
}
