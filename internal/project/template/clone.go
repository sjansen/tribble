package template

import (
	"io"
	"os"
)

type Project interface {
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(path string, data []byte, perm os.FileMode) error
}

func (t *Template) Clone(p Project) error {
	t.clone("/", p)
	return nil
}

func (t *Template) clone(path string, proj Project) error {
	files, err := t.fs.ReadDir(path)
	if err != nil {
		return err
	}

	for _, fi := range files {
		name := fi.Name()
		switch name {
		case ".git", ".tribble", "_tribble":
			continue
		}

		path := t.fs.Join(path, name)
		switch mode := fi.Mode(); {
		case mode.IsDir():
			if err = proj.MkdirAll(path, fi.Mode()); err != nil {
				return err
			}
			if err = t.clone(path, proj); err != nil {
				return err
			}
		case mode.IsRegular():
			if err = t.copy(path, fi, proj); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Template) copy(path string, fi os.FileInfo, proj Project) error {
	f, err := t.fs.Open(fi.Name())
	if err != nil {
		return err
	}

	buf := make([]byte, fi.Size())
	_, err = io.ReadFull(f, buf)
	if err != nil {
		return err
	}

	err = proj.WriteFile(path, buf, fi.Mode())
	if err != nil {
		return err
	}

	return nil
}
