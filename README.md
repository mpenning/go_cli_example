## pinger

This builds a simple golang CLI ping binary using:

- [`probing`][1]
- [`urfave/cli`][2]
- [`gookit/slog`][3]

## Syntax

The pinger must run as root.  CLI syntax follows...

```none
NAME:
   pinger - A simple ping application

USAGE:
   pinger [global options]

VERSION:
   0.0.1

GLOBAL OPTIONS:
   --count int, -c int         Number of times to ping (default: 10)
   --size int, -s int          Size of the ping payload (default: 100)
   --interval float, -i float  ping interval (milliseconds); default is 100ms (default: 100)
   --help, -h                  show help
```

## Changing dependency versions

If you want to change versions of the dependencies, the best way to do so is when they are installed via the `Makefile`.

## Building from source

On linux, use `make all`, it should do the rest (including building static / dynamically linked libraries)

  [1]: https://github.com/prometheus-community/pro-bing
  [2]: https://github.com/urfave/cli
  [3]: https://github.com/gookit/slog


