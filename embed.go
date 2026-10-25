// Package ridgelight holds the files that are compiled into the Ridgelight binary.
package ridgelight

import "embed"

// Static holds everything in the static/ folder, compiled into the binary.
//
//go:embed static
var Static embed.FS
