package arraylist_test

import (
	"algorithmsAndDataStructures/arraylist"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestArrayList(t *testing.T) {
	sut := arraylist.Create[int]()

	sut.Append(5)
	sut.Append(7)
	sut.Append(9)

	assert.Equal(t, 9, sut.ElementAt(2))
	assert.Equal(t, 7, sut.RemoveAt(1))
	assert.Equal(t, 2, sut.Length())

	sut.Append(11)
	assert.Equal(t, 9, sut.RemoveAt(1))
	assert.Nil(t, sut.Remove(9))
	assert.Equal(t, 5, sut.RemoveAt(0))
	assert.Equal(t, 11, sut.RemoveAt(0))
	assert.Equal(t, 0, sut.Length())

    sut.Append(10)
    assert.Equal(t, 10, *sut.Remove(10))
    
	sut.Prepend(5)
	sut.Prepend(7)
	sut.Prepend(9)

	assert.Equal(t, 5, sut.ElementAt(2))
	assert.Equal(t, 9, sut.ElementAt(0))
	assert.Equal(t, 9, *sut.Remove(9))
	assert.Equal(t, 2, sut.Length())
	assert.Equal(t, 7, sut.ElementAt(0))
}
