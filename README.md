## pinger

This builds a simple golang CLI ping binary using:

- [`probing`][1]
- [`urfave/cli`][2]

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

  [1]: https://github.com/prometheus-community/pro-bing
  [2]: https://github.com/urfave/cli

