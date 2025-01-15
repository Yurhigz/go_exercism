package main

import "fmt"

var lyrics = map[string]string{
	"first":    "a Partridge in a Pear Tree",
	"second":   "two Turtle Doves",
	"third":    "three French Hens",
	"fourth":   "four Calling Birds",
	"fifth":    "five Gold Rings",
	"sixth":    "six Geese-a-Laying",
	"seventh":  "seven Swans-a-Swimming",
	"eighth":   "eight Maids-a-Milking",
	"ninth":    "nine Ladies Dancing",
	"tenth":    "ten Lords-a-Leaping",
	"eleventh": "eleven Pipers Piping",
	"twelfth":  "twelve Drummers Drumming",
}
var order = []string{
	"first", "second", "third", "fourth", "fifth",
	"sixth", "seventh", "eighth", "ninth", "tenth",
	"eleventh", "twelfth",
}

func Verse(i int) string {
	if i > 11 || i < 0 {
		return ""
	}
	gift := ""
	nbDay := order[i-1]
	if i == 1 {
		return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", nbDay, lyrics[order[i-1]])
	}
	for e := i - 1; e >= 0; e-- {
		if e == 0 {
			// Dernier cadeau (Partridge)
			gift += "and " + lyrics[order[e]]
		} else {
			gift += lyrics[order[e]] + ", "
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", nbDay, gift)
}

func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}

func main() {

	fmt.Println(Verse(1))
	fmt.Println(Verse(2))
	fmt.Println(Verse(3))
	// fmt.Println(Song())
}
