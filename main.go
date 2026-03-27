package main

import (
	"os"

	"github.com/kimmoller/minilist/commands"
	"github.com/kimmoller/minilist/data"
	"github.com/spf13/cobra"
)

// TODO: Add tests for commands and data functions

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "minilist",
		Short: "Minimalistic todo list",
		Long:  "A simple todo list with basic commands",
	}

	cmd.AddCommand(commands.NewAddCmd())
	cmd.AddCommand(commands.NewDeleteCmd())
	cmd.AddCommand(commands.NewListCmd())
	cmd.AddCommand(commands.NewCompleteCmd())

	return cmd
}

func main() {
	data.EnsureDataFileExists()
	cmd := NewCmd()
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Execute()
}
