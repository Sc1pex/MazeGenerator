package algs

import (
	"math/rand/v2"
	. "mazes/internal/data"
)

func Sidewinder(g *Grid) {
	for x := 0; x < g.Width-1; x++ {
		g.AddPath(x, 0, "E")
	}
	for y := 1; y < g.Height; y++ {
		run_len := 1
		for x := range g.Width {
			val := rand.IntN(2)
			if val == 0 || x == g.Width-1 {
				north_at := x - rand.IntN(run_len)
				g.AddPath(north_at, y, "N")
				run_len = 1
			} else {
				g.AddPath(x, y, "E")
				run_len++
			}
		}
	}
}
