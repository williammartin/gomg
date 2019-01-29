package build

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
	"github.com/williammartin/gomg/containers"
	"github.com/williammartin/gomg/schema"
	"github.com/williammartin/gomg/schema/generator"
	"github.com/williammartin/gomg/ui"
	"github.com/williammartin/gomg/validate"
	"github.com/williammartin/jsonschema"
	"github.com/williammartin/omg"
	yaml "gopkg.in/yaml.v2"
)

var Command = cli.Command{
	Name: "build",
	Action: func(ctx *cli.Context) error {
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
			return &validate.ValidationFailedError{ValidationErrors: result.Errors}
		}

		if _, err := os.Stat("Dockerfile"); os.IsNotExist(err) {
			return errors.New("the current directory must contain a 'Dockerfile' file")
		}

		UI.DisplayText("building...")

		endpoint := "unix:///var/run/docker.sock"
		dockerClient, err := docker.NewClient(endpoint)
		if err != nil {
			return err
		}

		client := &containers.DockerClient{
			DockerClient: dockerClient,
		}

		r, w := io.Pipe()
		go func(reader io.Reader) {
			UI.DisplayStream(reader)
		}(r)

		repository := fmt.Sprintf("omg-%s", convertTitleToImageName(microservice.Info.Title))
		name := fmt.Sprintf("%s:%s", repository, "latest")
		err = client.Build(name, containers.WithContextDir("."), containers.WithOutputStream(w))
		if err != nil {
			return err
		}

		UI.DisplayText("built {{.Repository}} with tag latest", map[string]interface{}{"Repository": repository})

		return nil
	},
}

func convertTitleToImageName(title string) string {
	return strings.ToLower(strings.Replace(title, " ", "", -1))
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
	actor := &validate.Actor{SchemaGenerator: schemaGenerator, DocumentGenerator: docGenerator, SchemaValidator: validator}

	return actor.ValidateMicroservice(microservice)
}
