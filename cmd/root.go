package cmd

import (
	"github.com/grindlemire/clean-cobra/cmd/opts"
	"github.com/grindlemire/clean-cobra/cmd/query"
	"github.com/grindlemire/clean-cobra/cmd/serve"
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// Run starts configures and executes the entire cli interface
func Run() error {
	options := opts.Options{
		UnconfiguredOpt: true,
	}

	cmd := &cobra.Command{
		Use:   "app",
		Short: "root command",
		Long:  `this is the long description of the root command`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(options)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	attachOpts(cmd, &options)
	cmd.AddCommand(
		query.CreateQueryCmd(options),
		serve.CreateServeCmd(options),
	)
	return cmd.Execute()
}

func attachOpts(cmd *cobra.Command, opts *opts.Options) {
	flags := cmd.PersistentFlags()
	flags.StringVarP(&opts.Global, "global", "g", "", "global usage string")
}

func run(opts opts.Options) error {
	log.Infof("global options: %+v", opts)
	return nil
}
