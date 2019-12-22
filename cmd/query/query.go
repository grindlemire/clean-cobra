package query

import (
	rootOpts "github.com/grindlemire/clean-cobra/cmd/opts"
	"github.com/grindlemire/clean-cobra/cmd/query/dev"
	"github.com/grindlemire/clean-cobra/cmd/query/opts"
	"github.com/grindlemire/clean-cobra/cmd/query/prod"
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// CreateQueryCmd generates the configuration for the query subcommand.
// It can be attached to any upstream cobra command.
func CreateQueryCmd(globalOpts rootOpts.Options) *cobra.Command {
	options := opts.Options{
		Options: globalOpts,
	}

	cmd := &cobra.Command{
		Use:   "query",
		Short: "this is the short description for the query subcommand",
		Long:  `this is the long description for the query subcommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(options)
		},
	}
	attachOpts(cmd, &options)
	cmd.AddCommand(
		prod.NewCmd(options),
		dev.NewCmd(),
	)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *opts.Options) {
	flags := cmd.PersistentFlags()
	flags.StringVarP(&opts.Query, "query", "q", "", "the query to execute")
	cmd.MarkPersistentFlagRequired("query")
}

func run(opts opts.Options) error {
	log.Infof("running query command with %+v options", opts)
	return nil
}
