package prod

import (
	queryOpts "github.com/grindlemire/clean-cobra/cmd/query/opts"
	"github.com/grindlemire/clean-cobra/cmd/query/prod/opts"
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// NewCmd generates the configuration for the query prod subcommand.
// It can be attached to any upstream cobra command.
func NewCmd(globalOpts queryOpts.Options) *cobra.Command {
	opts := opts.Options{
		Options: globalOpts,
	}

	cmd := &cobra.Command{
		Use:   "prod",
		Short: "query prod subcommand",
		Long:  `this is the long description for the query prod subcommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, options *opts.Options) {
	flags := cmd.Flags()
	flags.StringVar(&options.Auth, "auth", "", "the auth to execute the query as")
	cmd.MarkFlagRequired("auth")
}

func run(options opts.Options) error {
	log.Infof("running query prod command with %+v options", options)
	return nil
}
