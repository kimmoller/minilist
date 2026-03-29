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
			Status:      true,
			Description: "Completed test todo item",
		},
		{
			ID:          1,
			Status:      false,
			Description: "Second test todo item",
		},
		{
			ID:          2,
			Status:      false,
			Description: "Third test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	stdOut, _ := utils.ExecuteCommand(fs, fmt.Sprint("list"))

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
1    IN PROGRESS          Second test todo item
2    IN PROGRESS          Third test todo item
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
			Status:      false,
			Description: "First test todo item",
		},
		{
			ID:          1,
			Status:      true,
			Description: "Second test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	stdOut, _ := utils.ExecuteCommand(fs, fmt.Sprint("list --all"))

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          First test todo item
1    COMPLETED            Second test todo item
	`

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
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
	`

	utils.AssertOutput(t, stdOut, expected)
}
