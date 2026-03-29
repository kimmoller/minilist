package main

import (
	"os"

	"github.com/kimmoller/minilist/cli"
	"github.com/kimmoller/minilist/commands"
	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()

	cli.EnsureDataFileExists(fs)

	cmd := commands.NewCmd(fs)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Execute()
}
