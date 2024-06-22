package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

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

const basepath string = "https://docs.gitguardian.com"

func main() {
	var versionFlag string
	flag.StringVar(&versionFlag, "version", "", "Specify the version")
	flag.Parse()

	if versionFlag == "" {
		fmt.Println("Version is mandatory ex: -version 2.115.0")
		return
	}

	version := versionFlag
	// Read the JSON file
	filePath := "/Users/mathieu.bellon/Desktop/tokenscanner/tokenscanner/data/DETECTORS_CHANGELOG.json"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the JSON data
	var data map[string]interface{}
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Retrieve the value at key "2.115.0"
	value := data[version]

	var release Release

	valueBytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error converting value to []byte:", err)
		return
	}

	err = json.Unmarshal(valueBytes, &release)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Printf("Added %v Detectors:\n", len(release.AddedDetectors.Generic)+len(release.AddedDetectors.Specific))
	for _, detector := range release.AddedDetectors.Generic {
		markdownURL, err := buildMarkdownURL(basepath, "generics", detector)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(markdownURL)
	}
	for _, detector := range release.AddedDetectors.Specific {
		markdownURL, err := buildMarkdownURL(basepath, "specifics", detector)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(markdownURL)
	}

	fmt.Printf("Modified %v Detectors:\n", len(release.ModifiedDetectors.Generic)+len(release.ModifiedDetectors.Specific))
	for _, detector := range release.ModifiedDetectors.Generic {
		markdownURL, err := buildMarkdownURL(basepath, "generics", detector)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(markdownURL)
	}
	for _, detector := range release.ModifiedDetectors.Specific {
		markdownURL, err := buildMarkdownURL(basepath, "specifics", detector)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(markdownURL)
	}

	fmt.Printf("Modified %v Detectors:\n", len(release.RemovedDetectors.Generic)+len(release.RemovedDetectors.Specific))
	for _, detector := range release.RemovedDetectors.Generic {
		markdownURL, err := buildMarkdownURL(basepath, "generics", detector)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(markdownURL)
	}
	for _, detector := range release.RemovedDetectors.Specific {
		markdownURL, err := buildMarkdownURL(basepath, "specifics", detector)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(markdownURL)
	}
}

func findDetectorYAML(detectorName string) (string, error) {
	dir := "/Users/mathieu.bellon/Desktop/tokenscanner/tokenscanner/config"
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
	// Read the YAML file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// Parse the YAML data
	var data map[string]interface{}
	err = yaml.Unmarshal(fileContent, &data)
	if err != nil {
		return "", fmt.Errorf("error parsing YAML: %w", err)
	}

	// Retrieve the value of display_name
	displayName, ok := data["display_name"].(string)
	if !ok {
		return "", fmt.Errorf("display_name not found or not a string")
	}

	return displayName, nil
}

func buildMarkdownURL(basepath, category, detector string) (string, error) {
	filePath, err := findDetectorYAML(detector)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	displayName, err := extractDisplayNameFromYAML(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	markdownURL := fmt.Sprintf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/%s/%s)", displayName, basepath, category, detector)
	return markdownURL, nil
}
