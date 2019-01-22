package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/williammartin/gomg/build"
	"github.com/williammartin/gomg/validate"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomg"
	app.Commands = []cli.Command{
		validate.Command,
		build.Command,
	}

	app.Run(os.Args)
}
