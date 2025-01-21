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
	fmt.Println("ptNormalized : ", ptNormalized)
	l := len(ptNormalized)
	c, r := int(math.Ceil(math.Sqrt(float64(l)))), int(math.Floor(math.Sqrt(float64(l))))
	// fmt.Println("c : ", c, " r : ", r, " l : ", l)
	NormalizedMessage := make([]string, r)
	if r*c > l {
		ptNormalized += strings.Repeat(" ", c*r-l)
	}
	if r*c < l {
	}

	// fmt.Println(c*r - l)
	// fmt.Println("ptNormalized : ", ptNormalized)
	// NormalizedMessage est un tableau de string
	if r*c < l {
		for i := 0; i < r; i++ {
			NormalizedMessage[i] = ptNormalized[i*c : (i+1)*c]
		}
	}
	for i := 0; i < r; i++ {
		NormalizedMessage[i] = ptNormalized[i*c : (i+1)*c]
	}
	// fmt.Println("NormalizedMessage : ", NormalizedMessage)
	NormalizedMessageList := make([]string, c)
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
	fmt.Println(Encode("Time is an illusion. Lunchtime doubly so."))
}
