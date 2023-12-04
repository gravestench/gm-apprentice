package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Belonging = string

//go:embed data/belongings.yaml
var belongingData []byte

func LoadBelongings() (belongings []Belonging, err error) {
	if err = yaml.Unmarshal(belongingData, &belongings); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
