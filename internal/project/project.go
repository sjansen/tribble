package project

import (
	"os"

	"github.com/go-git/go-billy/v5"
)

const answersPath = ".tribble/answers.toml"
const configPath = ".tribble/config.toml"

type Project struct {
	fs billy.Filesystem
}

func New(fs billy.Filesystem) *Project {
	return &Project{
		fs: fs,
	}
}

func (p *Project) MkdirAll(path string, perm os.FileMode) error {
	return p.fs.MkdirAll(path, perm)
}

func (p *Project) WriteFile(path string, data []byte, perm os.FileMode) error {
	f, err := p.fs.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}
