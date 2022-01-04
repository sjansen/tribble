package project

import (
	"github.com/BurntSushi/toml"

	"github.com/sjansen/tribble/internal/project/template"
)

// LoadAnswers reads Tribble-specific project settings.
func (p *Project) LoadAnswers() (*template.Answers, error) {
	f, err := p.fs.Open(answersPath)
	if err != nil {
		return nil, err
	}

	a := &template.Answers{}
	if _, err = toml.NewDecoder(f).Decode(a); err != nil {
		return nil, err
	}

	return a, nil
}

// SaveAnswers writes Tribble-specific project settings.
func (p *Project) SaveAnswers(a *template.Answers) error {
	f, err := p.fs.Create(answersPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(a)
}
