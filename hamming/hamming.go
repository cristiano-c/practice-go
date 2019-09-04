// Package hamming provides functionality to calculate the Hamming distange between two DNA strands
package hamming

import "errors"

// Distance calculates the Hamming distange between two given DNA strands
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	var err error
	var distance int
	if len(ar) != len(br) {
		err = errors.New("math: square root of negative number")
	} else {
		for i := range ar {
			if ar[i] != br[i] {
				distance++
			}
		}
	}
	return distance, err
}
