// Package acronym provides funtion to generate acronyms from string inputs
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns string acronyms of input string
func Abbreviate(s string) string {
	flag := true
	var abbreviation string
	for _, v := range s {

		if unicode.IsLetter(v) {
			if flag {
				flag = false
				abbreviation += strings.ToUpper(string(v))
			}
		} else if unicode.IsSpace(v) || string(v) == "-" {
			flag = true
		}
	}

	// inputString := strings.Split(s, " ")
	// var abbreviation string
	// for _, val := range inputString {
	// 	if firstletter := firstLetter(val); firstletter != "" {
	// 		abbreviation += firstletter
	// 	}
	// }
	return abbreviation
}

func firstLetter(input string) string {
	for _, val := range input {
		if unicode.IsLetter(val) {
			return string(val)
		}
	}
	return ""
}
