package queue_test

import (
	"algorithmsAndDataStructures/queue"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestQueue(t *testing.T) {
	sut := queue.Queue[int]{}

	sut.Enqueue(5)
	sut.Enqueue(7)
	sut.Enqueue(9)

	assert.Equal(t, *sut.Dequeue(), 5)
	assert.Equal(t, sut.Length(), 2)

	sut.Enqueue(11)

	assert.Equal(t, *sut.Dequeue(), 7)
	assert.Equal(t, *sut.Dequeue(), 9)
	assert.Equal(t, *sut.Peek(), 11)
	assert.Equal(t, *sut.Dequeue(), 11)
	assert.Equal(t, sut.Dequeue(), (*int)(nil))
	assert.Equal(t, sut.Length(), 0)

	sut.Enqueue(69)
	assert.Equal(t, *sut.Peek(), 69)
	assert.Equal(t, sut.Length(), 1)
}
