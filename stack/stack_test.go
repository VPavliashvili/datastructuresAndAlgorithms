package stack_test

import (
	"algorithmsAndDataStructures/stack"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestStack(t *testing.T) {
    sut := stack.Stack[int]{}

    sut.Push(5)
    sut.Push(7)
    sut.Push(9)

    assert.Equal(t, *sut.Pop(), 9)
    assert.Equal(t, sut.Length(), 2)

    sut.Push(11)
    assert.Equal(t, *sut.Pop(), 11)
    assert.Equal(t, *sut.Pop(), 7)
    assert.Equal(t, *sut.Peek(), 5)
    assert.Equal(t, *sut.Pop(), 5)
    assert.Equal(t, sut.Pop(), (*int)(nil))

    sut.Push(69)
    assert.Equal(t, *sut.Peek(), 69)
    assert.Equal(t, sut.Length(), 1)

    sut.Pop()
    sut.Pop()
    sut.Pop()

    assert.Nil(t, sut.Peek())
    assert.Equal(t, sut.Length(), 0)
}
