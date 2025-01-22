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
		if unicode.IsLetter(i) || unicode.IsDigit(i) {
			sb.WriteRune(unicode.ToLower(i))
		}
	}
	return sb.String()
}

func Encode(pt string) string {
	ptNormalized := normalize(pt)
	l := len(ptNormalized)
	c, r := int(math.Ceil(math.Sqrt(float64(l)))), int(math.Floor(math.Sqrt(float64(l))))

	if c*r < l {
		r = r + 1
	}

	NormalizedMessageList := make([]string, c)
	NormalizedMessage := make([]string, r)
	ptNormalized += strings.Repeat(" ", c*r-l)

	for i := 0; i < r; i++ {
		NormalizedMessage[i] = ptNormalized[i*c : (i+1)*c]
	}
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			NormalizedMessageList[i] += string(NormalizedMessage[j][i])
		}
	}

	return strings.Join(NormalizedMessageList, " ")
}

func main() {

	// fmt.Println(normalize("1, 2, 3 GO!"))
	// fmt.Println(Encode("If man was meant to stay on the ground, god would have given us roots."))
	// fmt.Println(Encode("Time is an illusion. Lunchtime doubly so."))
	// fmt.Println(Encode("1 2 3"))
	fmt.Println(Encode("Time is an illusion. Lunchtime doubly so."))
	// fmt.Println(Encode("We all know interspecies romance is weird."))
	// fmt.Println(Encode("12345678"))
}
