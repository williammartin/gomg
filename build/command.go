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
	"github.com/williammartin/gomg/ui"
	"github.com/williammartin/gomg/validator"
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

		if _, err := os.Stat("microservice.yml"); os.IsNotExist(err) {
			UI.DisplayErrorAndFailed(errors.New("the current directory must contain a 'microservice.yml' file"))
			return cli.NewExitError("", 1)
		}

		b, err := ioutil.ReadFile("microservice.yml")
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		var microservice omg.Microservice
		err = yaml.Unmarshal(b, &microservice)
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		validator := &validator.Validator{}
		result, err := validator.Validate(&microservice)
		if err != nil {
			UI.DisplayErrorAndFailed(err)
			cli.NewExitError("", 1)
		}

		if !result.IsValid {
			fmt.Fprintln(os.Stderr, "validation errors occurred:")
			for _, e := range result.Errors {
				UI.DisplayError(fmt.Errorf(" - %s", e))
			}
			UI.DisplayFailed()
			return cli.NewExitError("", 1)
		}

		if _, err := os.Stat("Dockerfile"); os.IsNotExist(err) {
			return cli.NewExitError("the current directory must contain a 'Dockerfile' file", 1)
		}

		UI.DisplayText("building...")

		endpoint := "unix:///var/run/docker.sock"
		dockerClient, err := docker.NewClient(endpoint)
		if err != nil {
			cli.NewExitError(err, 1)
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
			return cli.NewExitError(err, 1)
		}

		UI.DisplayText("built {{.Repository}} with tag latest", map[string]interface{}{"Repository": repository})
		UI.DisplaySuccess()

		return nil
	},
}

func convertTitleToImageName(title string) string {
	return strings.ToLower(strings.Replace(title, " ", "", -1))
}
