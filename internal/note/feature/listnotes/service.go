package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/core/apperr"
	"context"
)

type Service struct {
	repo domain.NoteRepository
}

func NewService(repo domain.NoteRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ListNotes(ctx context.Context) ([]*domain.Note, error) {
	notes, err := s.repo.List(ctx)
	if err != nil {
		return nil, apperr.Internal("failed to list notes", err)
	}

	return notes, nil
}
