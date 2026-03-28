package utils

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/commands"
	"github.com/spf13/afero"
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

func ExecuteCommand(fs afero.Fs, command string) (*bytes.Buffer, *bytes.Buffer) {
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
