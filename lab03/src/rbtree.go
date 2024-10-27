package src

import (
	"errors"
)

var ValueDuplicationError = errors.New("value is duplicated")
var ValueNotFoundError = errors.New("value not found")

type RBTree struct {
	root *RBTreeNode
}

// rotate left
func (tree *RBTree) rotateLeft(x *RBTreeNode) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		tree.root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
}

// rotate right
func (tree *RBTree) rotateRight(y *RBTreeNode) {
	x := y.Left
	y.Left = x.Right
	if x.Right != nil {
		x.Right.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent == nil {
		tree.root = x
	} else if y == y.Parent.Right {
		y.Parent.Right = x
	} else {
		y.Parent.Left = x
	}
	x.Right = y
	y.Parent = x
}

// insert new node into the Red-Black Tree
func (tree *RBTree) Insert(value int) (*RBTreeNode, error) {
	newNode := &RBTreeNode{Value: value, IsRed: true}
	newRoot, err := tree.bstInsert(tree.root, newNode)
	if err != nil {
		return nil, err
	}

	tree.root = newRoot
	tree.fixAfterInsertion(newNode)

	return newNode, nil
}

// standard Binary Search Tree insert
func (tree *RBTree) bstInsert(root, node *RBTreeNode) (*RBTreeNode, error) {
	if root == nil {
		return node, nil
	}
	if root.Value == node.Value {
		return nil, ValueDuplicationError
	}
	if node.Value < root.Value {
		newRoot, err := tree.bstInsert(root.Left, node)
		if err != nil {
			return nil, err
		}
		root.Left = newRoot
		root.Left.Parent = root
	} else if node.Value > root.Value {
		newRoot, err := tree.bstInsert(root.Right, node)
		if err != nil {
			return nil, err
		}
		root.Right = newRoot
		root.Right.Parent = root
	}
	return root, nil
}

// fix Red-Black Tree properties after insertion
func (tree *RBTree) fixAfterInsertion(z *RBTreeNode) {
	for z != tree.root && z.Parent.IsRed {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right // uncle
			if y != nil && y.IsRed {   // case 1: Uncle is RED
				z.Parent.IsRed = false
				y.IsRed = false
				z.Parent.Parent.IsRed = true
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right { // case 2: z is right child
					z = z.Parent
					tree.rotateLeft(z)
				}
				// case 3: z is left child
				z.Parent.IsRed = false
				z.Parent.Parent.IsRed = true
				tree.rotateRight(z.Parent.Parent)
			}
		} else { // mirror cases
			y := z.Parent.Parent.Left // uncle
			if y != nil && y.IsRed {
				z.Parent.IsRed = false
				y.IsRed = false
				z.Parent.Parent.IsRed = true
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					tree.rotateRight(z)
				}
				z.Parent.IsRed = false
				z.Parent.Parent.IsRed = true
				tree.rotateLeft(z.Parent.Parent)
			}
		}
	}
	tree.root.IsRed = false
}

// Search a node by value in the Red-Black Tree
func (tree *RBTree) search(root *RBTreeNode, value int) *RBTreeNode {
	if root == nil || value == root.Value {
		return root
	}
	if value < root.Value {
		return tree.search(root.Left, value)
	}
	return tree.search(root.Right, value)
}

// Helper function to replace one subtree as a child of its parent with another subtree
func (tree *RBTree) transplant(u, v *RBTreeNode) {
	if u.Parent == nil {
		tree.root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

// Remove node from Red-Black Tree
func (tree *RBTree) Remove(value int) error {
	node := tree.search(tree.root, value)
	if node == nil {
		return ValueNotFoundError
	}

	isRed := node.IsRed
	var x *RBTreeNode
	var y = node

	if node.Left == nil {
		x = node.Right
		tree.transplant(node, node.Right)
	} else if node.Right == nil {
		x = node.Left
		tree.transplant(node, node.Left)
	} else {
		y = minimum(node.Right)
		isRed = y.IsRed
		x = y.Right
		if y.Parent == node {
			if x != nil {
				x.Parent = y
			}
		} else {
			tree.transplant(y, y.Right)
			y.Right = node.Right
			y.Right.Parent = y
		}
		tree.transplant(node, y)
		y.Left = node.Left
		y.Left.Parent = y
		y.IsRed = node.IsRed
	}

	if !isRed {
		tree.fixAfterDelete(x)
	}

	return nil
}

// fix Red-Black Tree properties after deletion
func (tree *RBTree) fixAfterDelete(x *RBTreeNode) {
	for x != tree.root && (x == nil || !x.IsRed) {
		if x == x.Parent.Left { // todo: fix potential nil pointer dereference
			w := x.Parent.Right // sibling
			if w.IsRed {
				w.IsRed = false
				x.Parent.IsRed = true
				tree.rotateLeft(x.Parent)
				w = x.Parent.Right
			}
			if (w.Left == nil || !w.Left.IsRed) && (w.Right == nil || !w.Right.IsRed) {
				w.IsRed = true
				x = x.Parent
			} else {
				if w.Right == nil || !w.Right.IsRed {
					w.Left.IsRed = false
					w.IsRed = true
					tree.rotateRight(w)
					w = x.Parent.Right
				}
				w.IsRed = x.Parent.IsRed
				x.Parent.IsRed = false
				if w.Right != nil {
					w.Right.IsRed = false
				}
				tree.rotateLeft(x.Parent)
				x = tree.root
			}
		} else { // Mirror case: x is right child
			w := x.Parent.Left
			if w.IsRed {
				w.IsRed = false
				x.Parent.IsRed = true
				tree.rotateRight(x.Parent)
				w = x.Parent.Left
			}
			if (w.Left == nil || !w.Left.IsRed) && (w.Right == nil || !w.Right.IsRed) {
				w.IsRed = true
				x = x.Parent
			} else {
				if w.Left == nil || !w.Left.IsRed {
					w.Right.IsRed = false
					w.IsRed = true
					tree.rotateLeft(w)
					w = x.Parent.Left
				}
				w.IsRed = x.Parent.IsRed
				x.Parent.IsRed = false
				if w.Left != nil {
					w.Left.IsRed = false
				}
				tree.rotateRight(x.Parent)
				x = tree.root
			}
		}
	}
	if x != nil {
		x.IsRed = false
	}
}
