package main

import "time"

type Detector struct {
	Type        string `yaml:"type" json:"type"`
	Nature      string `yaml:"nature,omitempty" json:"nature,omitempty"`
	Name        string `yaml:"name" json:"name"`
	Family      string `yaml:"family" json:"family"`
	GroupName   string `yaml:"group_name" json:"group_name"`
	DisplayName string `yaml:"display_name" json:"display_name"`
	URL         string `yaml:"url,omitempty" json:"url,omitempty"`
	BrandName   string `yaml:"brand_name,omitempty" json:"brand_name,omitempty"`
	//Postvalidators []string `yaml:"postvalidators,omitempty" json:"postvalidators,omitempty"`
}

// Metadata contains detection metadata configuration
type Metadata struct {
	Config                      MetadataConfig `yaml:"config"`
	Category                    string         `yaml:"category"`
	PercentageValid             float64        `yaml:"percentage_valid"`
	FrequencyEstimateAfterCheck float64        `yaml:"frequency_estimate_after_check"`
	FrequencyEstimate           float64        `yaml:"frequency_estimate"`
	Company                     string         `yaml:"company"`
	Provider                    string         `yaml:"provider"`
}

// MetadataConfig contains required check settings
type MetadataConfig struct {
	RequiredCheck bool `yaml:"required_check"`
}

// PreValidator represents a validation step performed before detection
type PreValidator struct {
	Type                            string   `yaml:"type"`
	IncludeDefaultBanlistExtensions bool     `yaml:"include_default_banlist_extensions,omitempty"`
	BanlistExtensions               []string `yaml:"banlist_extensions,omitempty"`
	BanlistFilenames                []string `yaml:"banlist_filenames,omitempty"`
	Patterns                        []string `yaml:"patterns,omitempty"`
}

// Matcher represents a detection rule
type Matcher struct {
	Type       string `yaml:"type"`
	LookBehind string `yaml:"look_behind,omitempty"`
	Pattern    string `yaml:"pattern"`
	LookAhead  string `yaml:"look_ahead,omitempty"`
}

// PostValidator represents a validation step performed after detection
type PostValidator struct {
	Type        string   `yaml:"type"`
	Digits      int      `yaml:"digits,omitempty"`
	Entropy     float64  `yaml:"entropy,omitempty"`
	WindowWidth int      `yaml:"window_width,omitempty"`
	WindowType  string   `yaml:"window_type,omitempty"`
	Patterns    []string `yaml:"patterns,omitempty"`
}

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
