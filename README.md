clistruct
-------

[![Build Status](https://travis-ci.org/corpix/clistruct.svg?branch=master)](https://travis-ci.org/corpix/clistruct)

Go struct mapper for [urfave/cli](https://github.com/urfave/cli).

## Mapping what?

- Struct field types into flags
- Parsed flags into struct fields

## Limitations

- Has no support for `github.com/urfave/cli.Global*` getters(idk how to map them, don't think you will ever need to do this, if you need then tell me your case)
- You can't pass default value for [generic](https://github.com/urfave/cli/blob/6a87e37dffb000993f7c2831579e271d8fb298aa/flag.go#L99) at this time, clear solution required(at this time it will return an error about incompatible types)

> Also, reflection is full of shit so... there could be bugs.
> Feel free to send pull requests or open an issue if you have problems.

## Example

Let's write a simple program which will accept two special flags:

- `--debug`
- `--say`

``` go
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
```

Now let's run it:

``` shell
go run examples/simple/main.go --say="hello"
Here is what I say: hello
```

We define a `bool` `--debug` flag, let's try to run application with it:

``` shell
go run examples/simple/main.go --say="hello" --debug
I am in debug mode
Here is what I say: hello
```

That's it. Little summary:

- You could use `FlagsFromStruct` to construct flags from the structure.
- You could use `FlagsToStruct` to write parsed flags into the structure.

## License

MIT
