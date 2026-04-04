package commands_test

import (
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/utils"
	"github.com/spf13/afero"
)

// TODO_MIGRATION: Remove in a future version
func TestMigrateWithOldData(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	oldItems := []cli.OldItem{
		cli.OldItem{
			ID:          0,
			Status:      false,
			Description: "Old item with false status",
		},
		cli.OldItem{
			ID:          1,
			Status:      true,
			Description: "Old item with true status",
		},
	}

	err = utils.PopulateTestDataForMigration(fs, filePath, oldItems)
	if err != nil {
		t.Fatal(err)
	}

	stdOut, _ := utils.ExecuteCommand(fs, "migrate")
	expected := "Migration complete"
	utils.AssertOutput(t, stdOut, expected)
}

func TestMigrateWithMigratedData(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	items := []cli.Item{
		cli.Item{
			ID:          0,
			Status:      cli.StatusInProgress,
			Description: "Old item with in prgress status",
		},
		cli.Item{
			ID:          1,
			Status:      cli.StatusCompleted,
			Description: "Old item with completed status",
		},
	}

	err = utils.PopulateTestData(fs, filePath, items)
	if err != nil {
		t.Fatal(err)
	}

	stdOut, _ := utils.ExecuteCommand(fs, "migrate")
	expected := "Data already in the new format, nothing to migrate"
	utils.AssertOutput(t, stdOut, expected)
}
