package project

import (
	"github.com/sjansen/tribble/internal/variables"
)

func (p *Project) Create(c *Config, a *variables.Answers) error {
	if err := p.SaveConfig(c); err != nil {
		return err
	}

	if err := p.SaveAnswers(a); err != nil {
		return err
	}

	return nil
}
