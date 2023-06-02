package treetraversals_test

import (
	treetraversals "algorithmsAndDataStructures/treeTraversals"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestInOrder(t *testing.T) {
	sut := treetraversals.CreateTree()

	got := sut.TraverseInOrder()
	want := []int{5, 23, 4, 7, 18, 3, 21}

	assert.Equal(t, want, got)
}

func TestPreOrder(t *testing.T) {
	sut := treetraversals.CreateTree()

	got := sut.TraversePreOrder()
	want := []int{7, 23, 5, 4, 3, 18, 21}

	assert.Equal(t, want, got)
}

func TestPostOrder(t *testing.T) {
	sut := treetraversals.CreateTree()

	got := sut.TraversePostOrder()
	want := []int{5, 4, 23, 18, 21, 3, 7}

	assert.Equal(t, want, got)
}

func TestBreadthFirst(t *testing.T) {
	sut := treetraversals.CreateTree()

	got := sut.TraverseBreadthFirst()
	want := []int{7, 23, 3, 5, 4, 18, 21}

	assert.Equal(t, want, got)
}
