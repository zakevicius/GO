package word

import (
	"strings"
)

// UseCount returns how many times every word repeats in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns total number of words
func Count(s string) int {
	return len(strings.Fields(s))
}
