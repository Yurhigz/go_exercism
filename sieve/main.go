package main

import "fmt"

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i <= (n/2)+1; i++ {
		if n%i == 0 {
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("True")
	return true
}

func Sieve(limit int) []int {
	primeNumbers := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}

func main() {
	isPrime(7)
	isPrime(6)
	fmt.Println(Sieve(13))

}
