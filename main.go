package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var ValidateCommand = cli.Command{
	Name: "validate",
	Action: func(ctx *cli.Context) error {
		if _, err := os.Stat("microservice.yml"); os.IsNotExist(err) {
			return cli.NewExitError("the current directory must contain a 'microservice.yml' file", 1)
		}

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
