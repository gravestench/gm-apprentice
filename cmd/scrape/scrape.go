package main

import (
	_ "embed"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"dm-apprentice/pkg/models"
)

type T struct {
	Event1       []string `yaml:"Event1"`
	Event2       []string `yaml:"Event2"`
	Event3       []string `yaml:"Event3"`
	SensoryHear  []string `yaml:"SensoryHear"`
	SensorySee   []string `yaml:"SensorySee"`
	SensoryTouch []string `yaml:"SensoryTouch"`
	SensorySmell []string `yaml:"SensorySmell"`
	Belongings   []string `yaml:"Belongings"`
	Name1        []string `yaml:"Name1"`
	Name2        []string `yaml:"Name2"`
	Name3        []string `yaml:"Name3"`
	Catalyst     []string `yaml:"Catalyst"`
	Virtue       []string `yaml:"Virtue"`
	Location     []string `yaml:"Location"`
	Vice         []string `yaml:"Vice"`
}

//go:embed scraped.yaml
var data []byte

func main() {
	var raw T
	data := data

	if err := yaml.Unmarshal(data, &raw); err != nil {
		panic(err)
	}

	{
		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(raw.Belongings)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("belongings.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	{
		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(raw.Catalyst)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("catalysts.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	{
		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(raw.Virtue)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("virtues.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	{
		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(raw.Location)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("locations.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	{
		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(raw.Vice)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("vices.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	{
		list := raw.Event1
		list = append(list, raw.Event2...)
		list = append(list, raw.Event3...)

		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(list)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("events.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	{
		list := raw.Name1
		list = append(list, raw.Name2...)
		list = append(list, raw.Name3...)

		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(list)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("names.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	{
		sensorySnippets := make([]models.SensorySnippet, 0)

		for _, s := range raw.SensorySmell {
			sensorySnippets = append(sensorySnippets, models.SensorySnippet{
				Type:        "Olfactory",
				Description: s,
			})
		}

		for _, s := range raw.SensorySee {
			sensorySnippets = append(sensorySnippets, models.SensorySnippet{
				Type:        "Visual",
				Description: s,
			})
		}

		for _, s := range raw.SensoryHear {
			sensorySnippets = append(sensorySnippets, models.SensorySnippet{
				Type:        "Auditory",
				Description: s,
			})
		}

		for _, s := range raw.SensoryTouch {
			sensorySnippets = append(sensorySnippets, models.SensorySnippet{
				Type:        "Tactile",
				Description: s,
			})
		}

		// Marshal the struct into YAML format
		yamlData, err := yaml.Marshal(sensorySnippets)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}

		// Write the YAML data to a file
		err = ioutil.WriteFile("sensory_snippets.yaml", yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}
}
