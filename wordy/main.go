package main

import (
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
		if b == 0 {
			return 0
		}
		return a / b
	default:
		return 0
	}
}

func Answer(question string) (int, bool) {

	re := regexp.MustCompile(`What is (-?\d+)(?: (plus|minus|multiplied by|divided by) (-?\d+))*\?`)

	if re.MatchString(question) {

		reSubNumbers, reSubOperators := regexp.MustCompile(`-?\d+`), regexp.MustCompile(`(plus|minus|multiplied by|divided by)`)
		numbers, operators := reSubNumbers.FindAllString(question, -1), reSubOperators.FindAllString(question, -1)
		results, _ := strconv.Atoi(numbers[0])

		if len(numbers) == 1 {
			return results, true
		}

		for i := 1; i < len(numbers); i++ {
			number, _ := strconv.Atoi(numbers[i])
			operator := operators[i-1]
			results = Operator(operator, results, number)
		}
		return results, true

	} else {
		return 0, false
	}

}

func main() {
	tests := []string{
		"What is 5?",
		"What is -1 plus -10?",
		"What is 5 plus 13?",
		"What is 7 minus 5?",
		"What is 6 multiplied by 4?",
		"What is 25 divided by 5?",
		"What is 3 plus 2 multiplied by 3?",
		"What is 52 cubed?",
		"Who is the President?",
		"What is 1 plus plus 2?",
		"What is -12 divided by 2 divided by -3?",
		"What is -3 plus 7 multiplied by -2?",
	}
	for _, question := range tests {
		result, ok := Answer(question)
		if ok {
			println(question, "=>", result)
		} else {
			println(question, "=>", "I don't know")
		}
	}
}
