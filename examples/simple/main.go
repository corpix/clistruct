package main

import (
	"fmt"

	"github.com/corpix/clistruct"
	"github.com/urfave/cli"
)

var (
	flags = &Flags{}
)

type Flags struct {
	Debug bool   `usage:"Enable debug mode"`
	Say   string `usage:"Tell me what to say" value:"I could say nothing"`
}

func rootAction(context *cli.Context) error {
	if flags.Debug {
		fmt.Println("I am in debug mode")
	}

	fmt.Println(
		"Here is what I say:",
		flags.Say,
	)

	return nil
}

func main() {
	cliFlags, err := clistruct.FlagsFromStruct(flags)
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Flags = cliFlags
	app.Before = func(context *cli.Context) error {
		return clistruct.FlagsToStruct(context, flags)
	}
	app.Action = rootAction

	app.RunAndExitOnError()
}
