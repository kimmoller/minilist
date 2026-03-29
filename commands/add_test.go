package commands_test

import (
	"fmt"
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/utils"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestAddItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	// TODO: Change this back to spaces once the test arg splitter has been fixed
	utils.ExecuteCommand(fs, fmt.Sprintf("add %s", "Test_todo_item"))

	stdOut, _ := utils.ExecuteCommand(fs, "list")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          Test_todo_item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestAddItemWithGapInIds(t *testing.T) {
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
			ID:          2,
			Status:      false,
			Description: "Second test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	// TODO: Change this back to spaces once the test arg splitter has been fixed
	utils.ExecuteCommand(fs, fmt.Sprintf("add %s", "Test_todo_item"))

	stdOut, _ := utils.ExecuteCommand(fs, "list")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          First test todo item
2    IN PROGRESS          Second test todo item
3    IN PROGRESS          Test_todo_item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestAddMultipleItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	// TODO: Change this back to spaces once the test arg splitter has been fixed
	utils.ExecuteCommand(fs, fmt.Sprintf("add %s", "Test_todo_item"))
	utils.ExecuteCommand(fs, fmt.Sprintf("add %s", "Test_todo_item_2"))
	utils.ExecuteCommand(fs, fmt.Sprintf("add %s", "Test_todo_item_3"))

	stdOut, _ := utils.ExecuteCommand(fs, "list")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          Test_todo_item
1    IN PROGRESS          Test_todo_item_2
2    IN PROGRESS          Test_todo_item_3
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestAddItemWithoutDescription(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "add")
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", errOut.String())
}
