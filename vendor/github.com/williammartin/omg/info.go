package omg

type Info struct {
	// Version specifies version of the microservice.
	Version string `json:"version,omitempty" jsonschema:"required,pattern=[0-9]*\\.[0-9]*\\.[0-9]*"`

	// Title specifies a human friendly name.
	Title string `json:"title,omitempty" jsonschema:"required"`

	// Description specifies a human friendly description.
	Description string `json:"description,omitempty" jsonschema:"required"`

	// Contact specifies a contact point.
	Contact *Contact `json:"contact,omitempty"`

	// License describes the licensing model.
	License *License `json:"license,omitempty" jsonschema:"required"`
}

type Contact struct {
	// Name is the name of the person or company.
	Name string `json:"name,omitempty"`

	// URL is the homepage of the person or company.
	URL string `json:"url,omitempty"`

	// Email is the email address of the person or company.
	Email string `json:"email,omitempty"`
}

type License struct {
	// License is the name of the license.
	Name string `json:"name,omitempty" jsonschema:"required"`

	// URL is a url to an indepth explanation of the license.
	URL string `json:"url,omitempty" jsonschema:"required"`
}
