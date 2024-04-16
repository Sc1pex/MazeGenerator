package main

import "mazes/cmd/cli/cmds"

func main() {
	err := cmds.Execute()
	if err != nil {
		panic(err)
	}
}
