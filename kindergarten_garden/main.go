package main

import (
	"errors"
	"fmt"
)

type Garden struct {
	diagram  string
	children []string
}

func containsDuplicate(liste []string) bool {
	for i := 0; i < len(liste); i++ {
		for j := i + 1; j < len(liste); j++ {
			if liste[i] == liste[j] {
				return true
			}
		}
	}
	return false
}

func sortNames(names []string) []string {
    sorted := make([]string, len(names))
    copy(sorted, names)
    
    for i := 0; i < len(sorted)-1; i++ {
        for j := 0; j < len(sorted)-i-1; j++ {
            if sorted[j] > sorted[j+1] {
                sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
            }
        }
    }
    return sorted
}

func validDiagram(diagram string) bool {
	for i := 0; i < len(diagram); i++ {
		if diagram[i] != '\n' && diagram[i] != 'V' && diagram[i] != 'R' && diagram[i] != 'C' && diagram[i] != 'G' {
			return false
		}
	}
	if diagram[0:1] == "\n" {
		return true
	}
	return false
}

func PlantsTranscription(plantsAbrev string) []string {
	fullPlants := []string{}
	for _, letter := range plantsAbrev {
		switch string(letter) {
		case "V":
			fullPlants = append(fullPlants, "violets")
		case "R":
			fullPlants = append(fullPlants, "radishes")
		case "C":
			fullPlants = append(fullPlants, "clover")
		case "G":
			fullPlants = append(fullPlants, "grass")
		}
	}

	return fullPlants
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if validDiagram(diagram) && len(diagram)-2 == len(children)*4 && !containsDuplicate(children) {
		sortedChildren := sortNames(children)
		return &Garden{diagram: diagram, children: sortedChildren}, nil
	}
	return nil, errors.New("Error")
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants := ""
	l := len(g.diagram) / 2
	for i, name := range g.children {
		if name == child {
			firstRowPot := string(g.diagram[2*i+1 : 2*i+3])
			secondRowPot := string(g.diagram[l+2*i+1 : l+2*i+3])
			plants += firstRowPot + secondRowPot
		}
	}
	if plants == "" {
		return []string{}, false
	}
	return PlantsTranscription(plants), true
}

func main() {

	test, _ := NewGarden("\nVVCG\nVVRC", []string{"Alice", "Bob"})
	fmt.Println(test.Plants("Alice"))
	students_2, _ := NewGarden("\nVVCG\nVVRC", []string{"Alice", "Bob"})
	fmt.Println(students_2.Plants("Bob"))
	students_3, _ := NewGarden("\nVVCCGG\nVVCCGG", []string{"Alice", "Bob", "Charlie"})
	fmt.Println(students_3.Plants("Charlie"))
	// test := []string{"Alice", "Bob", "Charlie"}
	// fmt.Println(test[0:3])
}
