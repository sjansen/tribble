package template

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/sjansen/tribble/internal/errors"
)

// ConfigVersion indicates project template settings compatibility.
// Currently only one version is defined: 0
type ConfigVersion int

const (
	V0 ConfigVersion = iota // unstable / experimental
)

// Config defines Tribble-specific project template settings.
type Config struct {
	Version ConfigVersion `toml:"version"`
}

// LoadConfig reads Tribble-specific project template settings.
func (t *Template) LoadConfig() (*Config, error) {
	f, err := t.fs.Open(configPath)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if _, err = toml.NewDecoder(f).Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

// SaveConfig writes Tribble-specific project template settings.
func (t *Template) SaveConfig(c *Config, force bool) error {
	if !force {
		if _, err := t.fs.Stat(configPath); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		} else {
			return &errors.ErrExists{
				Path: configPath,
			}
		}
	}

	f, err := t.fs.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(c)
}
