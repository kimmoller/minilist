package commands

import (
	"strconv"

	"github.com/kimmoller/minilist/cli"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use:     "delete",
		Short:   "Delete a todo item",
		Long:    "Delete a todo item by giving the ID as the only argument.",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"d", "remove", "rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return cli.DeleteItem(fs, id)
		},
	}
}
