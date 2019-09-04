package isogram

import "sort"
import "strings"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// IsIsogram returns a boolean about a phrase is an isogram or not
func IsIsogram(phrase string) bool {
	cleanedString := strings.ReplaceAll(phrase, "-", "")
	cleanedString = strings.ReplaceAll(cleanedString, " ", "")
	loweredString := strings.ToLower(cleanedString)
	sortedString := sortString(loweredString)
	sortedRunes := []rune(sortedString)
	numOfRunes := len(sortedRunes)
	if numOfRunes == 0 {
		return true
	}
	print(string(sortedRunes))
	for i, c := range sortedRunes {
		if i < numOfRunes-1 && c == sortedRunes[i+1] {
			return false
		}
	}
	return true
}
