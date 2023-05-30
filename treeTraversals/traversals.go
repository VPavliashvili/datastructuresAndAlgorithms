package treetraversals

type node struct {
	value int
	left  *node
	right *node
}

type BinaryTree struct {
	root node
}

func CreateTree() BinaryTree {
	tree := BinaryTree{
		root: node{
			value: 7,
			left: &node{
				value: 23,
				left: &node{
					value: 5,
					left:  nil,
					right: nil,
				},
				right: &node{
					value: 4,
					left:  nil,
					right: nil,
				},
			},
			right: &node{
				value: 3,
				left: &node{
					value: 18,
					left:  nil,
					right: nil,
				},
				right: &node{
					value: 21,
					left:  nil,
					right: nil,
				},
			},
		},
	}
	return tree
}

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
