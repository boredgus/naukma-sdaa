package src

// Helper function to find the minimum node
func minimum(node *RBTreeNode) *RBTreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
