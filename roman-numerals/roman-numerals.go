package romannumerals

import "errors"

var (
	chunk0  = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	chunk01 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	chunk02 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	chunk03 = []string{"", "M", "MM", "MMM"}
)

// ToRomanNumeral converts arabic numbers to roman notation
func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic >= 3001 {
		return "", errors.New("Out of range")
	}
	return chunk03[arabic%1e4/1e3] + chunk02[arabic%1e3/1e2] + chunk01[arabic%100/10] + chunk0[arabic%10], nil
}
