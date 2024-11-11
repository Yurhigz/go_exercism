package main

import (
	"errors"
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	numbers := ""
	for _, e := range phoneNumber {
		if unicode.IsNumber(e) {
			numbers = numbers + string(e)
		}
	}
	if len(numbers) == 10 && (int(numbers[0]) > '1') && (int(numbers[3]) > '1') {
		return numbers, nil
	}

	if len(numbers) == 11 && (numbers[0] == '1') && (int(numbers[1]) > '1') && (int(numbers[4]) > '1') {
		return numbers[1:], nil
	}

	return "", errors.New("It requires more than 9 digit to be a proper number.")
}

func AreaCode(phoneNumber string) (string, error) {
	value, err := Number(phoneNumber)
	if err != nil {
		return "", errors.New("It requires more than 9 digit to be a proper number.")
	}
	return value[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	value, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + value[0:3] + ") " + value[3:6] + "-" + value[6:], nil
}

func main() {
	fmt.Println(Number("123456789"))
	fmt.Println(Number("1234567891"))
	fmt.Println(Number("12234567890"))
	fmt.Println(Number("223.456.7890"))
	fmt.Println(len("2234567890"))

}
