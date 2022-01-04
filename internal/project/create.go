package project

import "github.com/sjansen/tribble/internal/project/template"

func (p *Project) Create(c *Config, a *template.Answers) error {
	if err := p.SaveConfig(c); err != nil {
		return err
	}

	if err := p.SaveAnswers(a); err != nil {
		return err
	}

	return nil
}
