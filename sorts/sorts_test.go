package sorts

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
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
		got := bubbleSort(item.list)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in bubbleSort\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
