package twoballs

import "testing"

func TestTwoballs(t *testing.T) {
	cases := []struct {
		input []bool
		want  int //distance
	}{
		{
			// is not sqrt int len == 8
			input: []bool{false, false, false, false, false, false, true, true},
			want:  7,
		},
		{
			// is sqrt int len == 9
			input: []bool{false, false, false, false, false, false, false, true, true},
			want:  8,
		},
		{
			input: []bool{true, true, true},
			want:  1,
		},
        {
        	input: []bool{false, true, true},
        	want:  2,
        },
        {
        	input: []bool{false, false, true},
        	want:  3,
        },
	}

	for _, item := range cases {
		got := twoballs(item.input)

		if got != item.want {
			t.Errorf("error\ngot\n%v\nwant\n%v", got, item.want)
		}
	}
}
