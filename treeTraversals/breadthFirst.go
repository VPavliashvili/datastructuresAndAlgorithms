package treetraversals

import "algorithmsAndDataStructures/queue"

// this method is aka level order tree traversal

func (bt BinaryTree) TraverseBreadthFirst() []int {
	queue := queue.Queue[node]{}
    res := []int{}

    queue.Enqueue(bt.root)
    for queue.Length() > 0 {
        curr := *queue.Dequeue()
        res = append(res, curr.value)

        if curr.left != nil {
            queue.Enqueue(*curr.left)
        }
        if curr.right != nil {
            queue.Enqueue(*curr.right)
        }
    }

	return res
}
