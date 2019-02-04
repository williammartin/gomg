package main

import (
	"os"

	"github.com/pivotal-cf/jhanda"
	"github.com/williammartin/gomg/build"
	"github.com/williammartin/gomg/ui"
	"github.com/williammartin/gomg/validate"
)

type globalOptions struct{}

func main() {
	UI := &ui.UI{
		Out: os.Stdout,
		Err: os.Stderr,
	}

	var globalOptions globalOptions

	args, err := jhanda.Parse(&globalOptions, os.Args[1:])
	if err != nil {
		UI.DisplayErrorAndFailed(err)
		os.Exit(1)
	}

	var command string
	if len(args) > 0 {
		command, args = args[0], args[1:]
	}

	commandSet := jhanda.CommandSet{}
	commandSet["validate"] = validate.Command{}
	commandSet["build"] = build.Command{}

	if err := commandSet.Execute(command, args); err != nil {
		UI.DisplayErrorAndFailed(err)
		os.Exit(1)
	}

	UI.DisplaySuccess()
}
