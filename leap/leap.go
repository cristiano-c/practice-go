// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear returns true if the given year is leap
func IsLeapYear(year int) bool {
	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%4 == 0 && year%100 != 0 {
		leap = true
	} else {
		leap = false
	}
	return leap
}
