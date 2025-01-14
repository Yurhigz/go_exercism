package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	if len(s) == 0 {
		return ""
	}
	abreviation := ""
	isWord := true
	for i := 0; i < len(s); i++ {
		char := s[i]
		if isWord && isLetter(char) {
			abreviation += string(char)
		}
		isWord = char == ' ' || char == '-' || char == '_'
	}
	return abreviation
}

func isLetter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

// Solution alternative à base de replace/fields
func Abbreviate2(s string) (out string) {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)
	words := strings.Fields(s)
	for i := range words {
		out += string(words[i][0])
	}
	return strings.ToUpper(out)
}

const (
	inWord int = iota
	seekingWord
)

// Abbreviate should have a comment documenting it.
func Abbreviate3(s string) string {
	abbr := strings.Builder{}
	state := seekingWord

	for _, c := range s {
		switch {
		case state == seekingWord && unicode.IsLetter(c):
			abbr.WriteRune(unicode.ToUpper(c))
			state = inWord
		case state == inWord && !(unicode.IsLetter(c) || c == '\''):
			state = seekingWord
		}
	}

	return abbr.String()
}

// bench
// BenchmarkAcronym-2        881917              1245 ns/op              88 B/op         10 allocs/op

func main() {
	fmt.Println(Abbreviate("The Road _Not_ Taken"))
}
