package strand

var transcription = map[rune]rune{
	'A': 'U',
	'T': 'A',
	'C': 'G',
	'G': 'C',
}

// ToRNA returns the RNA transcription of a given DNA
func ToRNA(dna string) string {
	rna := []rune(dna)
	for i, nucleotide := range dna {
		rna[i] = rune(transcription[nucleotide])
	}
	return string(rna)
}
