//Package reverse provides function that reverses the input string
package reverse

//Reverse takes a string and returns the reversed string
func Reverse(word string) string {
	// if len(strings.TrimSpace(word)) == 0 {
	// 	return ""
	// }

	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
