// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	yelled, _ := regexp.MatchString(`\A[A-Z1-9\s[:punct:]]+[A-Z][A-Z1-9\s[:punct:]]+\z`, remark)
	question, _ := regexp.MatchString(`\A.*\?\z`, remark)
	switch {
	case yelled && question:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case yelled:
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
