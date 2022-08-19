package template

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/sjansen/tribble/internal/errors"
)

// Questions are a collection of project-specific settings.
type Questions map[string]map[string]*Question

// Question is a project-specific setting.
type Question struct {
	Prompt   string      `toml:"prompt"`
	Type     string      `toml:"type,omitempty"`
	Default  interface{} `toml:"default"`
	Required bool        `toml:"required"`
}

func (q *Question) Lint() (errors, warnings []string) {
	if q.Prompt == "" {
		errors = append(errors, "missing required prompt")
	}
	// TODO validate Default is compatible with Type
	switch strings.ToLower(q.Type) {
	case "", "string":
		// noop
	default:
		errors = append(errors, fmt.Sprintf("invalid value type: %q", q.Type))
	}
	return
}

// LoadQuestions reads Tribble-specific project template settings.
func (t *Template) LoadQuestions() (Questions, error) {
	f, err := t.fs.Open(questionsPath)
	if err != nil {
		return nil, err
	}

	q := Questions{}
	if _, err = toml.NewDecoder(f).Decode(&q); err != nil {
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
