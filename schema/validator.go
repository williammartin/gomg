package schema

import (
	"io"
	"io/ioutil"

	"github.com/xeipuuv/gojsonschema"
)

type Validator struct{}

type Result struct {
	IsValid bool
	Errors  ValidationErrors
}

type ValidationErrors []string

func (v *Validator) Validate(schemaReader, documentReader io.Reader) (*Result, error) {
	schema, err := ioutil.ReadAll(schemaReader)
	if err != nil {
		return nil, err
	}

	document, err := ioutil.ReadAll(documentReader)
	if err != nil {
		return nil, err
	}

	schemaLoader := gojsonschema.NewStringLoader(string(schema))
	documentLoader := gojsonschema.NewStringLoader(string(document))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return &Result{}, err
	}

	if result.Valid() {
		return &Result{IsValid: true, Errors: ValidationErrors{}}, nil
	}

	validationErrors := ValidationErrors{}
	for _, e := range result.Errors() {
		validationErrors = append(validationErrors, e.String())
	}

	return &Result{IsValid: false, Errors: validationErrors}, nil
}
