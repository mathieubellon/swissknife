package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const GitGuardianBasePath = "https://docs.gitguardian.com"

func main() {
	app := cli.NewApp()
	app.Name = "Swissknife"
	app.Version = "0.5"
	app.Usage = "https://github.com/mathieubellon/swissknife"
	app.Description = "Swissknife is a multi-purposes utility command-line tool for managing detectors.\nIt can be used to generate markdown changelog links from the specified Detection Engine version."
	app.Commands = []*cli.Command{
		{
			Name:   "gpg",
			Usage:  "Generate a GPG key pair",
			Action: gen_gpg,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "Name of the key owner",
					Value: "John Doe", // Default value
				},
				&cli.StringFlag{
					Name:  "email",
					Usage: "Email of the key owner",
					Value: "john.doe@email.com", // Default value
				},
				&cli.StringFlag{
					Name:  "passphrase",
					Usage: "Passphrase to encrypt the private key",
					Value: "LongSecret", // Default value
				},
				&cli.IntFlag{
					Name:  "rsa-bits",
					Usage: "Size of RSA key to generate",
					Value: 2048, // Default value
				},
			},
		},
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
			Name:   "print",
			Usage:  "Generate markdown or html changelog links from the specified Detection Engine version",
			Action: printOutput,
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
				&cli.StringFlag{
					Name:  "format",
					Usage: "markdown or html",
					Value: "markdown", // Default value
				},
			},
		},
		{
			Name:   "list",
			Usage:  "Generate markdown or html detectors list",
			Action: printDetectorsList,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "repo",
					Usage: "Specify the Tokenscanner repo local path",
					Value: ".", // Default value
				},
				&cli.BoolFlag{
					Name:  "write-one-json",
					Usage: "Write the output to disk as one single json file",
				},
				&cli.BoolFlag{
					Name:  "write-multiple-json",
					Usage: "Write the output to disk as multiple json file, one by detector, in a detectors folder",
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
