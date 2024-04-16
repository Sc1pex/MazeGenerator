package data

type Grid struct {
	Width  int
	Height int

	Cells [][]Cell
}

type Cell struct {
	North bool
	East  bool
	South bool
	West  bool
}

func NewGrid(width, height int) Grid {
	grid := Grid{
		Width:  width,
		Height: height,
	}

	grid.Cells = make([][]Cell, height)
	for i := range grid.Cells {
		grid.Cells[i] = make([]Cell, width)
		for j := range grid.Cells[i] {
			grid.Cells[i][j] = Cell{North: true, East: true, South: true, West: true}
		}
	}

	return grid
}

func (g *Grid) AddPath(from_x int, from_y int, direction string) {
	switch direction {
	case "N":
		g.Cells[from_y][from_x].North = false
		if from_y > 0 {
			g.Cells[from_y-1][from_x].South = false
		}
	case "E":
		g.Cells[from_y][from_x].East = false
		if from_x < g.Width-1 {
			g.Cells[from_y][from_x+1].West = false
		}
	case "S":
		g.Cells[from_y][from_x].South = false
		if from_y < g.Height-1 {
			g.Cells[from_y+1][from_x].North = false
		}
	case "W":
		g.Cells[from_y][from_x].West = false
		if from_x > 0 {
			g.Cells[from_y][from_x-1].East = false
		}
	}
}

func (g Grid) String() string {
	output := "+"
	for i := 0; i < g.Width; i++ {
		output += "---+"
	}
	output += "\n"

	for y := 0; y < g.Height; y++ {
		top := "|"
		bottom := "+"

		for x := 0; x < g.Width; x++ {
			cell := g.Cells[y][x]
			top += "   "
			if cell.East {
				top += "|"
			} else {
				top += " "
			}

			if cell.South {
				bottom += "---"
			} else {
				bottom += "   "
			}
			bottom += "+"
		}

		output += top + "\n"
		output += bottom + "\n"
	}

	return output
}

func (g Grid) StringNumbers(numFmt func(int) string) string {
	output := "+"
	for i := 0; i < g.Width; i++ {
		output += "---+"
	}
	output += "\n"

	for y := 0; y < g.Height; y++ {
		top := "|"
		bottom := "+"

		for x := 0; x < g.Width; x++ {
			cell := g.Cells[y][x]
			top += numFmt(y*g.Width + x)
			if cell.East {
				top += "|"
			} else {
				top += " "
			}

			if cell.South {
				bottom += "---"
			} else {
				bottom += "   "
			}
			bottom += "+"
		}

		output += top + "\n"
		output += bottom + "\n"
	}

	return output
}
