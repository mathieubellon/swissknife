package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

func getDetectorsList(repo string) (detectors []Detector, err error) {
	DetectorsList := []Detector{}
	StartDir := fmt.Sprintf("%s/pkg-tokenscanner/tokenscanner/config/detectors", repo)
	err = filepath.WalkDir(StartDir, func(path string, d os.DirEntry, err error) error {
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
				return fmt.Errorf("failed to unmarshal detector: %s, %s", path, err)
			}
			// TODO : This is a trap
			if detector.Nature == "" || detector.Nature == "specific" {
				detector.Nature = "specifics"
			} else {
				detector.Nature = "generics"
			}
			detector.URL = fmt.Sprintf("%s/secrets-detection/secrets-detection-engine/detectors/%s/%s", GitGuardianPublicDocBasePath, detector.Nature, detector.GroupName)
			detector.BrandName = detector.GroupName[:strings.Index(detector.GroupName, "_")]
			DetectorsList = append(DetectorsList, detector)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return DetectorsList, nil
}
