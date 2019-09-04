package raindrops

import "fmt"
import "strconv"

func Convert(input int) string {

	pling := ""
	plang := ""
	plong := ""
	noPlxng := ""

	if input%3 == 0 {
		pling = "Pling"
	}

	if input%5 == 0 {
		plang = "Plang"
	}

	if input%7 == 0 {
		plong = "Plong"
	}

	if pling == "" && plang == "" && plong == "" {
		noPlxng = strconv.Itoa(input)
	}

	return fmt.Sprintf("%s%s%s%s", pling, plang, plong, noPlxng)
}
