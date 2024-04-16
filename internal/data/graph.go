package data

type Graph struct {
	Nodes int
	Edges [][]int
}

func NewGraph(nodes int) Graph {
	return Graph{
		Nodes: nodes,
		Edges: make([][]int, nodes),
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.Edges[from] = append(g.Edges[from], to)
}

func GraphFromGrid(g Grid) Graph {
	nodes := g.Width * g.Height
	graph := NewGraph(nodes)

	idx := func(x, y int) int {
		return y*g.Width + x
	}

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Cells[y][x].North == false {
				graph.AddEdge(idx(x, y), idx(x, y-1))
			}
			if g.Cells[y][x].South == false {
				graph.AddEdge(idx(x, y), idx(x, y+1))
			}
			if g.Cells[y][x].West == false {
				graph.AddEdge(idx(x, y), idx(x-1, y))
			}
			if g.Cells[y][x].East == false {
				graph.AddEdge(idx(x, y), idx(x+1, y))
			}
		}
	}

	return graph
}
