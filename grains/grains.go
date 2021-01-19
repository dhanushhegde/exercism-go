package grains

import (
	"errors"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Tile number should be greater than zero")
	}

	// return uint64(math.Pow(2, float64(n-1))), nil
	return 1 << (n - 1), nil

}

func Total() uint64 {

	// Sum of powe of 2 from 0 to n = 2^n+1-1
	// 2^3   1 2 4 8 -1  2^4-1
	// var n float64
	// var m uint64
	// n = math.Pow(2, float64(64)) - 1
	// m = uint64(math.Pow(2, float64(64)) - 1)
	// fmt.Println(n)
	// fmt.Println(m)
	// return uint64(n)
	// return math.MaxUint64
	return 1<<64 - 1
}
