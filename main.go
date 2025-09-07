package main

// gookit/slog, below, conflicts with the stdlib slog
// namespace.  For now, I'm keeping it because I
// like the output off gookit/slot more.
import (
	"context"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/gookit/slog"
	probing "github.com/prometheus-community/pro-bing"
	cli "github.com/urfave/cli/v3"
)

func main() {
	// Define the main ping command for the application
	app := &cli.Command{
		Name:    "pinger",
		Usage:   "A simple ping application",
		Version: "0.0.1",
		Flags: []cli.Flag{
			// Define an IntFlag as --count
			&cli.IntFlag{
				Name:    "count",
				Aliases: []string{"c"}, // Process -c
				Value:   10,            // Default value
				Usage:   "Number of times to ping",
			},
			// Define an IntFlag as --size
			&cli.IntFlag{
				Name:    "size",
				Aliases: []string{"s"}, // Process -s
				Value:   100,           // Default value
				Usage:   "Size of the ping payload",
			},
			// Define an Float64Flag as --interval
			&cli.Float64Flag{
				Name:    "interval",
				Aliases: []string{"i"}, // Process -i
				Value:   100.0,         // Default value
				Usage:   "ping interval (milliseconds); default is 100ms",
			},
		},

		//////////////////////////////////////////////////////
		// Set up a deferred action via urfave/cli here...
		//////////////////////////////////////////////////////
		Action: func(_ context.Context, cmd *cli.Command) error {
			/* no error will be returned.  Errors are handled inline */

			// Ensure that the program runs as root
			if os.Geteuid() != 0 {
				_error := "Program must run as root"
				slog.Fatal(_error)
				os.Exit(1)
			}

			// Get the CLI positional argument (hostname)
			hostname := cmd.Args().Get(0)

			// Access urfave/cli flag values...
			count := cmd.Int("count")
			size := cmd.Int("size")
			interval := cmd.Float64("interval")

			statistics := ping(hostname, count, size, interval)

			summary := "Packet loss: " + strconv.FormatFloat(statistics.PacketLoss, 'f', 3, 64) + "%"
			slog.Info(summary)

			return nil
		},
	}

	///////////////////////////////////////////////////////////////
	// Run the deferred application with command-line arguments
	///////////////////////////////////////////////////////////////
	if err := app.Run(context.Background(), os.Args); err != nil {
		slog.Fatal(err.Error())
		os.Exit(1)
	}
}

func ping(hostname string, count int, size int, intervalMs float64) *probing.Statistics {
	/*
		Attempt to ping a remote host with the Prometheus pro-bing library.
	*/

	// Check that the hostname resolves
	hosts, err := net.LookupHost(hostname)
	if len(hosts) == 0 {
		_error := "Invalid hostname: " + hostname + " " + err.Error()
		slog.Fatal(_error)
		panic(_error)
	}

	slog.Debug(hosts)

	if (23 >= size) || (size >= 1470) {
		_error := "Ping payload must be between 24 and 1470 bytes"
		slog.Error(_error)
		os.Exit(1)
	}

	// validate packet count is sane
	if count < 1 {
		_error := "Invalid count.  Cannot ping " + hostname + " with " + strconv.Itoa(count) + " packets"
		slog.Fatal(_error)
		os.Exit(1)
	}

	// validate ping interval is sane
	if intervalMs < 1.0 {
		_error := "Invalid packet interval. Specify an interval of at least 1.0 millisecond"
		slog.Fatal(_error)
		os.Exit(1)
	}

	/////////////////////////////////////////////
	// Set up the pinger instance
	/////////////////////////////////////////////
	pinger, err := probing.NewPinger(hostname)
	if err != nil {
		slog.Fatal(err.Error())
		os.Exit(1)
	}

	slog.Info("pinging " + pinger.Addr() + " " + pinger.IPAddr().String() + " " + strconv.Itoa(count) + " times with " + strconv.Itoa(size) + " byte ICMP payloads")

	// time.Duration() is nanoseconds... convert to milliseconds
	pinger.Interval = time.Duration(intervalMs * 1000000.0)
	pinger.Timeout = time.Duration(float64(count)*intervalMs) * time.Millisecond
	pinger.SetPrivileged(true)
	pinger.SetDoNotFragment(true)
	pinger.Count = count
	pinger.Size = size // ICMP payload size

	err = pinger.Run()
	if err != nil {
		slog.Fatal(err.Error())
		os.Exit(1)
	}

	return pinger.Statistics()
}
