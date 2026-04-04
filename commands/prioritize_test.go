package commands_test

import (
	"fmt"
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/utils"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestPrioritizeItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	items := []cli.Item{
		{
			ID:          0,
			Status:      cli.StatusInProgress,
			Description: "Normal in progress item",
		},
		{
			ID:          1,
			Status:      cli.StatusTodo,
			Description: "Prioritized todo item",
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	utils.ExecuteCommand(fs, "prioritize 1")

	stdOut, _ := utils.ExecuteCommand(fs, "list")

	toBold := "1    TODO                 Prioritized todo item"
	boldText := fmt.Sprintf("%s", "\033[1m"+toBold+"\033[0m")

	expected := fmt.Sprintf(`
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
%s
0    IN PROGRESS          Normal in progress item
		`, boldText)

	utils.AssertOutput(t, stdOut, expected)
}

func TestPrioritizePrioritizedItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	items := []cli.Item{
		{
			ID:          0,
			Status:      cli.StatusInProgress,
			Description: "Unprioritized item",
			Priority:    false,
		},
		{
			ID:          1,
			Status:      cli.StatusInProgress,
			Description: "Prioritized item",
			Priority:    true,
		},
	}

	utils.PopulateTestData(fs, filePath, items)

	utils.ExecuteCommand(fs, "prioritize 1")

	stdOut, _ := utils.ExecuteCommand(fs, "list")

	expected := `
	ID   STATUS               DESCRIPTION
--------------------------------------------------------------------------------
0    IN PROGRESS          Unprioritized item
1    IN PROGRESS          Prioritized item
	`

	utils.AssertOutput(t, stdOut, expected)
}

func TestPrioritizeNonExistingItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "prioritize 0")
	assert.Equal(t, "Error: item with ID 0 not found\n", errOut.String())
}

func TestPrioritizeItemWithoutArgs(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	utils.PopulateTestData(fs, filePath, []cli.Item{})

	_, errOut := utils.ExecuteCommand(fs, "prioritize")
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", errOut.String())
}
