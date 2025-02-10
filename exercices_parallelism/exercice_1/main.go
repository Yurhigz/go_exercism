package main

import (
	"fmt"
	"sync"
	"time"
)

func printA(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A")
		time.Sleep(50 * time.Millisecond)
	}
}

func printB(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B")
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // On attend 2 Goroutines

	go printA(&wg)
	go printB(&wg)

	wg.Wait()
}
