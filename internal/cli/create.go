package cli

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-billy/v5/osfs"

	"github.com/sjansen/tribble/internal/errors"
	"github.com/sjansen/tribble/internal/project"
	"github.com/sjansen/tribble/internal/project/template"
)

type createCmd struct {
	Project  string `kong:"arg,help='Path to new project.'"`
	Template string `kong:"arg,default='.',help='Path or URL of project template.'"`
	Branch   string `kong:"arg,optional,help='TODO'"`
	Timeout  int    `kong:"default='60',short='t',help='Timeout in seconds'"`
}

func (cmd *createCmd) Run() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cmd.Timeout)*time.Second,
	)
	defer cancel()

	dst := filepath.Clean(cmd.Project)
	if err := mkdir(dst); err != nil {
		return err
	}

	tmpl, err := template.Open(ctx, cmd.Template, "")
	if err != nil {
		return err
	}

	origin, refname, err := tmpl.Origin(dst)
	if err != nil {
		return err
	}

	cfg := &project.Config{}
	cfg.Origin.URL = origin
	cfg.Origin.RefName = refname

	answers := &template.Answers{
		"project": {
			"name": filepath.Base(dst),
		},
	}

	proj := project.New(osfs.New(dst))
	if err := proj.Create(cfg, answers); err != nil {
		return err
	}

	return tmpl.Clone(proj)
}

func mkdir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	} else {
		return &errors.ErrExists{
			Path: path,
		}
	}

	if err := os.MkdirAll(path, 0777); err != nil {
		return err
	}

	return nil
}
