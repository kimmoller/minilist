package commands_test

import (
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/utils"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestStartItem(t *testing.T) {
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

	utils.ExecuteCommand(fs, "start 1")

	stdOut, _ := utils.ExecuteCommand(fs, "list --all")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
1    IN PROGRESS          Second test todo item
0    TODO                 First test todo item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestStartInProgressItem(t *testing.T) {
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
			Status:      cli.StatusInProgress,
			Description: "Second test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	utils.ExecuteCommand(fs, "start 1")

	stdOut, _ := utils.ExecuteCommand(fs, "list --all")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          First test todo item
1    IN PROGRESS          Second test todo item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestStartNonExistingItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "start 0")
	assert.Equal(t, "Error: item with ID 0 not found\n", errOut.String())
}

func TestStartItemWithoutArgs(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "start")
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", errOut.String())
}
