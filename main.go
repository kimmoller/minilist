package main

import (
	"os"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/commands"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// TODO: Add tests for commands

func NewCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "minilist",
		Short: "Minimalistic todo list",
		Long:  "A simple todo list with basic commands",
	}

	cmd.AddCommand(commands.NewAddCmd(fs))
	cmd.AddCommand(commands.NewDeleteCmd(fs))
	cmd.AddCommand(commands.NewListCmd(fs))
	cmd.AddCommand(commands.NewCompleteCmd(fs))

	return cmd
}

func main() {
	fs := afero.NewOsFs()

	cli.EnsureDataFileExists(fs)

	cmd := NewCmd(fs)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Execute()
}
