package commands

import (
	"github.com/kimmoller/minilist/data"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "add",
		Short:   "add a new todo item",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"a", "new"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return data.AddItem(args[0])
		},
	}
}
