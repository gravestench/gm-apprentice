package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Rune struct {
	Name          string `yaml:"name"`
	Symbol        string `yaml:"symbol"`
	Meaning       string `yaml:"meaning"`
	InGameMeaning string `yaml:"in_game_meaning"`
}

//go:embed data/runes.yaml
var runeData []byte

func LoadRunes() (runes []Rune, err error) {
	if err = yaml.Unmarshal(runeData, &runes); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
