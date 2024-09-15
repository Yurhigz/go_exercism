package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number > 64 || number < 1 {
		return 0, errors.New("This square number does not exist.")
	}
	return uint64(math.Pow(2.0, float64(number-1))), nil
}

func Total() uint64 {
	var total uint64 = 0
	for i := 1; i <= 64; i++ {
		s, e := Square(i)
		if e != nil {
			return 0
		}
		total += s
	}
	return total
}
