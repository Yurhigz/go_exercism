package main

import (
	"errors"
	"fmt"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 || outputBase < 2 {
		return []int{}, errors.New("inputBase and outputBase must be >= 2")
	}
	for _, number := range inputDigits {
		if number < 0 || number > inputBase {
			return inputDigits, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}
	outputDigits := []int{}
	for i, number := range inputDigits {

	}
	panic("Please implement the ConvertToBase function")
}

func main() {
	testCases := []struct {
		description   string
		inputBase     int
		inputDigits   []int
		outputBase    int
		expected      []int
		expectedError string
	}{
		{
			description:   "single bit one to decimal",
			inputBase:     2,
			inputDigits:   []int{1},
			outputBase:    10,
			expected:      []int{1},
			expectedError: "",
		},
		{
			description:   "binary to single decimal",
			inputBase:     2,
			inputDigits:   []int{1, 0, 1},
			outputBase:    10,
			expected:      []int{5},
			expectedError: "",
		},
		{
			description:   "single decimal to binary",
			inputBase:     10,
			inputDigits:   []int{5},
			outputBase:    2,
			expected:      []int{1, 0, 1},
			expectedError: "",
		},
		{
			description:   "binary to multiple decimal",
			inputBase:     2,
			inputDigits:   []int{1, 0, 1, 0, 1, 0},
			outputBase:    10,
			expected:      []int{4, 2},
			expectedError: "",
		},
		{
			description:   "decimal to binary",
			inputBase:     10,
			inputDigits:   []int{4, 2},
			outputBase:    2,
			expected:      []int{1, 0, 1, 0, 1, 0},
			expectedError: "",
		},
		{
			description:   "trinary to hexadecimal",
			inputBase:     3,
			inputDigits:   []int{1, 1, 2, 0},
			outputBase:    16,
			expected:      []int{2, 10},
			expectedError: "",
		},
		{
			description:   "hexadecimal to trinary",
			inputBase:     16,
			inputDigits:   []int{2, 10},
			outputBase:    3,
			expected:      []int{1, 1, 2, 0},
			expectedError: "",
		},
		{
			description:   "15-bit integer",
			inputBase:     97,
			inputDigits:   []int{3, 46, 60},
			outputBase:    73,
			expected:      []int{6, 10, 45},
			expectedError: "",
		},
		{
			description:   "empty list",
			inputBase:     2,
			inputDigits:   []int{},
			outputBase:    10,
			expected:      []int{0},
			expectedError: "",
		},
		{
			description:   "single zero",
			inputBase:     10,
			inputDigits:   []int{0},
			outputBase:    2,
			expected:      []int{0},
			expectedError: "",
		},
		{
			description:   "multiple zeros",
			inputBase:     10,
			inputDigits:   []int{0, 0, 0},
			outputBase:    2,
			expected:      []int{0},
			expectedError: "",
		},
		{
			description:   "leading zeros",
			inputBase:     7,
			inputDigits:   []int{0, 6, 0},
			outputBase:    10,
			expected:      []int{4, 2},
			expectedError: "",
		},
		{
			description:   "input base is one",
			inputBase:     1,
			inputDigits:   []int{0},
			outputBase:    10,
			expected:      []int(nil),
			expectedError: "input base must be >= 2",
		},
	}

	for _, test := range testCases {
		fmt.Printf("\nTest: %s\n", test.description)
		result, err := ConvertToBase(test.inputBase, test.inputDigits, test.outputBase)

		// Vérification des erreurs
		if test.expectedError != "" {
			if err == nil {
				fmt.Printf("❌ Expected error: %s, but got no error\n", test.expectedError)
			} else if err.Error() != test.expectedError {
				fmt.Printf("❌ Expected error: %s, but got: %v\n", test.expectedError, err)
			} else {
				fmt.Printf("✅ Got expected error: %s\n", err)
			}
			continue
		}

		// Vérification des résultats
		if err != nil {
			fmt.Printf("❌ Unexpected error: %v\n", err)
			continue
		}

		// Comparaison des résultats
		matches := len(result) == len(test.expected)
		if matches {
			for i := range result {
				if result[i] != test.expected[i] {
					matches = false
					break
				}
			}
		}

		if matches {
			fmt.Printf("✅ Test passed! Result: %v\n", result)
		} else {
			fmt.Printf("❌ Expected: %v, but got: %v\n", test.expected, result)
		}
	}
}
