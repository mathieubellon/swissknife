package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const GitGuardianBasePath = "https://docs.gitguardian.com"

func main() {
	app := cli.NewApp()
	app.Name = "Swissknife"
	app.Version = "0.6.0"
	app.Usage = "https://github.com/mathieubellon/swissknife"
	app.Description = "Swissknife is a multi-purposes utility command-line tool for managing detectors.\nIt can be used to generate markdown changelog links from the specified Detection Engine version."
	app.Commands = []*cli.Command{
		{
			Name:   "changelog",
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
		{
			Name:   "vscode",
			Usage:  "Get VScode competitors download count and save to Supabase",
			Action: vscode,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "verbose",
					Usage: "Print verbose output",
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
