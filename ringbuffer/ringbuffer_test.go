package ringbuffer_test

import (
	"algorithmsAndDataStructures/ringbuffer"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestRingBuffer(t *testing.T) {
	sut := ringbuffer.Create[int]()

	sut.Enqueue(5)
	assert.Equal(t, 5, *sut.Get(0))
	assert.Equal(t, 5, *sut.Dequeue())
	assert.Nil(t, sut.Dequeue())

	sut.Enqueue(42)
	sut.Enqueue(9)
	assert.Equal(t, 42, *sut.Dequeue())
	assert.Equal(t, 9, *sut.Dequeue())
	assert.Nil(t, sut.Dequeue())
	assert.Nil(t, sut.Get(1))

	sut.Enqueue(1)
	sut.Enqueue(2)
	sut.Enqueue(3)
	sut.Enqueue(4)
	sut.Enqueue(5)
	sut.Enqueue(6)

	assert.Equal(t, 3, *sut.Get(2))
	assert.Equal(t, 2, *sut.Get(1))
	assert.Equal(t, 1, *sut.Get(0))
	assert.Equal(t, 6, *sut.Get(5))
	assert.Nil(t, sut.Get(6))

	assert.Equal(t, 1, *sut.Dequeue())
	assert.Equal(t, 2, *sut.Dequeue())
	assert.Equal(t, 3, *sut.Dequeue())
	assert.Equal(t, 4, *sut.Dequeue())
	assert.Equal(t, 5, *sut.Dequeue())

	sut.Enqueue(42)
	sut.Enqueue(9)
	sut.Enqueue(12)
	assert.Equal(t, 12, *sut.Get(2))
	assert.Equal(t, 9, *sut.Get(1))
	assert.Equal(t, 42, *sut.Get(0))
}
