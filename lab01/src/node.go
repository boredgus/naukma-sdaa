package src

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func (n *Node) String() string {
	return fmt.Sprint("{", n.value, "}")
}
