package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subjectSlice := []rune(strings.ToLower(subject))
	var returnSlice []string
	sort.Slice(subjectSlice, func(i, j int) bool {
		return subjectSlice[i] < subjectSlice[j]
	})

	for _, candidate := range candidates {
		candidateSlice := []rune(strings.ToLower(candidate))
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			break
		}
		sort.Slice(candidateSlice, func(i, j int) bool {
			return candidateSlice[i] < candidateSlice[j]
		})
		if string(subjectSlice) == string(candidateSlice) {
			returnSlice = append(returnSlice, candidate)
		}
	}
	return returnSlice

}
