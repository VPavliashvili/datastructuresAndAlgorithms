package queue_test

import (
	"algorithmsAndDataStructures/queue"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestQueue(t *testing.T) {
	q := queue.Queue[int]{}

	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)

	assert.Equal(t, *q.Dequeue(), 5)
	assert.Equal(t, q.Length, 2)

	q.Enqueue(11)

	assert.Equal(t, *q.Dequeue(), 7)
	assert.Equal(t, *q.Dequeue(), 9)
	assert.Equal(t, *q.Peek(), 11)
	assert.Equal(t, *q.Dequeue(), 11)
	assert.Equal(t, q.Dequeue(), (*int)(nil))
	assert.Equal(t, q.Length, 0)

	q.Enqueue(69)
	assert.Equal(t, *q.Peek(), 69)
	assert.Equal(t, q.Length, 1)
}
