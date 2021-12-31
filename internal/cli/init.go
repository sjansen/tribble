package cli

import (
	"os"

	"github.com/go-git/go-billy/v5/osfs"

	"github.com/sjansen/tribble/internal/template"
)

type initCmd struct {
	Force bool `kong:"short='f',help='Force replacement of existing files.'"`
}

func (cmd *initCmd) Run() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	return template.Initialize(osfs.New(cwd), cmd.Force)
}
