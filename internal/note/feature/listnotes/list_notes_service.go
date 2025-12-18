package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"context"
	"fmt"
)

type Service struct {
	repo domain.NoteRepository
}

func NewService(repo domain.NoteRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Execute(ctx context.Context) ([]domain.Note, error) {
	notes, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list notes: %w", err)
	}

	return notes, nil
}
