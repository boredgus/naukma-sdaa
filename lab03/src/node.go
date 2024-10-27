package src

import "fmt"

//type Color int
//
//func (c Color) String() string {
//	if c == 0 {
//		return "black"
//	}
//	return "red"
//}
//
//const (
//	BLACK_СOLOR Color = iota
//	RED_COLOR
//)

type RBTreeNode struct {
	IsRed  bool
	Value  int
	Parent *RBTreeNode
	Left   *RBTreeNode
	Right  *RBTreeNode
}

//func NewRBTreeNode(value int) *RBTreeNode {
//	return &RBTreeNode{
//		IsRed: false,
//		value: value,
//	}
//}

//func (n *RBTreeNode) SetParent(node *RBTreeNode) {
//	n.Parent = node
//}
//
//func (n *RBTreeNode) SetLeft(node *RBTreeNode) {
//	n.Left = node
//}
//
//func (n *RBTreeNode) SetRight(node *RBTreeNode) {
//	n.Right = node
//}

//func (n *RBTreeNode) Value() int {
//	return n.value
//}

func (n *RBTreeNode) String() string {
	color := "black"
	if n.IsRed {
		color = "red"
	}
	return fmt.Sprintf("{%v %s}", n.Value, color)
}
