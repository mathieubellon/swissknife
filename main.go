package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Swissknife"
	app.Version = "0.2"
	app.Usage = "https://github.com/mathieubellon/swissknife"
	app.Description = "Swissknife is a multi-purposes utility command-line tool for managing detectors.\nIt can be used to generate markdown changelog links from the specified Detection Engine version."
	app.Commands = []*cli.Command{
		{
			Name:   "gencert",
			Usage:  "Generate an example of self-signed certificate",
			Action: gen_certificate,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "host",
					Usage:    "Comma-separated hostnames and IPs to generate a certificate for",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "start-date",
					Usage: "Creation date formatted as Jan 1 15:04:05 2011",
				},
				&cli.DurationFlag{
					Name:  "duration",
					Usage: "Duration that certificate is valid for",
					Value: 365 * 24 * time.Hour, // Default value
				},
				&cli.BoolFlag{
					Name:  "ca",
					Usage: "Whether this cert should be its own Certificate Authority",
				},
				&cli.IntFlag{
					Name:  "rsa-bits",
					Usage: "Size of RSA key to generate. Ignored if --ecdsa-curve is set",
					Value: 2048, // Default value
				},
				&cli.StringFlag{
					Name:  "ecdsa-curve",
					Usage: "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521",
				},
				&cli.BoolFlag{
					Name:  "ed25519",
					Usage: "Generate an Ed25519 key",
					Value: false, // Default value
				},
			},
		},
		{
			Name:   "markdown",
			Usage:  "Generate markdown changelog links from the specified Detection Engine version",
			Action: generateMarkdown,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "version",
					Usage:    "Specify the Tokenscanner version (e.g. 2.115.0)",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "repo",
					Usage: "Specify the Tokenscanner repo local path",
					Value: ".", // Default value
				},
				&cli.BoolFlag{
					Name:  "absolute-url",
					Usage: "Use absolute URLs in the markdown output",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
