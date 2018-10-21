package hamming

import (
	"errors"
)

// Distance - calculate it
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("invalid input, sequences must be of equal length")
	}

	var hammingDistance = 0
	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
