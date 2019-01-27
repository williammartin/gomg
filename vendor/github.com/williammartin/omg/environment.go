package omg

// Environment specifies a map of names to environment variables.
type Environment map[string]*Variable

// Variable specifies the details of an environment variable.
type Variable struct {
	// Help specifies a human friendly description for the environment variable.
	Help string `json:"help,omitempty"`

	// Type specifies the type of the environment variable.
	Type string `json:"type,omitempty" jsonschema:"required,enum=int|float|string|boolean"`

	// Pattern specifies the a pattern which the environment variable must match.
	Pattern string `json:"pattern,omitempty"`

	// Required specifies whether the environment variable is required.
	Required bool `json:"required,omitempty"`
}
