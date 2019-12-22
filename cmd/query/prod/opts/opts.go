package opts

import (
	queryOpts "github.com/grindlemire/clean-cobra/cmd/query/opts"
)

// Options are the configuration options for the prod subcommand
type Options struct {
	queryOpts.Options
	Auth string
}
