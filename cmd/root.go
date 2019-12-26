package cmd

import (
	"github.com/grindlemire/clean-cobra/cmd/query"
	"github.com/grindlemire/clean-cobra/cmd/serve"
	"github.com/spf13/cobra"
)

// Run starts configures and executes the entire cli interface
func Run() error {
	cmd := &cobra.Command{
		Use:           "app",
		Short:         "root command",
		Long:          `this is the long description of the root command`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(
		query.CreateQueryCmd(),
		serve.CreateServeCmd(),
	)
	return cmd.Execute()
}
