package main

import (
	"fmt"
	"io"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	fmt.Println(reader, writer)
	return nil
}

func main() {

	r := strings.NewReader(`Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win`)
	b, _ := io.ReadAll(r)

	score := make(map[string][]int)

}
