package commands

import (
	"strconv"

	"github.com/kimmoller/minilist/cli"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewPrioritizeCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use:     "prioritize",
		Short:   "Prioritize a todo item",
		Long:    "Toggle the priority of a todo item. If the item is already marked as prioritized, it will be unprioritized.",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"p"},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return cli.TogglePriority(fs, id)
		},
	}
}
