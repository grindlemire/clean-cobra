package query

import (
	"github.com/grindlemire/clean-cobra/cmd/query/db"
	"github.com/spf13/cobra"
)

// CreateQueryCmd generates the configuration for the query subcommand.
// It can be attached to any upstream cobra command.
func CreateQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "this is the short description for the query subcommand",
		Long:  `this is the long description for the query subcommand`,
	}
	cmd.AddCommand(
		NewProdCmd(),
		NewDevCmd(),
		db.NewDBCmd(),
	)
	flags := cmd.PersistentFlags()
	flags.StringP("query", "q", "", "the query to execute")

	return cmd
}
