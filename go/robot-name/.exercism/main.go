package main

import (
	"fmt"
	"math/rand"
)

func randomName() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbers = []rune("0123456789")
	b := make([]rune, 5)
	for i := range b {
		if i < 2 {
			b[i] = letters[rand.Intn(len(letters))]
		} else {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
	}
	return string(b)
}

func main() {

	fmt.Println(randomName())
}
