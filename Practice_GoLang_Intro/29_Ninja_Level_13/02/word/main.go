//Package word provides custom functions for working with word and strings.
package word

import "strings"

//UseCount returns numbers of times the word used in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count returns number of words in a string.
func Count(s string) int {
	return len(strings.Fields(s))
}
