package query

import (
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// Options are options for the configured dev subcommand
type devOptions struct {
	Username string
}

// NewDevCmd generates the configuration for the query dev subcommand.
// It can be attached to any upstream cobra command.
func NewDevCmd() *cobra.Command {
	opts := devOptions{}

	cmd := &cobra.Command{
		Use:   "dev",
		Short: "query dev subcommand",
		Long:  `this is the long description for the query dev subcommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			query, err := cmd.Flags().GetString("query")
			if err != nil {
				return err
			}

			return runDev(opts, query)
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, options *devOptions) {
	flags := cmd.Flags()
	flags.StringVar(&options.Username, "user", "", "the username to execute the query as")
}

func runDev(options devOptions, query string) error {
	log.Infof("running query dev command with %+v options. query: [%s]", options, query)
	return nil
}
