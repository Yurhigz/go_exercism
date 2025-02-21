package tree

import "errors"

type Record struct {
	ID     int
	Parent *int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	var organizedNode *Node
	if len(records) == 0 {
		return nil, nil
	}
	for _, record := range records {
		if record.ID == 0 && record.Parent != nil {
			return nil, errors.New("root cannot have parents")
		}
	}

}

// https://www.geeksforgeeks.org/binary-tree-data-structure/
// https://www.geeksforgeeks.org/introduction-to-tree-data-structure/
