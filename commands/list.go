package commands

import (
	"fmt"
	"strings"

	"github.com/kimmoller/minilist/data"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	var withCompleted bool

	const withCompletedFlag = "all"

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all todo items",
		Args:    cobra.ExactArgs(0),
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := data.ReadData()
			if err != nil {
				return err
			}

			fmt.Printf("%-4s %-20s %s\n", "ID", "STATUS", "DESCRIPTION")
			fmt.Println(strings.Repeat("-", 80))

			for _, item := range data.Items {
				if item.Status && !withCompleted {
					continue
				}
				statusText := toStatusText(item.Status)
				fmt.Printf("%-4d %-20s %s\n", item.ID, statusText, item.Description)
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&withCompleted, withCompletedFlag, "a", false, "(optional) Print completed items")

	return cmd
}

func toStatusText(status bool) string {
	if status {
		return "COMPLETED"
	}

	return "IN PROGRESS"
}
