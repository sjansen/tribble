package template

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/sjansen/tribble/internal/errors"
)

// Question is a project-specific setting.
type Question struct {
	Default  interface{} `toml:"default"`
	Type     string      `toml:"type,omitempty"`
	Required bool        `toml:"required"`
}

// Questions are a collection of project-specific settings.
type Questions map[string]map[string]*Question

// LoadQuestions reads Tribble-specific project template settings.
func (t *Template) LoadQuestions() (*Questions, error) {
	f, err := t.fs.Open(questionsPath)
	if err != nil {
		return nil, err
	}

	q := &Questions{}
	if _, err = toml.NewDecoder(f).Decode(q); err != nil {
		return nil, err
	}

	return q, nil
}

// SaveQuestions writes Tribble-specific project template settings.
func (t *Template) SaveQuestions(q *Questions, force bool) error {
	if !force {
		if _, err := t.fs.Stat(questionsPath); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		} else {
			return &errors.ErrExists{
				Path: questionsPath,
			}
		}
	}

	f, err := t.fs.Create(questionsPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(q)
}
