package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Swissknife"
	app.Version = "0.1"
	app.Usage = "A multi-purposes utility command-line tool for managing detectors"
	app.Commands = []*cli.Command{
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
