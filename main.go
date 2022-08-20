package main

import (
	"github.com/sjansen/tribble/internal/cli"
)

var build string // set by goreleaser

func main() {
	if build == "" {
		build = version
	}
	cli.ParseAndRun(build)
}
