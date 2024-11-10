package bottlesong

import (
	"strings"
)

func verse(nbBottles int) []string {
	numbers := map[int]string{
		0:  "No green bottles",
		1:  "One green bottle",
		2:  "Two green bottles",
		3:  "Three green bottles",
		4:  "Four green bottles",
		5:  "Five green bottles",
		6:  "Six green bottles",
		7:  "Seven green bottles",
		8:  "Eight green bottles",
		9:  "Nine green bottles",
		10: "Ten green bottles"}

	firstLines := numbers[nbBottles] + " hanging on the wall,"
	secondLine := "And if one green bottle should accidentally fall,"
	thirdLine := "There'll be " + strings.ToLower(numbers[nbBottles-1]) + " hanging on the wall."

	return []string{firstLines, firstLines, secondLine, thirdLine}
}

func Recite(startBottles, takeDown int) []string {
	lyrics := []string{}

	if takeDown <= 1 {
		for _, element := range verse(startBottles) {
			lyrics = append(lyrics, element)
		}
		return lyrics
	}

	for i := 0; i < takeDown; i++ {
		for _, element := range verse(startBottles - i) {
			lyrics = append(lyrics, element)
		}
		if i < takeDown-1 {
			lyrics = append(lyrics, "")
		}
	}
	return lyrics
}
