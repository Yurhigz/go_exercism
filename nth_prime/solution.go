package prime
import (
    "errors"
    "math"
)
// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero) 
func isPrime(n int) bool {
    if n <= 1 {
        return false
    } else if n == 2 {
        return true
    } else if n % 2 == 0 {
        return false
    }
    sqrt := int(math.Sqrt(float64(n)))
    for i := 3; i <= sqrt; i += 2 {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func Nth(n int) (int, error) {
    rank := 0
    i := 0
	if n < 1 {
        return 0, errors.New("N cannot be equal or less than 0.")
    }
    if n == 1 {
        return 2, nil
    }
    for rank != n {
        if isPrime(i) {
            rank += 1
            i += 1
        } else {
            i += 1
        }
    }
    return i-1, nil
}
