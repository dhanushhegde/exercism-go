// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey returns a string reply depending on the input string
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if strings.HasSuffix(remark, "?") {
		if isUpper(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."

	} else if isUpper(remark) {
		return "Whoa, chill out!"
	} else if remark == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}

}

func isUpper(input string) bool {
	for _, val := range input {
		if unicode.IsLetter(val) && !unicode.IsUpper(val) {
			return false
		}

	}
	if strings.ToUpper(input) == strings.ToLower(input) {
		return false
	}
	return true
}
