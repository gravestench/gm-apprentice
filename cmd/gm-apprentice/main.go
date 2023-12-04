package main

import (
	"fmt"
	"log"

	"dm-apprentice/pkg/models"
)

func main() {

	snippets, err := models.LoadSensorySnippets()
	if err != nil {
		log.Fatalf("loading sensory snippets: %v", err)
	}

	for _, r := range snippets {
		fmt.Printf("(%s) %s\n", r.Type, r.Description)
	}
}
