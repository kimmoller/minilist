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
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Add a new todo item")
			data.AddItem("Test item")
		},
	}
}
