package main

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	lyrics := []string{}
	numbers := []string{
		"One green bottle",
		"Two green bottles",
		"Three green bottles",
		"Four green bottles",
		"Five green bottles",
		"Six green bottles",
		"Seven green bottles",
		"Eight green bottles",
		"Nine green bottles",
		"Ten green bottles"}

	for i := startBottles; i > takeDown; i-- {
		nextBottles := "no green bottles"
		if i > 1 {
			nextBottles = strings.ToLower(numbers[i-2])
		}
		verse := fmt.Sprintf(
			"%[1]s hanging on the wall,\n%[1]s hanging on the wall,\nAnd if one green bottle should accidentally fall,\nThere'll be %[2]s hanging on the wall.\n",
			numbers[i-1], nextBottles,
		)
		lyrics = append(lyrics, verse)
	}
	return lyrics
}

func main() {
	fmt.Println(Recite(10, 1))

}
