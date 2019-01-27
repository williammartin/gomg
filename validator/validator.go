package validator

import (
	"encoding/json"

	"github.com/williammartin/jsonschema"
	"github.com/williammartin/omg"
	"github.com/xeipuuv/gojsonschema"
)

type Validator struct{}

type Result struct {
	IsValid bool
	Errors  ValidationErrors
}

type ValidationErrors []string

func (v *Validator) Validate(microservice *omg.Microservice) (*Result, error) {
	reflector := &jsonschema.Reflector{AllowAdditionalProperties: false, RequiredFromJSONSchemaTags: true}
	js := reflector.Reflect(&omg.Microservice{})
	jm, err := json.Marshal(js)
	if err != nil {
		return nil, err
	}

	schemaLoader := gojsonschema.NewStringLoader(string(jm))
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
