package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Catalyst = string

//go:embed data/catalysts.yaml
var catalystData []byte

func LoadCatalysts() (catalysts []Catalyst, err error) {
	if err = yaml.Unmarshal(catalystData, &catalysts); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
