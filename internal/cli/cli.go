package cli

import (
	"os"

	"github.com/alecthomas/kong"
)

var cli struct {
	Init   initCmd   `kong:"cmd"`
	New    newCmd    `kong:"cmd"`
	Update updateCmd `kong:"cmd"`
}

// ParseAndRun parses command line arguments, then runs the matching command.
func ParseAndRun() {
	ctx := parse(os.Args[1:])
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
