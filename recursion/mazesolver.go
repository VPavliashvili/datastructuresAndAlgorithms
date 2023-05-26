package recursion

import "algorithmsAndDataStructures/stack"

type Point struct {
	X int
	Y int
}

var dir = []Point{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: 0, Y: 1},
}

func walk(maze []string, wall string, end Point, cur Point, seen *[][]bool, path *stack.Stack[Point]) bool {

	/// base

	// if off the map
	if cur.Y < 0 || cur.Y >= len(maze) || cur.X < 0 || cur.X >= len(maze[0]) {
		return false
	}

	// if hit the wall
	if string(maze[cur.Y][cur.X]) == wall {
		return false
	}

	// got to end(solved)
	if end.X == cur.X && end.Y == cur.Y {
		path.Push(end)
		return true
	}

	// has already visited
	if (*seen)[cur.Y][cur.X] {
		return false
	}

	/// pre
	// pushing to path
	path.Push(cur)
	(*seen)[cur.Y][cur.X] = true

	/// recursion
	for i := 0; i < len(dir); i++ {
		x, y := dir[i].X, dir[i].Y
		next := Point{X: cur.X + x, Y: cur.Y + y}
		if walk(maze, wall, end, next, seen, path) {
			return true
		}
	}

	/// post
	// popping extra seen cells from path
	path.Pop()

	return false
}

func Solve(maze []string, wall string, start Point, end Point) []Point {
	cur := Point{X: start.X, Y: start.Y}

	sizex := len(maze[0])
	sizey := len(maze)
	seen := make([][]bool, sizey)
	for i := range seen {
		seen[i] = make([]bool, sizex)
	}

	// using stack instead of built in slices for more convenience
	var path = stack.Stack[Point]{}

	res := walk(maze, wall, end, cur, &seen, &path)
	if !res {
		panic("could not resolv maze")
	}

	var result []Point

	for i, size := 0, path.Length(); i < size; i++ {
		val := *path.Pop()
		result = append(result, val)
	}

	result = reverse(result)
	return result
}

// not mine
func reverse(input []Point) []Point {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
