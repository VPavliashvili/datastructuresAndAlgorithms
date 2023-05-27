package sorts_test

import (
	"algorithmsAndDataStructures/sorts"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		list []int
		want []int
        function func([]int) []int
	}{
		{
			list: []int{5, 2, 10, 7, 220, 22, 1},
			want: []int{1, 2, 5, 7, 10, 22, 220},
            function: sorts.BubbleSort,
		},
		{
			list: []int{5, 2, 10, 7, 220, 22, 1, 150},
			want: []int{1, 2, 5, 7, 10, 22, 150, 220},
            function: sorts.BubbleSort,
		},
		{
			list: []int{1},
			want: []int{1},
            function: sorts.BubbleSort,
		},
		{
			list: []int{},
			want: []int{},
            function: sorts.BubbleSort,
		},
	}

	for _, item := range cases {
		got := item.function(item.list)
		want := item.want

        assert.Equal(t, want, got)
	}
}
