// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer hoses function that detials shares out of totals
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith returns a string detailing the share of the total
func ShareWith(name string) string {
	if strings.TrimSpace(name) == "" {
		return "One for you, one for me."
	}
	//return "One for " + name + ", one for me."
	return fmt.Sprintf("One for %s, one for me.", name)

}
