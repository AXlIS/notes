package service

import (
	"github.com/AXlIS/notes/internal/model"
	"github.com/AXlIS/notes/pkg/speller"
	"strings"
)

type CorrectorService struct {
	Client *speller.SpellClient
}

func NewCorrectorService() *CorrectorService {
	return &CorrectorService{
		Client: speller.NewSpellClient(),
	}
}

func (s *CorrectorService) ValidateText(note *model.Note) error {
	spellResponse, err := s.Client.CheckText(note.Text)
	if err != nil {
		return err
	}

	for _, spell := range spellResponse {
		if len(spell.Suggestions) > 0 {
			note.Text = strings.Replace(note.Text, spell.Word, spell.Suggestions[0], -1)
		}
	}

	return nil
}
