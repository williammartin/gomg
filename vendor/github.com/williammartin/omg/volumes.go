package omg

// Volume specifies a map of names to volumes.
type Volumes map[string]*Volume

// Volume specifies the details of an volume.
type Volume struct {
	// Target specifies the directory this volume should be mounted.
	Target string `json:"target,omitempty" jsonschema:"required"`

	// Persist specifies whether the volume must be persisted.
	Persist bool `json:"persist,omitempty"`
}
