package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node)

	for i, record := range records {

		if record.ID != i || (record.ID != 0 && record.Parent >= record.ID) || (record.ID == 0 && record.Parent > 0) {
			return nil, errors.New("données invalides")
		}

		nodes[record.ID] = &Node{
			ID: record.ID,
		}
		if record.ID != 0 {
			parent, exists := nodes[record.Parent]
			if !exists {
				return nil, errors.New("parent inexistant")
			}
			parent.Children = append(parent.Children, nodes[record.ID])
		}
	}
	return nodes[0], nil
}

// https://www.geeksforgeeks.org/binary-tree-data-structure/
// https://www.geeksforgeeks.org/introduction-to-tree-data-structure/
