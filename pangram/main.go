package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	alphabet := make(map[string]int)
	for _, letter := range strings.ToLower(input) {
		if unicode.IsLetter(letter) {
			_, exists := alphabet[string(letter)]
			if !exists {
				alphabet[string(letter)] += 1
			}
		}
	}
	return len(alphabet) == 26
}
