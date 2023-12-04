package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Name = string

//go:embed data/names.yaml
var nameData []byte

func LoadNames() (names []Name, err error) {
	if err = yaml.Unmarshal(nameData, &names); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
