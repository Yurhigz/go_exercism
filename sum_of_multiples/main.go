package main

import (
	"fmt"
	"slices"
)

func SumMultiples(limit int, divisors ...int) int {
	divisorsNotSorted := []int{}
	sum := 0
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor > 0 {
				if i%divisor == 0 {
					divisorsNotSorted = append(divisorsNotSorted, i)
				}
			}
		}
	}
	divisorsSorted := slices.Compact(divisorsNotSorted)
	for _, i := range divisorsSorted {
		sum += i
	}
	return sum
}

func main() {
	limit := 20
	divisors := []int{3, 5}
	sum := SumMultiples(limit, divisors...)
	fmt.Println(sum)
}
