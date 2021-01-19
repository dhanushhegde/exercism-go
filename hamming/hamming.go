package hamming

import "errors"

// Distance returns hamming distance between provided inputs
// it returns an error if unequal length inputs are provided
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Inputs are of different length")
	}
	numberOfDifference := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			numberOfDifference++
		}
	}
	return numberOfDifference, nil
}
