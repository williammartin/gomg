package omg

// Microservice is the the top level configuration
type Microservice struct {
	// OMG specifies the version of the OMG specification with which this microservice complies.
	OMG int `json:"omg,omitempty" jsonschema:"required"`

	// Info specifies general information about the microservice.
	Info *Info `json:"info,omitempty" jsonschema:"required"`

	// Lifecycle specifies the lifecycle of the microservice.
	Lifecycle *Lifecycle `json:"lifecycle,omitempty"`

	// Actions specifies how to interact with the microservice.
	Actions Actions `json:"actions,omitempty"`

	// Environment specifies environment variables that should be exposed to the microservice.
	Environment Environment `json:"environment,omitempty"`

	// Volumes specifies volumes that should be provided to the microservice.
	Volumes Volumes `json:"volumes,omitempty"`

	// Healthcheck specifies how the health of the microservice should be checked.
	Healthcheck *Healthcheck `json:"health,omitempty"`

	// System specifies the system requirements of the microservice.
	System *System `json:"system,omitempty"`
}
