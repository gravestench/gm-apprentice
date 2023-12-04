package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Vice = string

//go:embed data/vices.yaml
var viceData []byte

func LoadVices() (vices []Vice, err error) {
	if err = yaml.Unmarshal(viceData, &vices); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
