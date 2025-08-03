package main

import (
	"errors"
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("char must be a capital letter")
	}
	if char == 'A' {
		return "A", nil
	}

	var diamond strings.Builder
	varChar := int(char) - int('A')
	// Triangle supérieur
	for i := 0; i < varChar; i++ {
		currentLetter := string(rune('A' + i))
		outterSpace := strings.Repeat(".", varChar-i)
		if i == 0 {
			diamond.WriteString(outterSpace + "A" + outterSpace + "\n")
		} else {
			innerSpace := strings.Repeat(".", i*2-1)
			diamond.WriteString(outterSpace + currentLetter + innerSpace + currentLetter + outterSpace + "\n")
		}
	}

	// Middle line
	diamond.WriteString(string(char) + strings.Repeat(".", 2*varChar-1) + string(char) + "\n")

	// Triangle inférieur
	for i := varChar - 1; i >= 0; i-- {
		currentLetter := string(rune('A' + i))
		outterSpace := strings.Repeat(".", varChar-i)
		if i == 0 {
			diamond.WriteString(outterSpace + currentLetter + outterSpace + "\n")
		} else {
			innerSpace := strings.Repeat(".", i*2-1)
			diamond.WriteString(outterSpace + currentLetter + innerSpace + currentLetter + outterSpace + "\n")
		}

	}

	return diamond.String(), nil
}

func main() {
	test_1, _ := Gen('B')
	test_2, _ := Gen('C')
	test_3, _ := Gen('D')
	test_4, _ := Gen('E')
	fmt.Println(test_1)
	fmt.Println(test_2)
	fmt.Println(test_3)
	fmt.Println(test_4)
}
