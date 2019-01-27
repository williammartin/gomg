package omg

// Healthcheck specifies how to check the health of a microservice.
type System struct {
	// Soft specifies soft limits for the microservice.
	Soft *Limits `json:"soft,omitempty"`

	// Hard specifies hard limits for the microservice.
	Hard *Limits `json:"hard,omitempty"`
}

// Limits specify various system limits.
type Limits struct {
	// CPU specifies as a float the percentage of CPU.
	CPU float64 `json:"cpu,omitempty"`

	// Memory specifies as a string representation e.g. `10MB` an amount of memory.
	Memory string `json:"memory,omitempty"`
}
