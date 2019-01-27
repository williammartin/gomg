package omg

type Lifecycle struct {
	// Startup specifies the startup configuration for the microservice.
	Startup *Startup `json:"startup,omitempty"`

	// Shutdown specifies the shutdown configuration for the microservice.
	Shutdown *Shutdown `json:"shutdown,omitempty"`
}

type Startup struct {
	// Command is the startup command that should be executed. It must be a string or array of strings.
	Command interface{} `json:"command,omitempty"`
}

type Shutdown struct {
	// Command is the shutdown command that should be executed. It must be a string or array of strings.
	Command interface{} `json:"command,omitempty"`

	// Timeout is the time allowed to gracefully shutdown.
	Timeout int `json:"timeout,omitempty"`
}
