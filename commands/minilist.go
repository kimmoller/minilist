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

	return cmd
}
