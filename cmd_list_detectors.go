package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

func listDetectorsList(Ctx *cli.Context) error {
	repo := Ctx.String("repo")

	detectors, err := getDetectorsList(repo)
	if err != nil {
		return err
	}

	detectorJSON, err := json.Marshal(detectors)
	if err != nil {
		return fmt.Errorf("error marshaling detectors to JSON: %w", err)
	}

	fmt.Println(string(detectorJSON))

	return nil
}
