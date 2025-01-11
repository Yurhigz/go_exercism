package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n <= 3:
		return true
	case n%2 == 0 || n%3 == 0:
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrt; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func Factors(n int64) []int64 {
    // Préallocation avec une taille estimée
    factors := make([]int64, 0, int(math.Log2(float64(n))))
    
    // Traiter les facteurs 2 séparément
    for n%2 == 0 {
        factors = append(factors, 2)
        n /= 2
    }
    
    // Traiter les facteurs 3 séparément
    for n%3 == 0 {
        factors = append(factors, 3)
        n /= 3
    }
    
    // Vérifier les nombres impairs jusqu'à sqrt(n)
    for i := int64(5); i*i <= n; i += 6 {
        for n%i == 0 {
            factors = append(factors, i)
            n /= i
        }
        for n%(i+2) == 0 {
            factors = append(factors, i+2)
            n /= i+2
        }
    }
    
    // Si n > 1, c'est un nombre premier
    if n > 1 {
        factors = append(factors, n)
    }
    
    return factors
}

func main() {
	fmt.Println(Factors(60))
}
