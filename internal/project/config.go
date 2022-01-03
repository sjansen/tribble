package project

import (
	"github.com/BurntSushi/toml"
	"github.com/go-git/go-billy/v5"
)

// ConfigVersion indicates project settings compatibility.
// Currently only one version is defined: 0
type ConfigVersion int

const (
	V0 ConfigVersion = iota // unstable / experimental
)

// Config defines Tribble-specific settings.
type Config struct {
	Version ConfigVersion `toml:"version"`
	Origin  struct {
		URL     string `toml:"url"`
		RefName string `toml:"refname,omitempty"`
	} `toml:"origin"`
}

// LoadConfig reads Tribble-specific project settings.
func (p *Project) LoadConfig(fs billy.Filesystem) (*Config, error) {
	f, err := p.fs.Open(configPath)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if _, err = toml.NewDecoder(f).Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

// SaveConfig writes Tribble-specific project settings.
func (p *Project) SaveConfig(c *Config) error {
	f, err := p.fs.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(c)
}
