package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type performances struct {
	team        string
	matchPlayed int
	wins        int
	even        int
	loss        int
	totalPoint  int
}

func Tally(reader io.Reader) error {
	r := bufio.NewReader(reader)
	// data := make([]byte, 1000)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}

// func Tally(reader io.Reader, writer io.Writer) error {
// 	r := bufio.NewReader(reader)
// 	for {

// 	}
// 	return nil
// }

func main() {

	// 	r := strings.NewReader(`Allegoric Alaskians;Blithering Badgers;win
	// Devastating Donkeys;Courageous Californians;draw
	// Devastating Donkeys;Allegoric Alaskians;win
	// Courageous Californians;Blithering Badgers;loss
	// Blithering Badgers;Devastating Donkeys;loss
	// Allegoric Alaskians;Courageous Californians;win`)
	// 	b, _ := Read(r)

	// Exemple : un io.Reader simulé
	data := `Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win`
	Tally(strings.NewReader(data))
}
