package interact

import "errors"

type algCmd struct{}

func (a algCmd) Execute(args []string) error {
	if len(args) != 2 {
		return errors.New("Expected 1 argument")
	}
	err := S.SetAlg(args[1])
	if err != nil {
		return err
	}
	return nil
}

func (a algCmd) Names() []string {
	return []string{"alg"}
}

func (a algCmd) Info() string {
	return "alg <algorithm> - Set the algorithm to use (binaryTree, sidewindwer)"
}
