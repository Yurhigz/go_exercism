// Package isbn is a small library for validating input as an isbn.
package isbn

func getDigit(chr rune) (digit int, ok bool) {
	if chr < '0' || chr > '9' {
		return digit, ok
	}
	return int(chr - '0'), true
}

// IsValidISBN calculates the validity of the input as an ISBN.
func IsValidISBN(input string) bool {
	pos := 10
	sum := 0

	for _, chr := range input {
		if digit, ok := getDigit(chr); ok {
			sum += (digit * pos)
			pos--
		} else if pos == 1 && chr == 'X' {
			sum += 10
			pos--
		} else if chr != '-' {
			return false
		}
	}
	return pos == 0 && sum%11 == 0
}