package db

import (
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// createOpts are the configuration options for the db subcommand
type createOpts struct{}

// NewCreateCmd generates the configuration for the query db subcommand.
// It can be attached to any upstream cobra command.
func NewCreateCmd() *cobra.Command {
	opts := createOpts{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create db subcommand",
		Long:  `this is the long description for the create db subcommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			query, err := cmd.Flags().GetString("query")
			if err != nil {
				return err
			}

			db, err := cmd.Flags().GetString("db")
			if err != nil {
				return err
			}
			return runDB(opts, query, db)
		},
	}
	return cmd
}

func runDB(options createOpts, query string, db string) error {
	log.Infof("running query prod command with %+v options. Query [%s]. DB: [%s]", options, query, db)
	return nil
}
