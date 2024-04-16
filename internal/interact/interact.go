package interact

import (
	"bufio"
	"errors"
	"fmt"
	"mazes/internal/data"
	"os"
	"strings"
)

type Status struct {
	width  int
	height int
	alg    string
	grid   *data.Grid
	graph  *data.Graph
}

var S Status = Status{
	width:  5,
	height: 5,
	alg:    "binaryTree",
}

func (s *Status) SetWidth(width int) {
	s.width = width
}

func (s *Status) SetHeight(height int) {
	s.height = height
}

func (s *Status) SetAlg(alg string) error {
	switch alg {
	case "binaryTree":
		s.alg = alg
	case "sidewinder":
		s.alg = alg
	default:
		return errors.New("Unknown algorithm")
	}
	return nil
}

func (s *Status) SetGrid(g data.Grid) {
	s.grid = &g
	s.width = g.Width
	s.height = g.Height
}

func StartInteractive() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for reader.Scan() {
		text := reader.Text()
		if text == "exit" {
			break
		}
		Execute(strings.Split(text, " "))

		fmt.Print("> ")
	}
}
