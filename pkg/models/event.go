package models

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Event = string

//go:embed data/events.yaml
var eventData []byte

func LoadEvents() (events []Event, err error) {
	if err = yaml.Unmarshal(eventData, &events); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return
}
