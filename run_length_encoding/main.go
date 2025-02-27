package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {
	encoded := ""
	occurrences := 1
	for i, l := range input {
		if i == len(input)-1 || input[i] != input[i+1] {
			if occurrences > 1 {
				encoded += strconv.Itoa(occurrences) + string(l)
				occurrences = 1
			} else {
				encoded += string(l)
			}
		} else {
			occurrences++
		}

	}
	return encoded
}

func RunLengthDecode(input string) string {
	decoded := ""
	count := ""

	for i := 0; i < len(input); i++ {
		if unicode.IsDigit(rune(input[i])) {
			count += string(input[i])
		} else {
			n := 1
			if count != "" {
				var err error
				n, err = strconv.Atoi(count)
				if err != nil {
					return ""
				}
				count = ""
			}
			for j := 0; j < n; j++ {
				decoded += string(input[i])
			}
		}
	}
	return decoded
}

func main() {
	fmt.Println(RunLengthEncode("AAAABBBCCDAA")) // 4A3B2C1D2A
	fmt.Println(RunLengthDecode("4A3B2C1D2A"))   // AAAABBBCCDAA
}
