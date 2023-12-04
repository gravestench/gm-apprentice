package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Location = string

//go:embed data/locations.yaml
var locationData []byte

func LoadLocations() (locations []Location, err error) {
	if err = yaml.Unmarshal(locationData, &locations); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
