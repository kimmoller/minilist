package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "delete a todo item",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Delete a todo list item")
		},
	}
}
