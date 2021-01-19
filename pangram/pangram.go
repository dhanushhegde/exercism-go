package pangram

import (
	"strings"
)

func IsPangram(input string) bool {

	var alphabetMap = map[string]bool{
		"q": false,
		"w": false,
		"e": false,
		"r": false,
		"t": false,
		"y": false,
		"u": false,
		"i": false,
		"o": false,
		"p": false,
		"a": false,
		"s": false,
		"d": false,
		"f": false,
		"g": false,
		"h": false,
		"j": false,
		"k": false,
		"l": false,
		"z": false,
		"x": false,
		"c": false,
		"v": false,
		"b": false,
		"n": false,
		"m": false,
	}

	input = strings.Replace(strings.ToLower(input), " ", "", -1)

	// for _, val := range input {
	// 	alphabetMap[string(val)] = true
	// }
	// for _, val := range alphabetMap {

	// 	if !val {
	// 		return false
	// 	}
	// }
	// return true

	for _, val := range input {
		_, ok := alphabetMap[string(val)]
		if ok {
			delete(alphabetMap, string(val))
		}
	}
	if len(alphabetMap) > 0 {
		return false
	}

	return true
}
