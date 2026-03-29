package commands_test

import (
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/utils"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestDeleteItem(t *testing.T) {
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
			Status:      false,
			Description: "Second test todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	utils.ExecuteCommand(fs, "delete 1")

	stdOut, _ := utils.ExecuteCommand(fs, "list --all")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          First test todo item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestDeleteNonExistingItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "delete 0")
	assert.Equal(t, "Error: item with ID 0 not found\n", errOut.String())
}

func TestDeleteItemWithoutArgs(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "delete")
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", errOut.String())
}
