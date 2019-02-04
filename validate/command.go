package validate

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/pivotal-cf/jhanda"
	"github.com/williammartin/gomg/schema"
	"github.com/williammartin/gomg/schema/generator"
	"github.com/williammartin/gomg/ui"
	"github.com/williammartin/jsonschema"
	"github.com/williammartin/omg"
	yaml "gopkg.in/yaml.v2"
)

type Command struct{}

func (c Command) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description: "Validate OMG microservice",
	}
}

func (c Command) Execute(args []string) error {
	UI := &ui.UI{
		Out: os.Stdout,
		Err: os.Stderr,
	}

	microservice, err := loadMicroservice()
	if err != nil {
		return err
	}

	result, err := validateMicroservice(microservice)
	if err != nil {
		return err
	}

	if !result.IsValid {
		return &ValidationFailedError{ValidationErrors: result.Errors}
	}

	UI.DisplayText("validation succeeded")

	return nil
}

func loadMicroservice() (*omg.Microservice, error) {
	if _, err := os.Stat("microservice.yml"); os.IsNotExist(err) {
		return nil, errors.New("the current directory must contain a 'microservice.yml' file")
	}

	bytes, err := ioutil.ReadFile("microservice.yml")
	if err != nil {
		return nil, err
	}

	var microservice omg.Microservice
	err = yaml.Unmarshal(bytes, &microservice)
	if err != nil {
		return nil, err
	}

	return &microservice, nil
}

func validateMicroservice(microservice *omg.Microservice) (*schema.Result, error) {
	reflector := &jsonschema.Reflector{AllowAdditionalProperties: false, RequiredFromJSONSchemaTags: true}
	schemaGenerator := &generator.SchemaGenerator{Reflector: reflector}
	docGenerator := &generator.DocumentGenerator{}
	validator := &schema.Validator{}
	actor := &Actor{SchemaGenerator: schemaGenerator, DocumentGenerator: docGenerator, SchemaValidator: validator}

	return actor.ValidateMicroservice(microservice)
}
