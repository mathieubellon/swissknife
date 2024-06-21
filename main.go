package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Release struct {
	ReleaseDate       time.Time         `json:"release_date"`
	AddedDetectors    AddedDetectors    `json:"added_detectors"`
	RemovedDetectors  RemovedDetectors  `json:"removed_detectors"`
	ModifiedDetectors ModifiedDetectors `json:"modified_detectors"`
}
type AddedDetectors struct {
	Generic  []any    `json:"generic"`
	Specific []string `json:"specific"`
}
type RemovedDetectors struct {
	Generic  []any `json:"generic"`
	Specific []any `json:"specific"`
}
type ModifiedDetectors struct {
	Generic  []any    `json:"generic"`
	Specific []string `json:"specific"`
}

func main() {
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
	value := data["2.114.0"]

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

	basepath := "https://docs.gitguardian.com"

	fmt.Println("Added Detectors:")
	for _, detector := range release.AddedDetectors.Generic {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/generics/%s)\n", detector, basepath, detector)
	}
	for _, detector := range release.AddedDetectors.Specific {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/specifics/%s)\n", detector, basepath, detector)
	}

	fmt.Println("Modified Detectors:")
	for _, detector := range release.ModifiedDetectors.Generic {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/generics/%s)\n", detector, basepath, detector)
	}
	for _, detector := range release.ModifiedDetectors.Specific {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/specifics/%s)\n", detector, basepath, detector)
	}

	fmt.Println("Removed Detectors:")
	for _, detector := range release.RemovedDetectors.Generic {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/generics/%s)\n", detector, basepath, detector)
	}
	for _, detector := range release.RemovedDetectors.Specific {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/specifics/%s)\n", detector, basepath, detector)
	}
}
