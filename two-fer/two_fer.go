// Package provides an example of string concatenation
package twofer

import "fmt"

// ShareWith returns 2-fer, if argument string is null, default is 'you'
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
