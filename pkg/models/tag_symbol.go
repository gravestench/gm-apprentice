package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type TagSymbol struct {
	Symbol      string `yaml:"symbol"`
	Description string `yaml:"description"`
}

//go:embed data/tag_symbols.yaml
var tagSymbolDataData []byte

func LoadTagSymbols() (symbols []TagSymbol, err error) {
	if err = yaml.Unmarshal(tagSymbolDataData, &symbols); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
