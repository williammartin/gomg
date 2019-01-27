package omg

// Healthcheck specifies how to check the health of a microservice.
type Healthcheck struct {
	// Interval specifies how often the healthcheck should be executed in seconds.
	Interval int `json:"interval,omitempty"`

	// Timeout specifies how long to wait before assuming the healthcheck failed in seconds.
	Timeout int `json:"timeout,omitempty"`

	// Retries specifies the number of times the healthcheck should be retried.
	Retries int `json:"retries,omitempty"`

	// Command specifies the command to execute to check the health of the microservice.
	Command []string `json:"command,omitempty"`
}
