package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ghodss/yaml"
	"github.com/williammartin/gomg/ui"
	"github.com/williammartin/gomg/validator"
	"github.com/williammartin/omg"
)

var ValidateCommand = cli.Command{
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
			cli.NewExitError(err, 1)
		}

		if !result.IsValid {
			fmt.Fprintln(os.Stderr, "validation errors occurred:")
			for _, e := range result.Errors {
				fmt.Fprintf(os.Stderr, " - %s\n", e)
			}
			UI.DisplayNewline()
			UI.DisplayFailed()
			return cli.NewExitError("", 1)
		}

		UI.DisplayText("validation succeeded\n")
		UI.DisplayNewline()
		UI.DisplaySuccess()

		return nil
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "gomg"
	app.Commands = []cli.Command{
		ValidateCommand,
	}

	app.Run(os.Args)
}
