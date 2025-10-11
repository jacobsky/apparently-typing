// Package static contains the reference to all the static HTML that will be included in the binary to be served.
// This package is currently only used for blogging.
package static

import "embed"

//go:embed blog/*.md
var BlogFiles embed.FS
