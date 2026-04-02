package commands_test

import (
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/utils"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCompleteItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	items := []cli.Item{
		{
			ID:          0,
			Status:      cli.StatusTodo,
			Description: "First test todo item",
		},
		{
			ID:          1,
			Status:      cli.StatusTodo,
			Description: "Second test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	utils.ExecuteCommand(fs, "complete 1")

	stdOut, _ := utils.ExecuteCommand(fs, "list --all")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    TODO                 First test todo item
1    COMPLETED            Second test todo item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestCompleteCompletedItem(t *testing.T) {
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

	utils.ExecuteCommand(fs, "complete 1")

	stdOut, _ := utils.ExecuteCommand(fs, "list --all")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          First test todo item
1    COMPLETED            Second test todo item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestCompleteNonExistingItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "complete 0")
	assert.Equal(t, "Error: item with ID 0 not found\n", errOut.String())
}

func TestCompleteItemWithoutArgs(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "complete")
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", errOut.String())
}
