package main

import (
	"fmt"
)

func All(n int, s string) []string {
	if n <= 0 || n > len(s) {
		return nil
	}
	substrings := []string{}
	for i := 0; i <= len(s)-n; i++ {
		substrings = append(substrings, string(s[i:i+n]))
	}
	return substrings
}
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func main() {
	s := "49142"
	n := 1
	n1 := 3

	fmt.Println(All(n, s))
	fmt.Println(All(n1, s))
}
