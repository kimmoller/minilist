package cli_test

import (
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

	err = utils.PopulateTestData(fs, filePath, []cli.Item{})
	if err != nil {
		t.Fatal(err)
	}

	err = cli.AddItem(fs, "Test todo item")
	if err != nil {
		t.Fatal(err)
	}

	data, err := utils.DataFromFile(fs, filePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(data.Items))
	item := data.Items[0]
	assert.Equal(t, 0, item.ID)
	assert.Equal(t, false, item.Status)
	assert.Equal(t, "Test todo item", item.Description)
}

func TestCompleteItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	item := cli.Item{
		ID:          0,
		Status:      false,
		Description: "Test todo item",
	}
	err = utils.PopulateTestData(fs, filePath, []cli.Item{item})
	if err != nil {
		t.Fatal(err)
	}

	err = cli.CompleteItem(fs, 0)
	if err != nil {
		t.Fatal(err)
	}

	data, err := utils.DataFromFile(fs, filePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(data.Items))
	modifiedItem := data.Items[0]
	assert.Equal(t, 0, modifiedItem.ID)
	assert.Equal(t, true, modifiedItem.Status)
	assert.Equal(t, "Test todo item", modifiedItem.Description)
}

func TestDeleteItem(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	item := cli.Item{
		ID:          0,
		Status:      false,
		Description: "Test todo item",
	}
	err = utils.PopulateTestData(fs, filePath, []cli.Item{item})
	if err != nil {
		t.Fatal(err)
	}

	data, err := utils.DataFromFile(fs, filePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(data.Items))

	err = cli.DeleteItem(fs, 0)
	if err != nil {
		t.Fatal(err)
	}

	data, err = utils.DataFromFile(fs, filePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, len(data.Items))
}

func TestReadData(t *testing.T) {
	fs := afero.NewMemMapFs()

	filePath, err := cli.DataFilePath()
	if err != nil {
		t.Fatal(err)
	}

	item := cli.Item{
		ID:          0,
		Status:      false,
		Description: "Test todo item",
	}
	err = utils.PopulateTestData(fs, filePath, []cli.Item{item})
	if err != nil {
		t.Fatal(err)
	}

	data, err := cli.ReadData(fs)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(data.Items))
	item = data.Items[0]
	assert.Equal(t, 0, item.ID)
	assert.Equal(t, false, item.Status)
	assert.Equal(t, "Test todo item", item.Description)
}
