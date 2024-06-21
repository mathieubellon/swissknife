package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
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

	fmt.Println("Added Detectors:")
	for _, detector := range release.AddedDetectors.Generic {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/generics/%s)\n", detector, basepath, detector)
	}
	for _, detector := range release.AddedDetectors.Specific {
		markdownURL, err := getURL("specifics", detector)
		if err != nil {
			fmt.Println("Error getting title:", err)
			return
		}
		fmt.Printf("%s", markdownURL)
	}

	fmt.Println("Modified Detectors:")
	for _, detector := range release.ModifiedDetectors.Generic {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/generics/%s)\n", detector, basepath, detector)
	}
	for _, detector := range release.ModifiedDetectors.Specific {
		markdownURL, err := getURL("specifics", detector)
		if err != nil {
			fmt.Println("Error getting title:", err)
			return
		}
		fmt.Printf("%s", markdownURL)
	}

	fmt.Println("Removed Detectors:")
	for _, detector := range release.RemovedDetectors.Generic {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/generics/%s)\n", detector, basepath, detector)
	}
	for _, detector := range release.RemovedDetectors.Specific {
		fmt.Printf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/specifics/%s)\n", detector, basepath, detector)
	}
}

func getURL(category string, detector any) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/secrets-detection/secrets-detection-engine/detectors/%s/%s", basepath, category, detector))
	if err != nil {
		return "Error during get", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "Error during lookup", err
	}
	title := doc.Find("#__docusaurus_skipToContent_fallback > div > main > div > div > div > div > article > div.theme-doc-markdown.markdown > div > div.section_AGm0.is--w-border-btm_hYla.is--rte-section > div.rich-text-block_RQVQ.w-richtext > h1").Text()
	markdown_url := fmt.Sprintf("[%s](%s/secrets-detection/secrets-detection-engine/detectors/%s/%s)\n", title, basepath, category, detector)
	return markdown_url, nil
}
