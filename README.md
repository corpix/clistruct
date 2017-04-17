clistruct
-------

[![Build Status](https://travis-ci.org/corpix/clistruct.svg?branch=master)](https://travis-ci.org/corpix/clistruct)

Go struct mapper for [urfave/cli](https://github.com/urfave/cli).

## Mapping what?

- Struct field types into flags
- Parsed flags into struct fields

## Limitations

- Has no support for `github.com/urfave/cli.Global*` getters(idk how to map them, don't think you will ever need to do this, if you need then tell me your case)
- You can't pass default value as struct tag at this time, they require additional parsing to be implemented in this library

## License

MIT
