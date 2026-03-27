package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/core/apperr"
	"context"
)

type Service struct {
	noteRepo domain.NoteRepository
}

func NewService(noteRepo domain.NoteRepository) *Service {
	return &Service{
		noteRepo: noteRepo,
	}
}

func (s *Service) ListNotes(ctx context.Context) ([]*domain.Note, error) {
	notes, err := s.noteRepo.List(ctx)
	if err != nil {
		return nil, apperr.Internal("failed to list notes", err)
	}

	return notes, nil
}
