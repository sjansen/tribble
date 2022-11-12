package cli

import (
	"os"

	"github.com/alecthomas/kong"
)

var cli struct {
	Create  createCmd  `kong:"cmd,help='Create a new project from a project template'"`
	Help    helpCmd    `kong:"cmd,help='Show this help text'"`
	Init    initCmd    `kong:"cmd,help='Convert an existing directory into a project template'"`
	Lint    lintCmd    `kong:"cmd,help='Check a project template for mistakes'"`
	Update  updateCmd  `kong:"cmd,help='Apply templates changes to an existing project'"`
	Version versionCmd `kong:"cmd,help='Print Tribble’s version number'"`
}

// ParseAndRun parses command line arguments, then runs the matching command.
func ParseAndRun(version string) {
	ctx := parse(os.Args[1:])
	cli.Help.ctx = ctx
	cli.Version.version = version
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}

func parse(args []string) *kong.Context {
	parser, err := kong.New(&cli,
		kong.UsageOnError(),
	)
	if err != nil {
		panic(err)
	}

	ctx, err := parser.Parse(args)
	parser.FatalIfErrorf(err)

	return ctx
}
