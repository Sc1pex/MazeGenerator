package algs

import (
	"math/rand/v2"
	. "mazes/internal/data"
)

func BinTree(g *Grid) {
	for y := range g.Height {
		for x := range g.Width {
			if y == 0 && x == g.Width-1 {
				continue
			}
			if y == 0 {
				g.AddPath(x, y, "E")
				continue
			}
			if x == g.Width-1 {
				g.AddPath(x, y, "N")
				continue
			}

			val := rand.IntN(2)
			if val == 0 {
				g.AddPath(x, y, "E")
			} else {
				g.AddPath(x, y, "N")
			}
		}
	}
}
