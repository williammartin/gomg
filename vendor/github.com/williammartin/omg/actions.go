package omg

// Actions specifies a map of action name to action.
type Actions map[string]*Action

// Action specifies an microservice action.
type Action struct {
	// Help specifies a human friendly description.
	Help string `json:"help,omitempty"`

	// Format specifies that this is a CLI based microservice.
	Format *Format `json:"format,omitempty" jsonschema:"oneOf"`

	// HTTP specifies that this is an http based microservice.
	HTTP *HTTP `json:"http,omitempty" jsonschema:"oneOf"`

	// Arguments specifes named arguments.
	Arguments Arguments `json:"arguments,omitempty"`

	// Output specifies the structure of data returned.
	Output *Output `json:"output,omitempty"`
}

// Format specifies a CLI action.
type Format struct {
	// Command specifies the command to be executed. It must be a string or array of strings.
	Command []string `json:"command,omitempty" jsonschema:"required"`
}

// HTTP specifies an HTTP action.
type HTTP struct {
	// Port specifies the port on which a connection should be established.
	Port int `json:"port,omitempty" jsonschema:"required"`

	// Method specifies the HTTP method to be used.
	Method string `json:"method,omitempty" jsonschema:"required"`

	// Path specifies the path on which this action should be executed.
	Path string `json:"path,omitempty" jsonschema:"required"`

	// ContentType specifies the type of the request body. If any microservice arguments
	// have their `in` field set to `requestBody`, they should be encoded with this content type.
	ContentType string `json:"contentType,omitempty"`
}

// Arguments specifies a map of argument name to argument.
type Arguments map[string]*Argument

// Argument specifies the details of a action's argument.
type Argument struct {
	// Help specifies help text.
	Help string `json:"help,omitempty"`

	// Type specifies the data type e.g. int, float, string, etc
	Type string `json:"type,omitempty" jsonschema:"required,enum=int|float|string|list|map|boolean|enum"`

	// In specifies the location of this argument e.g. requestBody, query or path.
	In string `json:"in,omitempty" jsonschema:"required,enum=requestBody|query|path"`

	// Required specifies whether this argument is required or not.
	Required bool `json:"required,omitempty"`

	// Pattern specifies a regex pattern which this argument must match.
	Pattern string `json:"pattern,omitempty"`

	// Enum specifies a list of options which this argument can be.
	Enum []string `json:"enum,omitempty"`

	// Range specifies a min and max bounds for this argument.
	Range *Range `json:"range,omitempty"`
}

// Range specifies a min and max bounds.
type Range struct {
	// Min specifies the minimum value of the range
	Min int `json:"min,omitempty"`

	// Max specifies the maximum value of the range
	Max int `json:"max,omitempty"`
}

// Output specifies the details of a action's output.
type Output struct {
	// Type specifies the data type e.g. int, float, string, object, etc.
	Type string `json:"type,omitempty" jsonschema:"required,enum=int|float|string|list|map|boolean|object"`

	// ContentType specifies the content-type if the type was object.
	ContentType string `json:"contentType,omitempty"`

	// Properties specifies the properties available to the caller.
	Properties map[string]*Property `json:"properties,omitempty"`
}

// Proprty specifies details about the content of an output.
type Property struct {
	// Type specifies the data type e.g. int, float, string, object, etc.
	Type string `json:"type,omitempty" jsonschema:"required,enum=int|float|string|list|map|boolean|object"`
}
