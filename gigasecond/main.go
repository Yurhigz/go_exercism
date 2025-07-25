// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	secondToAdd := 1000000000
	// I use the add function from time package to add one gigasecond which was previously transformed into a duration and then multiplied by time.Second
	return t.Add(time.Duration(secondToAdd) * time.Second)
}
