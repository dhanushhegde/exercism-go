//Package isogram implements function to check if a word is an isogram
package isogram

import "unicode"

// IsIsogram takes a string and returns a boolen if the string is an isogram
func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, val := range word {
		if !unicode.IsLetter(val) {
			continue
		}
		if _, ok := letters[unicode.ToUpper(val)]; ok {
			return false
		}
		letters[unicode.ToUpper(val)] = true

	}
	return true
}
