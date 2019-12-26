package query

import (
	"github.com/grindlemire/log"
	"github.com/spf13/cobra"
)

// Options are the configuration options for the prod subcommand
type prodOptions struct {
	Auth string
}

// NewProdCmd generates the configuration for the query prod subcommand.
// It can be attached to any upstream cobra command.
func NewProdCmd() *cobra.Command {
	opts := prodOptions{}

	cmd := &cobra.Command{
		Use:   "prod",
		Short: "query prod subcommand",
		Long:  `this is the long description for the query prod subcommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			query, err := cmd.Flags().GetString("query")
			if err != nil {
				return err
			}
			return runProd(opts, query)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.Auth, "auth", "", "the auth to execute the query as")
	cmd.MarkFlagRequired("auth")

	return cmd
}

func runProd(options prodOptions, query string) error {
	log.Infof("running query prod command with %+v options. Query [%s]", options, query)
	return nil
}
