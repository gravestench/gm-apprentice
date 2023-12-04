package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type SensorySnippet struct {
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

//go:embed data/sensory_snippets.yaml
var sensorySnippetData []byte

func LoadSensorySnippets() (sensorySnippets []SensorySnippet, err error) {
	if err = yaml.Unmarshal(sensorySnippetData, &sensorySnippets); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
