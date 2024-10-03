package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	total := 0
	number := strings.ReplaceAll(id, " ", "")
	if len(number) < 2 {
		return false
	}
	doubled := false
	for i := len(number) - 1; i >= 0; i-- {
		nb, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}
		if doubled {
			nb *= 2
			if nb > 9 {
				nb -= 9
			}
		}
		doubled = !doubled
		total += nb
	}

	return total%10 == 0
}
