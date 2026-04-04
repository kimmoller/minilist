package commands

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "minilist",
		Short: "Minimalistic todo list",
		Long:  "A simple todo list with basic commands",
	}

	cmd.AddCommand(NewAddCmd(fs))
	cmd.AddCommand(NewDeleteCmd(fs))
	cmd.AddCommand(NewListCmd(fs))
	cmd.AddCommand(NewCompleteCmd(fs))
	cmd.AddCommand(NewStartCmd(fs))
	cmd.AddCommand(NewPrioritizeCmd(fs))
	// TODO_MIGRATION: Remove in a future version
	cmd.AddCommand(NewMigrateCmd(fs))

	return cmd
}
