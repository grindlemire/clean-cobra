package opts

import (
	rootOpts "github.com/grindlemire/clean-cobra/cmd/opts"
)

// Options are the options for the query subcommand
type Options struct {
	rootOpts.Options
	Query string
}
