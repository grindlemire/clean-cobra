package opts

import (
	rootOpts "github.com/grindlemire/clean-cobra/cmd/opts"
)

// Options for the serve cli
type Options struct {
	rootOpts.Options

	Port int
}
