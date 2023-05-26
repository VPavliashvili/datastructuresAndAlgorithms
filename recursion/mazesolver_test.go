package recursion_test

import (
	"algorithmsAndDataStructures/recursion"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

var wall = "x"

var maze = []string{
	"xxxxxxxxxx x",
	"x        x x",
	"x        x x",
	"x xxxxxxxx x",
	"x          x",
	"x xxxxxxxxxx",
}

var answer = []recursion.Point{
	{X: 10, Y: 0},
	{X: 10, Y: 1},
	{X: 10, Y: 2},
	{X: 10, Y: 3},
	{X: 10, Y: 4},
	{X: 9, Y: 4},
	{X: 8, Y: 4},
	{X: 7, Y: 4},
	{X: 6, Y: 4},
	{X: 5, Y: 4},
	{X: 4, Y: 4},
	{X: 3, Y: 4},
	{X: 2, Y: 4},
	{X: 1, Y: 4},
	{X: 1, Y: 5},
}

func TestSolveMaze(t *testing.T) {

	cases := []struct {
		maze   []string
		answer []recursion.Point
		start  recursion.Point
		end    recursion.Point
	}{
		{
			maze:   maze,
			answer: answer,
			start:  recursion.Point{X: 10, Y: 0},
			end:    recursion.Point{X: 1, Y: 5},
		},
		// {
		// 	maze: []string{
		// 		"xxxxx xxx",
		// 		"x      xx",
		// 		"xxxxx xxx",
		// 		"xxxxx xxx",
		// 	},
		// 	answer: []recursion.Point{
		// 		{5, 0},
		// 		{5, 1},
		// 		{5, 2},
		// 		{5, 3},
		// 	},
		// 	start: recursion.Point{X: 5, Y: 0},
		// 	end:   recursion.Point{X: 5, Y: 3},
		// },
	}

	for _, cs := range cases {
		result := recursion.Solve(cs.maze, wall, cs.start, cs.end)

		assert.Equal(t, cs.answer, result)
	}
}
