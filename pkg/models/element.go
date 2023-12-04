package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Element struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

//go:embed data/tag_symbols.yaml
var elementData []byte

func LoadElements() (symbols []Element, err error) {
	if err = yaml.Unmarshal(elementData, &symbols); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
