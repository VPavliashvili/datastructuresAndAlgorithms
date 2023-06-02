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
