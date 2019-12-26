package db

import (
	"github.com/spf13/cobra"
)

// NewDBCmd generates the configuration for the query subcommand.
// It can be attached to any upstream cobra command.
func NewDBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "this is the short description for the db subcommand",
		Long:  `this is the long description for the db subcommand`,
	}
	cmd.AddCommand(
		NewCreateCmd(),
	)
	flags := cmd.PersistentFlags()
	flags.String("db", "", "the name of the db")

	return cmd
}
