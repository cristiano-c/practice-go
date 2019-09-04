package reverse

// Reverse returns a reverse string of
func Reverse(word string) string {
	runes := []rune(word)
	for i := 0; i < int(len(runes)/2); i++ {
		if i < len(runes)-1 {
			runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
		}
	}
	return string(runes)
}
