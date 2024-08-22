package random_test

import (
	"caga/websocket/random"
	"math/big"
	"testing"
)

func TestUniqueRandom(t *testing.T) {
	var generated = map[string]bool{}
	var maxNumber = new(big.Int).SetUint64(1000)

	for i := 0; i < 1000; i++ {
		num, err := random.Random(generated, maxNumber)
		if err != nil {
			t.Fatal(err)
		}

		if !generated[num.String()] {
			t.Fatal("generated number should be saved in the map")
		}

		if num.Cmp(maxNumber) >= 0 {
			t.Fatal("generated number should smaller than max")
		}
	}

	_, err := random.Random(generated, maxNumber)
	if err.Error() != "out of number" {
		t.Fatal("shoud return our of number")
	}
}
