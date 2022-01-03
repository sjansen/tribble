package variables

// Question is a project-specific setting.
type Question struct {
	Default  interface{} `toml:"default"`
	Type     string      `toml:"type,omitempty"`
	Required bool        `toml:"required"`
}

// Questions are a collection of project-specific settings.
type Questions map[string]map[string]*Question
