package algs

import (
	"math/rand/v2"
	. "mazes/internal/data"
)

func AldousBroder(g *Grid) {
	idx := func(x int, y int) int {
		return y*g.Width + x
	}

	x := rand.IntN(g.Width)
	y := rand.IntN(g.Height)

	visited := make([]bool, g.Height*g.Width)
	visited[idx(x, y)] = true
	remaning_squares := g.Height*g.Width - 1

	// W, E, N, S
	dir_x := []int{-1, 1, 0, 0}
	dir_y := []int{0, 0, -1, 1}
	step := func() {
		dir := rand.IntN(4)
		n_x := x + dir_x[dir]
		n_y := y + dir_y[dir]
		for {
			if n_x >= 0 && n_x < g.Width && n_y >= 0 && n_y < g.Height {
				break
			}
			dir = rand.IntN(4)
			n_x = x + dir_x[dir]
			n_y = y + dir_y[dir]
		}

		if !visited[idx(n_x, n_y)] {
			visited[idx(n_x, n_y)] = true
			remaning_squares -= 1
			switch dir {
			case 0:
				g.AddPath(x, y, "W")
			case 1:
				g.AddPath(x, y, "E")
			case 2:
				g.AddPath(x, y, "N")
			case 3:
				g.AddPath(x, y, "S")
			}
		}
		x = n_x
		y = n_y
	}

	for {
		step()
		if remaning_squares == 0 {
			return
		}
	}
}
