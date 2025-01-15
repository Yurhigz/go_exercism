package main

import "fmt"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if s.Length() == 0 {
		return s
	}
	var trueList IntList
	for _, v := range s {
		if fn(v) {
			trueList = append(trueList, v)
		}
	}
	return trueList
}

func (s IntList) Length() int {
	len := 0
	for i, _ := range s {
		len = i + 1
	}
	return len
}

func (s IntList) Map(fn func(int) int) IntList {
	if s.Length() == 0 {
		return s
	}
	var newList IntList
	for _, v := range s {
		newList = append(newList, fn(v))
	}
	return newList
}

func (s IntList) Reverse() IntList {
	if s.Length() == 0 {
		return s
	}
	var newList IntList
	for i := s.Length() - 1; i >= 0; i-- {
		newList = append(newList, s[i])
	}
	return newList
}

func (s IntList) Append(lst IntList) IntList {
	if s.Length() == 0 {
		return lst
	}
	if lst.Length() == 0 {
		return s
	}

	totalLen := s.Length() + lst.Length()
	newList := make(IntList, totalLen)
	copy(newList, []int(s))
	copy(newList[s.Length():], []int(lst))

	return newList

}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}
	return s
}

func main() {
	var list IntList = []int{1, 2, 3, 4, 5}
	var list2 IntList = []int{6, 7, 8, 9, 10}
	newList := list.Append(list2)
	copy(newList, list)
	fmt.Println(list.Length())
	fmt.Println(list2.Length())
	fmt.Println(newList)
}
