package validate

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
	"github.com/williammartin/gomg/ui"
	"github.com/williammartin/gomg/validator"
	"github.com/williammartin/omg"
	yaml "gopkg.in/yaml.v2"
)

var Command = cli.Command{
	Name: "validate",
	Action: func(ctx *cli.Context) error {
		UI := &ui.UI{
			Out: os.Stdout,
			Err: os.Stderr,
		}

		if _, err := os.Stat("microservice.yml"); os.IsNotExist(err) {
			return cli.NewExitError("the current directory must contain a 'microservice.yml' file", 1)
		}

		bytes, err := ioutil.ReadFile("microservice.yml")
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		var microservice omg.Microservice
		err = yaml.Unmarshal(bytes, &microservice)
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		validator := &validator.Validator{
			MicroserviceSchema: "file:///Users/wmartin/go/src/github.com/williammartin/omg/schemas/microservice.json",
		}

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

		UI.DisplayText("validation succeeded")
		UI.DisplaySuccess()

		return nil
	},
}
