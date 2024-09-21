// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := []string{}
	if len(rhyme) > 0 {
		for i := 0; i < len(rhyme)-1; i++ {
			text := fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1])
			proverb = append(proverb, text)
		}
		finalSentence := fmt.Sprintf("And all for the want of a %v.", rhyme[0])
		proverb = append(proverb, finalSentence)
		return proverb
	}
	return rhyme
}
