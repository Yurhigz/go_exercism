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
	done := make(chan bool)
	numGoroutines := len(texts)
	FreqMap := make(FreqMap)

	for _, element := range texts {
		go func(t string) {
			c <- Frequency(t)
			done <- true
		}(element)
	}

	go func() {
		for i := 0; i < numGoroutines; i++ {
			<-done // Attendre la fin d'une Goroutine
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

// ConcurrentFrequency concurrently counts the frequency of each rune in each item in a given list of texts
// and returns this data as a FreqMap.
func ConcurrentFrequency2(list []string) FreqMap {
	ch := make(chan FreqMap)
	for _, s := range list {
		go func(s string) { ch <- Frequency(s) }(s)
	}
	result := FreqMap{}
	for range list {
		for letter, count := range <-ch {
			result[letter] += count
		}
	}
	return result
}

func main() {

	texts := []string{
		"hello",
		"world",
		"go concurrency",
	}
	fmt.Println(ConcurrentFrequency(texts))

}
