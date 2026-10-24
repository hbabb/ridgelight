package ridgelight

import "embed"

// Static holds everything in the static/ folder, compiled into the binary.
//
//go:embed static
var Static embed.FS
