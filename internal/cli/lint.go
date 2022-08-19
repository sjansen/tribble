package cli

import (
	"context"
	"fmt"

	"github.com/sjansen/tribble/internal/project/template"
)

type lintCmd struct {
	Template string `kong:"arg,default='.',help='Path or URL of project template.'"`
	Branch   string `kong:"arg,optional,help='TODO'"`
	Strict   bool   `kong:"optional,help='Report warnings as errors.'"`
}

func (cmd *lintCmd) Run() error {
	ctx := context.TODO()
	t, err := template.Open(ctx, cmd.Template, cmd.Branch)
	if err != nil {
		return err
	}

	// TODO lint Config

	questions, err := t.LoadQuestions()
	if err != nil {
		return err
	}
	for group, questions := range questions {
		for name, q := range questions {
			errors, warnings := q.Lint()
			if len(errors) > 0 || len(warnings) > 0 {
				fmt.Print(group, ".", name, "\n")
				for _, e := range errors {
					fmt.Println("  ERROR:", e)
				}
				for _, w := range warnings {
					fmt.Println("  WARNING:", w)
				}
			}
		}
	}

	// TODO exit with error code if errors
	// TODO optionally exit with errors is warnings
	return nil
}
