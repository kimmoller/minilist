package commands_test

import (
	"fmt"
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/utils"
	"github.com/spf13/afero"
)

func TestListItems(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	items := []cli.Item{
		{
			ID:          0,
			Status:      cli.StatusCompleted,
			Description: "Completed test todo item",
		},
		{
			ID:          1,
			Status:      cli.StatusTodo,
			Description: "Second test todo item",
		},
		{
			ID:          2,
			Status:      cli.StatusInProgress,
			Description: "Third test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	stdOut, _ := utils.ExecuteCommand(fs, fmt.Sprint("list"))

	expected := `
--------------------------------- IN PROGRESS ----------------------------------

2    Third test todo item

------------------------------------- TODO -------------------------------------

1    Second test todo item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestListAllItems(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	items := []cli.Item{
		{
			ID:          0,
			Status:      cli.StatusInProgress,
			Description: "First test todo item",
		},
		{
			ID:          1,
			Status:      cli.StatusCompleted,
			Description: "Second test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	stdOut, _ := utils.ExecuteCommand(fs, fmt.Sprint("list --all"))

	expected := `
--------------------------------- IN PROGRESS ----------------------------------

0    First test todo item

------------------------------------- TODO -------------------------------------


---------------------------------- COMPLETED -----------------------------------

1    Second test todo item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestListItemsInCorrectOrder(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	items := []cli.Item{
		{
			ID:          0,
			Status:      cli.StatusInProgress,
			Description: "Only one in progress",
		},
		{
			ID:          1,
			Status:      cli.StatusCompleted,
			Description: "First in completed",
		},
		{
			ID:          2,
			Status:      cli.StatusTodo,
			Description: "Second in todo",
		},
		{
			ID:          3,
			Status:      cli.StatusTodo,
			Description: "First in todo",
			Priority:    true,
		},
		{
			ID:          4,
			Status:      cli.StatusCompleted,
			Description: "Last in completed",
			Priority:    true,
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	stdOut, _ := utils.ExecuteCommand(fs, fmt.Sprint("list --all"))

	firstDescription := "3    First in todo"
	firstBold := fmt.Sprintf("%s", "\033[1m"+firstDescription+"\033[0m")

	fourthDescription := "4    Last in completed"
	fourthBold := fmt.Sprintf("%s", "\033[1m"+fourthDescription+"\033[0m")

	expected := fmt.Sprintf(`
--------------------------------- IN PROGRESS ----------------------------------

0    Only one in progress

------------------------------------- TODO -------------------------------------

%s
2    Second in todo

---------------------------------- COMPLETED -----------------------------------

1    First in completed
%s
	`, firstBold, fourthBold)

	utils.AssertOutput(t, stdOut, expected)
}

func TestListNoItems(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	stdOut, _ := utils.ExecuteCommand(fs, fmt.Sprint("list"))

	expected := `
--------------------------------- IN PROGRESS ----------------------------------


------------------------------------- TODO -------------------------------------
	`

	utils.AssertOutput(t, stdOut, expected)
}
