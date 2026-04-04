package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/commands"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func DataFromFile(fs afero.Fs, filePath string) (*cli.Data, error) {
	byteData, err := afero.ReadFile(fs, filePath)
	if err != nil {
		return nil, err
	}

	var data cli.Data
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func PopulateTestData(fs afero.Fs, filePath string, items []cli.Item) error {
	testData := cli.Data{
		Items: items,
	}

	byteData, err := json.Marshal(testData)
	if err != nil {
		return err
	}

	return afero.WriteFile(fs, filePath, byteData, 0644)
}

// TODO_MIGRATION: Remove in a future version
func PopulateTestDataForMigration(fs afero.Fs, filePath string, items []cli.OldItem) error {
	testData := cli.OldData{
		Items: items,
	}

	byteData, err := json.Marshal(testData)
	if err != nil {
		return err
	}

	return afero.WriteFile(fs, filePath, byteData, 0644)
}

func ExecuteCommand(fs afero.Fs, command string) (*bytes.Buffer, *bytes.Buffer) {
	// FIXME: Currently the split does not correctly handle agruments with spaces, like the description
	args := strings.Split(command, " ")
	cmd := commands.NewCmd(fs)

	stdOut := bytes.NewBufferString("")
	errOut := bytes.NewBufferString("")

	cmd.SetArgs(args)

	cmd.SetOut(stdOut)
	cmd.SetErr(errOut)

	cmd.Execute()

	return stdOut, errOut
}

func AssertOutput(t *testing.T, stdOut *bytes.Buffer, expected string) {
	out, err := io.ReadAll(stdOut)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(string(out)))
}
