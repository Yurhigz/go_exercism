package linkedlist

import "errors"

// Define List and Node types here.
type List struct {
	FirstNode *Node
	LastNode  *Node
}

type Node struct {
	Value    interface{}
	NextNode *Node
	PrevNode *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...interface{}) *List {
	if len(elements) <= 0 {
		return &List{}
	}
	if len(elements) == 1 {
		node := &Node{
			Value: elements[0],
		}
		return &List{
			FirstNode: node,
			LastNode:  node,
		}
	}
	l := List{}
	var prev *Node
	for i, value := range elements {
		node := Node{
			Value:    value,
			PrevNode: prev,
		}
		if prev != nil {
			prev.NextNode = &node
		}
		if i == 0 {
			l.FirstNode = &node
		} else if i == len(elements)-1 {
			l.LastNode = &node
		}
		prev = &node

	}
	return &l
}

func (n *Node) Next() *Node {
	return n.NextNode
}

func (n *Node) Prev() *Node {
	return n.PrevNode
}

func (l *List) Push(v interface{}) {
	node := &Node{
		Value:    v,
		PrevNode: l.LastNode,
		NextNode: nil,
	}
	if l.LastNode == nil {
		l.LastNode = node
		l.FirstNode = node
	} else {
		l.LastNode.NextNode = node
		l.LastNode = node
	}

}

func (l *List) Unshift(v interface{}) {
	node := &Node{
		Value:    v,
		PrevNode: nil,
		NextNode: l.FirstNode,
	}
	if l.FirstNode == nil {
		l.FirstNode = node
		l.LastNode = node
	} else {
		l.FirstNode.PrevNode = node
		l.FirstNode = node
	}

}

func (l *List) Shift() (interface{}, error) {
	if l.FirstNode == nil {
		return nil, errors.New("There is no first node")
	}

	node := l.FirstNode
	if node.NextNode == nil {
		l.FirstNode = nil
		l.LastNode = nil
		return node.Value, nil
	}
	l.FirstNode = l.FirstNode.NextNode
	l.FirstNode.PrevNode = nil
	return node.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.LastNode == nil {
		return nil, errors.New("There is no last node")
	}
	node := l.LastNode
	if node.PrevNode == nil {
		l.LastNode = nil
		l.FirstNode = nil
		return node.Value, nil
	}

	l.LastNode = l.LastNode.PrevNode
	l.LastNode.NextNode = nil
	return node.Value, nil
}

func (l *List) Reverse() {

	currNode := l.FirstNode
	for currNode != nil {

		currNode.NextNode, currNode.PrevNode = currNode.PrevNode, currNode.NextNode
		currNode = currNode.PrevNode
	}
	l.FirstNode, l.LastNode = l.LastNode, l.FirstNode
}

func (l *List) First() *Node {
	return l.FirstNode
}

func (l *List) Last() *Node {
	return l.LastNode
}
