package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	result := ""
	patternOne := regexp.MustCompile(`^(?:[aeiou]|xr|yt)`)
	patternTwo := regexp.MustCompile(`^([^aeiuo])+`)
	patternThree := regexp.MustCompile(`^([^aeiuoy]*qu)`)
	patternFour := regexp.MustCompile(`^([^aeiuo]+)y`)

	for i, word := range words {
		currentWord := ""
		if patternOne.MatchString(word) {
			currentWord += word + "ay"
		} else if match := patternFour.FindString(word); match != "" {
			currentWord = word[len(match)-1:] + match[:len(match)-1] + "ay"
		} else if match := patternThree.FindString(word); match != "" {
			currentWord = word[len(match):] + match + "ay"
		} else if match := patternTwo.FindString(word); match != "" {
			currentWord = word[len(match):] + match + "ay"
		}

		if i < len(words)-1 {
			result += currentWord + " "
		} else {
			result += currentWord
		}
	}
	return result
}

func main() {
	test1 := "rhythm"

	fmt.Printf("La réponse est == %s VS reponse attendue : ythmrhay", Sentence(test1))
}
