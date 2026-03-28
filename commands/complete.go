package commands

import (
	"strconv"

	"github.com/kimmoller/minilist/cli"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewCompleteCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use:     "complete",
		Short:   "Complete a todo item",
		Long:    "Mark a todo item as completed by running the command with the item ID as the only argument.",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"c", "done"},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return cli.CompleteItem(fs, id)
		},
	}
}
