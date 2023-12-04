package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Virtue = string

//go:embed data/virtues.yaml
var virtueData []byte

func LoadVirtues() (virtues []Virtue, err error) {
	if err = yaml.Unmarshal(virtueData, &virtues); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
