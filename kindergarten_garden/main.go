package main

import "fmt"

// Define the Garden type here.
type Garden struct {
	diagram  string
	children []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

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
	if diagram[0:1] == "\n" && len(diagram)-2 == len(children)*4 {
		return &Garden{diagram: diagram, children: children}, nil
	}
	return nil, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants := ""
	for i, name := range g.children {

		if name == child {
			firstRowPot := string(g.diagram[i+1 : i+3])
			secondRowPot := string(g.diagram[len(g.diagram)/2+i+1 : len(g.diagram)/2+i+3])
			// fmt.Printf("Voici les deux premières plantes de la personne %s : %s \n", name, string(g.diagram[i+1:i+3]))
			// fmt.Printf("Voici les deux dernières plantes de la personne %s : %s \n", name, string(g.diagram[len(g.diagram)/2+i+1:len(g.diagram)/2+i+3]))
			plants += firstRowPot + secondRowPot
		}
	}
	return PlantsTranscription(plants), true
}

func main() {

	test, _ := NewGarden("\nVVCG\nVVRC", []string{"Alice", "Bob"})
	fmt.Println(test.Plants("Alice"))
	// NewGarden("RC\nGG", []string{"Alice"})
	// NewGarden("\nRCCC\nGG", []string{""})
	// NewGarden("\nRCC\nGGC", []string{"Alice"})
	// NewGarden("\nVRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV", []string{"Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"})

}
