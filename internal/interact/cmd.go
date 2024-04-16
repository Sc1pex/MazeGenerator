package interact

import (
	"fmt"
	"os"
	"slices"
)

type Command interface {
	Execute(args []string) error
	Names() []string
	Info() string
}

type commands struct {
	cmds []Command
}

var cmds commands = commands{
	cmds: []Command{
		widthCmd{},
		heightCmd{},
		sizeCmd{},
		algCmd{},
		generateCmd{},
		solveCmd{},
		printCmd{},
		longestPathCmd{},
	},
}

func Execute(args []string) {
	cmd := args[0]
	if cmd == "exit" {
		os.Exit(0)
	}

	for _, c := range cmds.cmds {
		if slices.Contains(c.Names(), cmd) {
			err := c.Execute(args)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}

	fmt.Println("Unknown command. Valid commands are:")
	for _, c := range cmds.cmds {
		fmt.Println(c.Info())
	}
}
