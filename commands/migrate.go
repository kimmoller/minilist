package commands

import (
	"github.com/kimmoller/minilist/cli"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// TODO_MIGRATION: Remove in a future version
func NewMigrateCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate to the new data model",
		Long:  "Migrate from the v0.1.1 data model to the new model introduced in v.0.2.0",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cli.Migrate(fs)
			if err != nil {
				if err.Error() == "Data already in the new format, nothing to migrate" {
					cmd.Println(err)
					return nil
				}
				return err
			}

			cmd.Println("Migration complete")
			return nil
		},
	}
}
