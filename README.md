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

## License

MIT
