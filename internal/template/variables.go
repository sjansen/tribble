package template

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-git/go-billy/v5"
)

const variablesPath = "_tribble/variables.toml"

// Variable is a project-specific setting.
type Variable struct {
	Type     string      `toml:"type,omitempty"`
	Default  interface{} `toml:"default"`
	Required bool        `toml:"required"`
}

// Variables define project-specific settings.
type Variables map[string]map[string]*Variable

// LoadVariables reads Tribble-specific project template settings.
func LoadVariables(fs billy.Filesystem) (*Variables, error) {
	f, err := fs.Open(variablesPath)
	if err != nil {
		return nil, err
	}

	variables := &Variables{}
	_, err = toml.NewDecoder(f).Decode(variables)
	if err != nil {
		return nil, err
	}

	return variables, nil
}

// SaveVariables writes Tribble-specific project template settings.
func SaveVariables(fs billy.Filesystem, variables *Variables, force bool) error {
	if !force {
		_, err := fs.Stat(variablesPath)
		if err == nil {
			return &ErrExists{
				Path: variablesPath,
			}
		}
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	f, err := fs.Create(variablesPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(variables)
}
