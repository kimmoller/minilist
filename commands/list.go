package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kimmoller/minilist/cli"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewListCmd(fs afero.Fs) *cobra.Command {
	var withCompleted bool

	const withCompletedFlag = "all"

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all todo items",
		Args:    cobra.ExactArgs(0),
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := cli.ReadData(fs)
			if err != nil {
				return err
			}

			statusMap := toStatusMap(data.Items)

			cmd.Printf("%s\n", strings.Repeat("-", 33)+" IN PROGRESS "+strings.Repeat("-", 34))
			cmd.Println()

			for _, item := range statusMap[cli.StatusInProgress] {
				printItem(cmd, item)
			}

			cmd.Println()
			cmd.Printf("%s\n", strings.Repeat("-", 37)+" TODO "+strings.Repeat("-", 37))
			cmd.Println()

			for _, item := range statusMap[cli.StatusTodo] {
				printItem(cmd, item)
			}

			cmd.Println()

			if withCompleted {

				cmd.Printf("%s\n", strings.Repeat("-", 34)+" COMPLETED "+strings.Repeat("-", 35))
				cmd.Println()

				for _, item := range statusMap[cli.StatusCompleted] {
					printItem(cmd, item)
				}
			}

			return nil
		},
	}
	cmd.Flags().BoolVarP(&withCompleted, withCompletedFlag, "a", false, "(optional) Print completed items")

	return cmd
}

func toStatusMap(items []cli.Item) map[cli.Status][]cli.Item {
	todoList := []cli.Item{}
	inProgressList := []cli.Item{}
	completedList := []cli.Item{}

	for _, item := range items {
		switch item.Status {
		case cli.StatusTodo:
			todoList = append(todoList, item)
		case cli.StatusInProgress:
			inProgressList = append(inProgressList, item)
		case cli.StatusCompleted:
			completedList = append(completedList, item)
		}
	}

	todoList = sortItems(todoList)
	inProgressList = sortItems(inProgressList)

	return map[cli.Status][]cli.Item{
		cli.StatusTodo:       todoList,
		cli.StatusInProgress: inProgressList,
		cli.StatusCompleted:  completedList,
	}
}

func printItem(cmd *cobra.Command, item cli.Item) {
	text := fmt.Sprintf("%-4d %s", item.ID, item.Description)
	if item.Priority {
		cmd.Printf("%s\n", "\033[1m"+text+"\033[0m")
	} else {
		cmd.Printf("%s\n", text)
	}
}

// Sort items by priority
func sortItems(items []cli.Item) []cli.Item {
	itemsCopy := slices.Clone(items)

	slices.SortFunc(itemsCopy, func(a cli.Item, b cli.Item) int {
		// Sort prioritized items over normal items
		if a.Priority && !b.Priority {
			return -1
		}
		if b.Priority && !a.Priority {
			return 1
		}

		return 0
	})

	return itemsCopy
}
