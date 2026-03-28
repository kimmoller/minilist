package commands

import (
	"github.com/kimmoller/minilist/cli"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewAddCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use:     "add",
		Short:   "add a new todo item",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"a", "new"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.AddItem(fs, args[0])
		},
	}
}
