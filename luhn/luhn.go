package luhn

import (
	"strconv"
	"strings"
)

func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	s = strings.TrimSpace(s)
	if len(s) <= 1 {
		return false
	}
	remainder := len(s) % 2
	sum := 0
	for key, val := range s {

		n, err := strconv.Atoi(string(val))
		if err != nil {
			return false
		}
		if key%2 == remainder {
			n = n * 2
		}
		if n > 9 {
			n = n - 9
		}
		sum += n
		// r[key] = n
	}
	if sum%10 == 0 {
		return true
	}
	return false

}
