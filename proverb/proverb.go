// Package proverb for want of...
package proverb

import "fmt"

// Proverb generates rhymes for the want of proverb...
func Proverb(rhymes []string) []string {
	var text []string
	baseline := "For want of a %s the %s was lost."
	endline := "And all for the want of a %s."
	if len(rhymes) != 0 {
		if len(rhymes) == 1 {
			text = append(text, fmt.Sprintf(endline, rhymes[0]))
		} else {
			for i := 0; i < len(rhymes)-1; i++ {
				text = append(text, fmt.Sprintf(baseline, rhymes[i], rhymes[i+1]))
			}
			text = append(text, fmt.Sprintf(endline, rhymes[0]))
		}
	}
	return text
}
