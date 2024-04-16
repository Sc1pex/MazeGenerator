package cmds

import (
	"mazes/internal/interact"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mazes",
	Short: "",
	Long:  "",

	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	interact.StartInteractive()
}

func Execute() error {
	return rootCmd.Execute()
}
