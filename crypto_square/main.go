package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func normalize(pt string) string {
	var sb strings.Builder
	for _, i := range pt {
		if unicode.IsLetter(i) {
			sb.WriteRune(unicode.ToLower(i))
		}
	}
	return sb.String()
}

func Encode(pt string) string {
	ptNormalized := normalize(pt)
	encoded := ""
	l := len(ptNormalized)
	c := int(math.Ceil(math.Sqrt(float64(l))))
	r := int(math.Floor(math.Sqrt(float64(l))))
	fmt.Println("l: ", l)
	fmt.Println("c: ", c)
	fmt.Println("r: ", r)
	if r*c > l {
		for i := 0; i < r; i++ {
			encoded += ptNormalized[i*r:i*r+r-1] + " "
		}
	}
	return encoded
}
func main() {

	// fmt.Println(normalize("1, 2, 3 GO!"))
	fmt.Println(normalize("If man was meant to stay on the ground, god would have given us roots."))
	fmt.Println(len(normalize("If man was meant to stay on the ground, god would have given us roots.")))
	fmt.Println(Encode("If man was meant to stay on the ground, god would have given us roots."))
}
