package records

import (
	"github.com/rs/zerolog"

	"dm-apprentice/pkg/models"
)

type Service struct {
	logger *zerolog.Logger

	Belongings      []models.Belonging
	Catalysts       []models.Catalyst
	Elements        []models.Element
	Events          []models.Event
	Locations       []models.Location
	Names           []models.Name
	Runes           []models.Rune
	SensorySnippets []models.SensorySnippet
	TagSymbols      []models.TagSymbol
	Vices           []models.Vice
	Virtues         []models.Virtue
}

func (s *Service) Name() string {
	return "Records"
}
