package interact

import (
	"errors"
	"fmt"
	"mazes/internal/algs"
	"mazes/internal/data"
	"slices"

	"github.com/charmbracelet/lipgloss"
)

type generateCmd struct{}

func (cmd generateCmd) Execute(args []string) error {
	g := data.NewGrid(S.width, S.height)
	S.grid = &g

	switch S.alg {
	case "binaryTree":
		algs.BinTree(S.grid)
	case "sidewinder":
		algs.Sidewinder(S.grid)
	case "aldous-broder":
		algs.AldousBroder(S.grid)
	default:
		return errors.New("Unknown algorithm (unreachable?)")
	}

	fmt.Println("Generated grid")
	fmt.Println(S.grid.String())
	return nil
}

func (cmd generateCmd) Names() []string {
	return []string{"generate", "gen"}
}

func (cmd generateCmd) Info() string {
	return "generate | gen: Generate a new maze"
}

type solveCmd struct{}

func (cmd solveCmd) Execute(args []string) error {
	if S.grid == nil {
		return errors.New("No grid to solve")
	}

	g := data.GraphFromGrid(*S.grid)
	S.graph = &g

	bfs := data.BFS(S.graph, 0)
	path := data.PathTo(0, S.graph.Nodes-1, bfs)

	fmt.Println(S.grid.StringNumbers(func(i int) string {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("#aaaaaa"))
		if slices.Contains(path, i) {
			style = lipgloss.NewStyle().Foreground(lipgloss.Color("#40a02b"))
		}
		return style.Render(fmt.Sprintf("%2d ", bfs.Distances[i]))
	}))

	return nil
}

func (cmd solveCmd) Names() []string {
	return []string{"solve"}
}

func (cmd solveCmd) Info() string {
	return "solve: Solve the maze"
}

type printCmd struct{}

func (cmd printCmd) Execute(args []string) error {
	if S.grid == nil {
		return errors.New("No grid to print")
	}

	fmt.Println(S.grid.String())
	return nil
}

func (cmd printCmd) Names() []string {
	return []string{"print"}
}

func (cmd printCmd) Info() string {
	return "print: Print the maze"
}

type longestPathCmd struct{}

func (cmd longestPathCmd) Execute(args []string) error {
	if S.grid == nil {
		return errors.New("No grid to find longest path")
	}

	g := data.GraphFromGrid(*S.grid)
	S.graph = &g
	r := data.Dfs(*S.graph)
	path := data.Path(r)

	fmt.Println(S.grid.StringNumbers(func(i int) string {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("#aaaaaa"))
		if slices.Contains(path, i) {
			style = lipgloss.NewStyle().Foreground(lipgloss.Color("#40a02b"))
		}
		return style.Render(fmt.Sprintf("%2d ", r.Depth[i]))
	}))

	return nil
}

func (cmd longestPathCmd) Names() []string {
	return []string{"solvelong", "sl"}
}

func (cmd longestPathCmd) Info() string {
	return "solvelong | sl: Find the longest path in the maze"
}
