package main

import "fmt"

var colorsMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Label(colors []string) string {
	multiplier := 1
	for i := 0; i < colorsMap[colors[2]]; i++ {
		multiplier *= 10
	}
	value := (colorsMap[colors[0]]*10 + colorsMap[colors[1]]) * multiplier
	if value < 1000 {
		return fmt.Sprintf("%d%s", value, " ohms")
	}
	if value < 1000000 {
		return fmt.Sprintf("%d%s", value/1000, " kiloohms")
	}
	if value < 100000000 {
		return fmt.Sprintf("%d%s", value/1000000, " megaohms")
	}
	return fmt.Sprintf("%d%s", value/1000000000, " gigaohms")
}

func main() {
	fmt.Println(Label([]string{"black", "brown", "red"}))
}
