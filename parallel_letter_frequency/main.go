package main

import (
	"fmt"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)

	numGoroutines := len(texts)
	FreqMap := make(FreqMap)

	for _, element := range texts {
		go func(t string) {
			c <- Frequency(t)

		}(element)
	}

	go func() {
		for i := 0; i < numGoroutines; i++ {

		}
		close(c)
	}()

	for freq := range c {
		for k, v := range freq {
			FreqMap[k] += v
		}
		// if len(FreqMap) == len(texts) {
		// 	break
		// }
	}
	return FreqMap
}

func main() {

	texts := []string{
		"hello",
		"world",
		"go concurrency",
	}
	fmt.Println(ConcurrentFrequency(texts))

}
