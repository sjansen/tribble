package template

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-git/go-billy/v5"
)

const configPath = "_tribble/config.toml"

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
func LoadConfig(fs billy.Filesystem) (*Config, error) {
	f, err := fs.Open(configPath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	_, err = toml.NewDecoder(f).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// SaveConfig writes Tribble-specific project template settings.
func SaveConfig(fs billy.Filesystem, config *Config, force bool) error {
	if !force {
		_, err := fs.Stat(configPath)
		if err == nil {
			return &ErrExists{
				Path: configPath,
			}
		}
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	f, err := fs.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(config)
}
