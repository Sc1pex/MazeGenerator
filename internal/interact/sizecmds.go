package interact

import (
	"errors"
	"strconv"
)

type widthCmd struct{}

func (w widthCmd) Execute(args []string) error {
	if len(args) != 2 {
		return errors.New("Expected 1 argument")
	}
	width, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Invalid width. Please enter a valid number.")
	}
	S.SetWidth(width)
	return nil
}

func (w widthCmd) Names() []string {
	return []string{"width", "w"}
}

func (w widthCmd) Info() string {
	return "width | w <number> - Set the width of the grid"
}

type heightCmd struct{}

func (h heightCmd) Execute(args []string) error {
	if len(args) != 2 {
		return errors.New("Expected 1 argument")
	}
	height, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Invalid height. Please enter a valid number.")
	}
	S.SetHeight(height)
	return nil
}

func (h heightCmd) Names() []string {
	return []string{"height", "h"}
}

func (h heightCmd) Info() string {
	return "height | h <number> - Set the height of the grid"
}

type sizeCmd struct{}

func (s sizeCmd) Execute(args []string) error {
	if len(args) != 3 {
		return errors.New("Expected 2 arguments")
	}
	width, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Invalid width. Please enter a valid number.")
	}
	height, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("Invalid height. Please enter a valid number.")
	}
	S.SetWidth(width)
	S.SetHeight(height)
	return nil
}

func (s sizeCmd) Names() []string {
	return []string{"size", "s"}
}

func (s sizeCmd) Info() string {
	return "size | s <width> <height> - Set the size of the grid"
}
