package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/williammartin/gomg/build"
	"github.com/williammartin/gomg/ui"
	"github.com/williammartin/gomg/validate"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomg"
	app.Commands = []cli.Command{
		validate.Command,
		build.Command,
	}

	UI := &ui.UI{
		Out: os.Stdout,
		Err: os.Stderr,
	}

	if err := app.Run(os.Args); err != nil {
		UI.DisplayErrorAndFailed(err)
		os.Exit(1)
	}

	UI.DisplaySuccess()
}
