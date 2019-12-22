package serve

import (
	rootOpts "github.com/grindlemire/clean-cobra/cmd/opts"
	"github.com/grindlemire/clean-cobra/cmd/serve/opts"
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// CreateServeCmd generates the configuration for the serve subcommand.
// It can be attached to any upstream cobra command.
func CreateServeCmd(globalOpts rootOpts.Options) *cobra.Command {
	opts := opts.Options{
		Options: globalOpts,
	}

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "this is the short description for serve",
		Long:  `this is the long description for serve`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, options *opts.Options) {
	flags := cmd.Flags()
	flags.IntVarP(&options.Port, "port", "p", 10000, "the port to serve on")
}

func run(options opts.Options) error {
	log.Infof("running query command with %+v options", options)
	return nil
}
