package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Operator(op string, a, b int) int {
	switch op {
	case "plus":
		return a + b
	case "minus":
		return a - b
	case "multiplied by":
		return a * b
	case "divided by":
		return a / b
	}
	return 0
}

func main() {

	question := []string{
		"What is 5?",
		"What is 1 plus 1?",
		"What is 53 plus 2?",
		"What is -1 plus -10?",
		"What is 4 minus -12?",
		"What is -3 multiplied by 25?",
		"What is 33 divided by -3?",
		"What is 15 plus 3 minus 5 multiplied by 2 divided by 4",
	}
	// What is (\d+)(?:\s*(plus|minus|multiplied by|divided by)\s*(\d+))*
	// What is (\d+)((plus|minus|multiplied by|divided)* (\d+)*)*
	re := regexp.MustCompile(`What is (-?\d+)(?:\s+(plus|minus|multiplied by|divided by)\s+(-?\d+))*\?`)
	matches := re.FindStringSubmatch(question[3])
	firstNumber, _ := strconv.Atoi(matches[1])
	operation := matches[2]
	secondNumber, _ := strconv.Atoi(matches[3])
	fmt.Println(firstNumber, operation, secondNumber)
	fmt.Println(Operator(operation, firstNumber, secondNumber))
}
