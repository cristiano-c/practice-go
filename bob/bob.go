// Package bob exercise
package bob

import "regexp"

var answers = map[string]string{
	"whatever": "Whatever.",
	"whoa":     "Whoa, chill out!",
	"sure":     "Sure.",
	"calm":     "Calm down, I know what I'm doing!",
	"fine":     "Fine. Be that way!",
}

// Hey returns the Bob's reaction
func Hey(remark string) string {

	m, _ := regexp.MatchString(`^$|^\s*$`, remark)
	if m == true {
		return answers["fine"]
	}

	m, _ = regexp.MatchString(`^[A-Z0-9 .,!?]*[A-Z]+[A-Z0-9 .,!?]*\?[ ]*$`, remark)
	if m == true {
		return answers["calm"]
	}

	m, _ = regexp.MatchString(`^[A-Za-z0-9 .,!?]*[a-z0-9A-Z\:\) ]+[A-Za-z0-9 .,!?]*\?[ ]*$`, remark)
	if m == true {
		return answers["sure"]
	}

	m, _ = regexp.MatchString(`^[A-Z0-9 \W]*[A-Z]+[A-Z0-9 \W]*[ ]*$`, remark)
	if m == true {
		return answers["whoa"]
	}

	m, _ = regexp.MatchString(`.*`, remark)
	if m == true {
		return answers["whatever"]
	}

	return "Whatever."
}
