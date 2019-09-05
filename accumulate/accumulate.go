package accumulate

// Accumulate iterates the collection applying the operator to each element
func Accumulate(collection []string, operator func(string) string) []string {
	var newCollection []string
	for _, word := range collection {
		newCollection = append(newCollection, operator(word))
	}
	return newCollection
}
