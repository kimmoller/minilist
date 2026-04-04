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

			cmd.Printf("%-4s %-20s %s\n", "ID", "STATUS", "DESCRIPTION")
			cmd.Println(strings.Repeat("-", 80))

			items := sortItems(data.Items)

			for _, item := range items {
				if item.Status == cli.StatusCompleted && !withCompleted {
					continue
				}
				text := fmt.Sprintf("%-4d %-20s %s", item.ID, item.Status, item.Description)
				if item.Priority {
					cmd.Printf("%s\n", "\033[1m"+text+"\033[0m")
				} else {
					cmd.Printf("%s\n", text)
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&withCompleted, withCompletedFlag, "a", false, "(optional) Print completed items")

	return cmd
}

// Sort items into a priority order: PRIORITY > IN_PROGRESS > TODO > COMPLETED
func sortItems(items []cli.Item) []cli.Item {
	itemsCopy := slices.Clone(items)

	slices.SortFunc(itemsCopy, func(a cli.Item, b cli.Item) int {
		// Sort everything before completed items
		if a.Status == cli.StatusCompleted && b.Status != cli.StatusCompleted {
			return 1
		}
		if b.Status == cli.StatusCompleted && a.Status != cli.StatusCompleted {
			return -1
		}

		// Sort prioritized items over normal items
		if a.Priority && !b.Priority {
			return -1
		}
		if b.Priority && !a.Priority {
			return 1
		}

		// Sort in_progress before todo items
		if a.Status == cli.StatusInProgress && b.Status == cli.StatusTodo {
			return -1
		}
		if b.Status == cli.StatusInProgress && a.Status == cli.StatusTodo {
			return 1
		}

		return 0
	})

	return itemsCopy
}
