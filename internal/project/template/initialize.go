package template

import (
	"github.com/sjansen/tribble/internal/variables"
)

// Initialize creates the files required for a project template.
func (t *Template) Initialize(force bool) error {
	c := &Config{}
	err := t.SaveConfig(c, force)
	if err != nil {
		return err
	}

	q := &variables.Questions{
		"project": {
			"name": {
				Required: true,
			},
			"owner": {
				Required: false,
			},
		},
	}
	return t.SaveQuestions(q, force)
}
