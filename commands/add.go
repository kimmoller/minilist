package commands

import (
	"fmt"

	"github.com/kimmoller/minilist/data"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "add a new todo item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing description")
			}

			if len(args) > 1 {
				return fmt.Errorf("too many arguments")
			}

			return data.AddItem(args[0])
		},
	}
}
