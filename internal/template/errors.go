package template

import (
	"fmt"
	"io/fs"
)

type ErrExists struct {
	Path string
}

func (e *ErrExists) Error() string {
	return fmt.Sprintf("%q already exists", e.Path)
}

func (e *ErrExists) Unwrap() error {
	return fs.ErrExist
}
