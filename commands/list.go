package commands

import (
	"fmt"

	"github.com/kimmoller/minilist/data"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all todo items",
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := data.ReadData()
			if err != nil {
				return err
			}

			for _, item := range data.Items {
				cmd.Println(fmt.Printf("%d: Status: %t, %s", item.ID, item.Status, item.Description))
			}
			return nil
		},
	}
}
