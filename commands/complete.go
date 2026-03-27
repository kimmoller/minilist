package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCompleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "complete",
		Short: "complete a todo item",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Complete a todo list item")
		},
	}
}
