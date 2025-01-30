package stringset

import (
	"fmt"
)

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

func (s Set) String() string {
	var e string
	i := 0
	for key := range s {
		if i < len(s)-1 {
			e += fmt.Sprintf("\"%v\"", key) + ", "
			i += 1
		} else {
			e += fmt.Sprintf("\"%v\"", key)
		}
	}
	return "{" + e + "}"
}

// Create a new set from a slice
func NewFromSlice(l []string) Set {
	set := New()
	for _, element := range l {
		set[element] = struct{}{}
	}
	return set
}

// check if the set does not contain any element
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// check for a specific key in the set
func (s Set) Has(elem string) bool {
	if _, keyExist := s[elem]; !keyExist {
		return false
	}
	return true
}

// add an element to the set
func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

// check if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	for element, _ := range s1 {
		exist := s2.Has(element)
		if !exist {
			return false
		}
	}
	return true
}

// Check if s1 and s2 have no element in common
func Disjoint(s1, s2 Set) bool {
	for element := range s1 {
		if s2.Has(element) {
			return false
		}
	}
	return true
}

// check if s1 = s2 which means same len and same elements
func Equal(s1, s2 Set) bool {
	if len(s1) == len(s2) && Subset(s1, s2) {
		return true
	}
	return false
}

// return a set with the element that s1 and s2 share
func Intersection(s1, s2 Set) Set {
	if s1.IsEmpty() && s2.IsEmpty() {
		return New()
	}
	s := New()
	for element := range s1 {
		if s2.Has(element) {
			s[element] = struct{}{}
		}
	}
	return s
}

// return a set with only the element of s1 which are not in s2
func Difference(s1, s2 Set) Set {
	if s1.IsEmpty() && s2.IsEmpty() {
		return New()
	}
	s := New()
	for element := range s1 {
		if !s2.Has(element) {
			s[element] = struct{}{}
		}
	}
	return s
}

// return a set with all the elements of both s1 and s2
func Union(s1, s2 Set) Set {
	if s1.IsEmpty() && s2.IsEmpty() {
		return New()
	}
	s := s2
	for element := range s1 {
		s.Add(element)
	}
	return s
}
