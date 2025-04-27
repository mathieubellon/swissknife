package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const ChangelogPath = "CHANGELOG.json"

type Changelog struct {
	ReleaseDate    string          `json:"release_date"`
	DetectorGroups []DetectorGroup `json:"detector_groups"`
	Engine         []string        `json:"engine"`
	Misc           []string        `json:"misc"`
}

type DetectorGroup struct {
	Name    string   `json:"detector_group"`
	URL     string   `json:"url,omitempty"`
	Changes []Change `json:"changes"`
}

type Change struct {
	NewDetector     string `json:"new_detector,omitempty"`
	NewChecker      string `json:"new_checker,omitempty"`
	DetectorUpgrade string `json:"detector_upgrade,omitempty"`
	CheckerUpgrade  string `json:"checker_upgrade,omitempty"`
	NewAnalyzer     string `json:"new_analyzer,omitempty"`
	AnalyzerUpgrade string `json:"analyzer_upgrade,omitempty"`
}

func buildOutput(data []byte, repo string) {

	var changelog Changelog
	if err := json.Unmarshal(data, &changelog); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	alldetectors, err := getDetectorsList(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting detectors list: %v\n", err)
		os.Exit(1)
	}

	// Process each detector group
	for i, detectorGroup := range changelog.DetectorGroups {
		// Find matching detector in alldetectors
		for _, detector := range alldetectors {
			if detector.GroupName == detectorGroup.Name {
				// Found a match, update the URL in the original changelog
				changelog.DetectorGroups[i].URL = detector.URL
				break
			}
		}
	}

	// Convert changelog to JSON
	jsonOutput, err := json.MarshalIndent(changelog, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating JSON output: %v\n", err)
		os.Exit(1)
	}

	// Print the JSON to stdout
	fmt.Println(string(jsonOutput))
}

func generateChangelog(Ctx *cli.Context) error {
	repo := Ctx.String("repo")
	version := Ctx.String("version")

	// Read JSON file
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", repo, ChangelogPath))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer jsonFile.Close()

	// Read file content
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse JSON
	var changelog map[string]interface{}
	err = json.Unmarshal(byteValue, &changelog)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	if content, exists := changelog[version]; exists {
		// Convert to pretty JSON for display
		prettyJSON, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			log.Fatalf("Error formatting JSON: %v", err)
		}
		buildOutput(prettyJSON, repo)
	} else {
		fmt.Printf("Version %s not found in changelog\n", version)
	}

	return nil
}
