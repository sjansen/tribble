package template

import "github.com/go-git/go-billy/v5"

// Initialize creates required files for a Tribble project template.
func Initialize(fs billy.Filesystem, force bool) error {
	err := SaveConfig(fs, &Config{}, force)
	if err != nil {
		return err
	}

	vars := &Variables{
		"project": {
			"name": {
				Required: true,
			},
			"owner": {
				Required: false,
			},
		},
	}
	return SaveVariables(fs, vars, force)
}
