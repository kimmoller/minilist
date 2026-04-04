package commands

import (
	"strconv"

	"github.com/kimmoller/minilist/cli"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewStartCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use:     "start",
		Short:   "Set an item to in progress",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return cli.SetToInProgress(fs, id)
		},
	}
}
