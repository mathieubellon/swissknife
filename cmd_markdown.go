package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type Release struct {
	ReleaseDate       time.Time         `json:"release_date"`
	AddedDetectors    AddedDetectors    `json:"added_detectors"`
	RemovedDetectors  RemovedDetectors  `json:"removed_detectors"`
	ModifiedDetectors ModifiedDetectors `json:"modified_detectors"`
}

type AddedDetectors struct {
	Generic  []string `json:"generic"`
	Specific []string `json:"specific"`
}

type RemovedDetectors struct {
	Generic  []string `json:"generic"`
	Specific []string `json:"specific"`
}

type ModifiedDetectors struct {
	Generic  []string `json:"generic"`
	Specific []string `json:"specific"`
}

func generateMarkdown(Ctx *cli.Context) error {
	basepath := ".."
	repo := Ctx.String("repo")
	version := Ctx.String("version")
	if Ctx.Bool("absolute-url") {
		basepath = "https://docs.gitguardian.com"
	}

	filePath := fmt.Sprintf("%s/tokenscanner/tokenscanner/data/DETECTORS_CHANGELOG.json", repo)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	value := data[version]

	var release Release

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error converting value to []byte: %w", err)
	}

	err = json.Unmarshal(valueBytes, &release)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	fmt.Printf("Added %v Detectors:\n", len(release.AddedDetectors.Generic)+len(release.AddedDetectors.Specific))
	for _, detector := range release.AddedDetectors.Generic {
		markdownURL, err := buildMarkdownURL(repo, basepath, "generics", detector)
		if err != nil {
			return err
		}
		fmt.Println(markdownURL)
	}
	for _, detector := range release.AddedDetectors.Specific {
		markdownURL, err := buildMarkdownURL(repo, basepath, "specifics", detector)
		if err != nil {
			return err
		}
		fmt.Println(markdownURL)
	}

	fmt.Printf("Modified %v Detectors:\n", len(release.ModifiedDetectors.Generic)+len(release.ModifiedDetectors.Specific))
	for _, detector := range release.ModifiedDetectors.Generic {
		markdownURL, err := buildMarkdownURL(repo, basepath, "generics", detector)
		if err != nil {
			return err
		}
		fmt.Println(markdownURL)
	}
	for _, detector := range release.ModifiedDetectors.Specific {
		markdownURL, err := buildMarkdownURL(repo, basepath, "specifics", detector)
		if err != nil {
			return err
		}
		fmt.Println(markdownURL)
	}

	fmt.Printf("Modified %v Detectors:\n", len(release.RemovedDetectors.Generic)+len(release.RemovedDetectors.Specific))
	for _, detector := range release.RemovedDetectors.Generic {
		markdownURL, err := buildMarkdownURL(repo, basepath, "generics", detector)
		if err != nil {
			return err
		}
		fmt.Println(markdownURL)
	}
	for _, detector := range release.RemovedDetectors.Specific {
		markdownURL, err := buildMarkdownURL(repo, basepath, "specifics", detector)
		if err != nil {
			return err
		}
		fmt.Println(markdownURL)
	}

	return nil
}

func findDetectorYAML(directory, detectorName string) (string, error) {
	dir := fmt.Sprintf("%s/tokenscanner/tokenscanner/config", directory)
	findThisFile := detectorName + ".yaml"
	var detectorYAML string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), findThisFile) {
			detectorYAML = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "Walk error (right path?)", err
	}

	if detectorYAML == "" {
		return "", fmt.Errorf("no detector.yaml file found")
	}

	return detectorYAML, nil
}

func extractDisplayNameFromYAML(filePath string) (string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(fileContent, &data)
	if err != nil {
		return "", fmt.Errorf("error parsing YAML: %w", err)
	}

	displayName, ok := data["display_name"].(string)
	if !ok {
		return "", fmt.Errorf("display_name not found or not a string")
	}

	return displayName, nil
}

func buildMarkdownURL(repo, basepath, category, detector string) (string, error) {
	filePath, err := findDetectorYAML(repo, detector)
	if err != nil {
		return "", err
	}
	displayName, err := extractDisplayNameFromYAML(filePath)
	if err != nil {
		return "", err
	}
	markdownURL := fmt.Sprintf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/%s/%s)", displayName, basepath, category, detector)
	return markdownURL, nil
}
