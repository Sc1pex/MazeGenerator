package cmds

import (
	"errors"
	"fmt"
	"mazes/internal/algs"
	"mazes/internal/data"
	"mazes/internal/interact"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("width", "w", 5, "Width of the grid")
	generateCmd.Flags().IntP("height", "h", 5, "Height of the grid")
	generateCmd.Flags().StringP("algorithm", "a", "", `One of: binaryTree, sidewindwer`)
	generateCmd.MarkFlagRequired("width")
	generateCmd.MarkFlagRequired("height")
	generateCmd.MarkFlagRequired("algorithm")

	generateCmd.Flags().BoolP("help", "H", false, "Help for the generate command")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Long:  "",

	RunE: generate,
}

func generate(cmd *cobra.Command, args []string) error {
	width, err := cmd.Flags().GetInt("width")
	if err != nil {
		return err
	}
	height, err := cmd.Flags().GetInt("height")
	if err != nil {
		return err
	}
	algorithm, err := cmd.Flags().GetString("algorithm")
	if err != nil {
		return err
	}

	g := data.NewGrid(width, height)
	switch algorithm {
	case "binaryTree":
		algs.BinTree(&g)
	case "sidewinder":
		algs.Sidewinder(&g)
	default:
		return errors.New("Unknown algorithm")
	}

	fmt.Println(g.String())

	interact.S.SetGrid(g)
	interact.S.SetAlg(algorithm)
	interact.StartInteractive()

	return nil
}
