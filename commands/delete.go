package commands

import (
	"fmt"
	"strconv"

	"github.com/kimmoller/minilist/data"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "delete a todo item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing ID")
			}

			if len(args) > 1 {
				return fmt.Errorf("too many arguments")
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return data.DeleteItem(id)
		},
	}
}
