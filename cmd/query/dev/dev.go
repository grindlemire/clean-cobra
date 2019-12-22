package dev

import (
	"github.com/grindlemire/clean-cobra/cmd/query/dev/opts"
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// NewCmd generates the configuration for the query dev subcommand.
// It can be attached to any upstream cobra command.
func NewCmd() *cobra.Command {
	opts := opts.Options{}

	cmd := &cobra.Command{
		Use:   "dev",
		Short: "query dev subcommand",
		Long:  `this is the long description for the query dev subcommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *opts.Options) {
	flags := cmd.Flags()
	flags.StringVar(&opts.Username, "user", "", "the username to execute the query as")
}

func run(opts opts.Options) error {
	log.Infof("running query dev command with %+v options", opts)
	return nil
}
