package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	orderedList := make(map[string]int)
	for point, listLetter := range in {
		for _, letter := range listLetter {
			letters := strings.ToLower(letter)
			orderedList[string(letters)] = point
		}
	}
	return orderedList
}
