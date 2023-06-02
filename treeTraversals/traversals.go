package treetraversals

// in this file there are 3 main traversal methods
// inorder, preorder and postorder

func (bt BinaryTree) TraverseInOrder() []int {
	result := []int{}

	traverseInOrder(bt.root, &result)

	return result
}

func traverseInOrder(n node, data *[]int) {
	if n.left != nil {
		traverseInOrder(*n.left, data)
	}
	*data = append(*data, n.value)
	if n.right != nil {
		traverseInOrder(*n.right, data)
	}
}

func (bt BinaryTree) TraversePreOrder() []int {
	result := []int{}

	traversePreOrder(bt.root, &result)

	return result
}

func traversePreOrder(n node, data *[]int) {
	*data = append(*data, n.value)
	if n.left != nil {
		traversePreOrder(*n.left, data)
	}
	if n.right != nil {
		traversePreOrder(*n.right, data)
	}
}

func (bt BinaryTree) TraversePostOrder() []int {
	result := []int{}

	traversePostOrder(bt.root, &result)

	return result
}

func traversePostOrder(n node, data *[]int) {
	if n.left != nil {
		traversePostOrder(*n.left, data)
	}
	if n.right != nil {
		traversePostOrder(*n.right, data)
	}
	*data = append(*data, n.value)
}
