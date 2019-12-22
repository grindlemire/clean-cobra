package opts

import (
	queryOpts "github.com/grindlemire/clean-cobra/cmd/query/opts"
)

// Options are options for the configured dev subcommand
type Options struct {
	queryOpts.Options
	Username string
}
