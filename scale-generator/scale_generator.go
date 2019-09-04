// Solution from https://exercism.io/tracks/go/exercises/scale-generator/solutions/1d75c67188ef4e3dab791038898ed807
// Study it!
package scale

import "strings"

var scales = [2][12]string{
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"},
}
var kind = map[string]struct {
	s int
	p int
}{
	"G": {0, 7}, "D": {0, 2}, "A": {0, 9}, "E": {0, 4}, "B": {0, 11}, "F#": {0, 6},
	"e": {0, 7}, "b": {0, 2}, "f#": {0, 9}, "c#": {0, 4}, "g#": {0, 11}, "d#": {0, 6},
	"F": {1, 5}, "Bb": {1, 10}, "Eb": {1, 3}, "Ab": {1, 8}, "Db": {1, 1}, "Gb": {1, 6},
	"d": {1, 5}, "g": {1, 10}, "c": {1, 3}, "f": {1, 8}, "bb": {1, 1}, "eb": {1, 6},
}
var step = map[byte]int{'A': 3, 'M': 2, 'm': 1}

// Scale generates a scale based on tonic and interval.
func Scale(t string, i string) []string {
	s, p := scales[kind[t].s], kind[t].p
	if strings.ToLower(t) == t {
		p = (p + len(s) - 3) % len(s)
	}

	if i == "" {
		return append(s[p:], s[:p]...)
	}

	r := []string{s[p]}
	for c := range i[:len(i)-1] {
		p = (p + step[i[c]]) % len(s)
		r = append(r, s[p])
	}
	return r
}
