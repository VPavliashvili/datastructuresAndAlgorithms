package sorts_test

import (
	"algorithmsAndDataStructures/sorts"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		list []int
		want []int
	}{
		{
			list: []int{5, 2, 10, 7, 220, 22, 1},
			want: []int{1, 2, 5, 7, 10, 22, 220},
		},
		{
			list: []int{5, 2, 10, 7, 220, 22, 1, 150},
			want: []int{1, 2, 5, 7, 10, 22, 150, 220},
		},
		{
			list: []int{1},
			want: []int{1},
		},
		{
			list: []int{},
			want: []int{},
		},
	}

	for _, item := range cases {
		got := sorts.BubbleSort(item.list)
		want := item.want

		assert.Equal(t, want, got)
	}

}

func TestQuickSort(t *testing.T) {
	cases := []struct {
		list []int
		want []int
	}{
		{
			list: []int{5, 2, 10, 7, 220, 22, 1},
			want: []int{1, 2, 5, 7, 10, 22, 220},
		},
		{
			list: []int{5, 2, 10, 7, 220, 22, 1, 150},
			want: []int{1, 2, 5, 7, 10, 22, 150, 220},
		},
		{
			list: []int{1},
			want: []int{1},
		},
		{
			list: []int{},
			want: []int{},
		},
	}

	for _, item := range cases {
		got := sorts.QuickSort(item.list)
		want := item.want

		assert.Equal(t, want, got)
	}
}
