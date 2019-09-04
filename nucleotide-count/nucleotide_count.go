package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

var validNucleotides = [4]rune{'A', 'G', 'C', 'T'}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := make(Histogram)
	for _, r := range validNucleotides {
		h[r] = 0
	}
	for _, nuc := range d {
		if _, f := h[nuc]; !f {
			return nil, errors.New("invalid input: this nucleotide does not exists")
		}
		h[nuc]++
	}
	return h, nil
}
