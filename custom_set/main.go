package main

import "fmt"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]struct{}

func New() Set {
	return Set{}
}

// Create a new set from a slice
func NewFromSlice(l []string) Set {
	panic("Please implement the NewFromSlice function")
}

func (s Set) String() string {
	var e string
	for key, _ := range s {
		e += key
	}
	return "{" + e + "}"
}

// check if the set does not contain any element
func (s Set) IsEmpty() bool {
	panic("Please implement the IsEmpty function")
}

// check for a specific key in the set
func (s Set) Has(elem string) bool {
	panic("Please implement the Has function")
}

// add an element to the set
func (s Set) Add(elem string) {
	panic("Please implement the Add function")
}

// check if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	panic("Please implement the Subset function")
}

// Check if s1 and s2 have no element in common
func Disjoint(s1, s2 Set) bool {
	panic("Please implement the Disjoint function")
}

// check if s1 = s2 which means same len and same elements
func Equal(s1, s2 Set) bool {
	panic("Please implement the Equal function")
}

// return a set with the element that s1 and s2 share
func Intersection(s1, s2 Set) Set {
	panic("Please implement the Intersection function")
}

// return a set with only the element of s1 which are not in s2
func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

// return a set with all the elements of both s1 and s2
func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}

func main() {
	s := New()
	fmt.Println(s)
}
