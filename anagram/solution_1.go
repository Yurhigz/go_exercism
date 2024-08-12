// Package anagram contains a solution to the anagram Exercism exercise.
package anagram

import (
	"sort"
	"strings"
)

// Detect determines which words in cases are anagrams of the subject.
func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0)

	subject = strings.ToLower(subject)

	for _, candidate := range candidates {
		c := strings.ToLower(candidate)

		if isAnagram(subject, c) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

// isAnagram determines whether a and b are anagrams of each other.
func isAnagram(a, b string) bool {
	return a != b && sortString(a) == sortString(b)
}

// sortString sorts a string lexicographically in non-decreasing order.
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}