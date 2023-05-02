package binarySearch

import "testing"

func TestBs_list(t *testing.T) {
	cases := []struct {
		list   []int
		needle int
		want   bool
	}{
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			needle: 3,
			want:   true,
		},
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			needle: 12,
			want:   false,
		},
        {
        	list:   []int{1, 4, 13, 33, 55},
        	needle: 33,
        	want:   true,
        },
        {
        	list:   []int{-10, -7, -4, -3, -2},
        	needle: -4,
        	want:   true,
        },
	}

	for _, item := range cases {
		got := bs_list(item.list, item.needle)

		if got != item.want {
			t.Errorf("error\ngot\n%v\nwant\n%v", got, item.want)
		}
	}
}
