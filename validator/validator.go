package validator

import (
	"github.com/williammartin/omg"
	"github.com/xeipuuv/gojsonschema"
)

type Validator struct {
	MicroserviceSchema string
}

type Result struct {
	IsValid bool
	Errors  ValidationErrors
}

type ValidationErrors []string

func (v *Validator) Validate(microservice *omg.Microservice) (*Result, error) {
	schemaLoader := gojsonschema.NewReferenceLoader(v.MicroserviceSchema)
	documentLoader := gojsonschema.NewGoLoader(microservice)

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
