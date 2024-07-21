package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type Detector struct {
	Type        string `yaml:"type" json:"type"`
	Nature      string `yaml:"nature,omitempty" json:"nature,omitempty"`
	Name        string `yaml:"name" json:"name"`
	Family      string `yaml:"family" json:"family"`
	GroupName   string `yaml:"group_name" json:"group_name"`
	DisplayName string `yaml:"display_name" json:"display_name"`
	URL         string `json:"url"`
	Metadata    struct {
		ShortTail                   bool    `yaml:"short_tail,omitempty" json:"short_tail,omitempty"`
		Category                    string  `yaml:"category,omitempty" json:"category,omitempty"`
		FrequencyEstimate           float64 `yaml:"frequency_estimate,omitempty" json:"frequency_estimate,omitempty"`
		PercentageValid             float64 `yaml:"percentage_valid,omitempty" json:"percentage_valid,omitempty"`
		FrequencyEstimateAfterCheck float64 `yaml:"frequency_estimate_after_check,omitempty" json:"frequency_estimate_after_check,omitempty"`
		SupportsOnPrem              bool    `yaml:"supports_on_prem,omitempty" json:"supports_on_prem,omitempty"`
		Config                      struct {
			RequiredCheck bool `yaml:"required_check,omitempty" json:"required_check,omitempty"`
		} `yaml:"config,omitempty" json:"config,omitempty"`
	} `yaml:"metadata,omitempty" json:"metadata,omitempty"`
}

func printDetectorsList(Ctx *cli.Context) error {
	repo := Ctx.String("repo")
	writeOneJson := Ctx.Bool("write-one-json")
	writeMultipleJson := Ctx.Bool("write-multiple-json")
	basepath := GitGuardianBasePath

	DetectorsList := []Detector{}
	StartDir := fmt.Sprintf("%s/tokenscanner/config/detectors", repo)
	err := filepath.WalkDir(StartDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var detector Detector
			err = yaml.Unmarshal(data, &detector)
			if err != nil {
				return err
			}
			// TODO : This is a trap
			if detector.Nature == "" || detector.Nature == "specific" {
				detector.Nature = "specifics"
			} else {
				detector.Nature = "generics"
			}
			detector.URL = fmt.Sprintf("%s/secrets-detection/secrets-detection-engine/detectors/%s/%s", basepath, detector.Nature, detector.GroupName)
			DetectorsList = append(DetectorsList, detector)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	detectorsJSON, err := json.Marshal(DetectorsList)
	if err != nil {
		return err
	}

	if writeOneJson {
		err = os.WriteFile("detectors.json", detectorsJSON, 0644)
		if err != nil {
			return err
		}
	} else if writeMultipleJson {
		for _, detector := range DetectorsList {
			detectorJSON, err := json.Marshal(detector)
			if err != nil {
				return err
			}
			err = os.MkdirAll("detectors", 0755)
			if err != nil {
				return err
			}
			err = os.WriteFile(fmt.Sprintf("detectors/%s.json", detector.Name), detectorJSON, 0644)
			if err != nil {
				return err
			}
		}
	} else {
		fmt.Println(string(detectorsJSON))
	}

	return nil
}
