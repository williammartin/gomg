package validate

import (
	"io"

	"github.com/williammartin/gomg/schema"
	"github.com/williammartin/omg"
)

//go:generate counterfeiter . SchemaGenerator

type SchemaGenerator interface {
	Generate(interface{}) (io.Reader, error)
}

//go:generate counterfeiter . DocumentGenerator

type DocumentGenerator interface {
	Generate(interface{}) (io.Reader, error)
}

//go:generate counterfeiter . SchemaValidator

type SchemaValidator interface {
	Validate(schemaReader, documentReader io.Reader) (*schema.Result, error)
}

type Actor struct {
	SchemaGenerator   SchemaGenerator
	DocumentGenerator DocumentGenerator
	SchemaValidator   SchemaValidator
}

func (a *Actor) ValidateMicroservice(microservice *omg.Microservice) (*schema.Result, error) {
	schema, err := a.SchemaGenerator.Generate(microservice)
	if err != nil {
		return nil, err
	}

	document, err := a.DocumentGenerator.Generate(microservice)
	if err != nil {
		return nil, err
	}

	return a.SchemaValidator.Validate(schema, document)
}
