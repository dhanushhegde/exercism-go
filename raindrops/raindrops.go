package raindrops

import "fmt"

func Convert(n int) string {
	var pattern string
	if n%3 == 0 {
		pattern += "Pling"
	}
	if n%5 == 0 {
		pattern += "Plang"
	}
	if n%7 == 0 {
		pattern += "Plong"
	}
	if pattern == "" {
		pattern = fmt.Sprint(n)
	}
	return pattern

}
