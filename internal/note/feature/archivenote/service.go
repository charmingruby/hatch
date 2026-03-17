package archivenote

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

func (s *Service) Execute(ctx context.Context, id string) error {
	note, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return fmt.Errorf("failed to find note: %w", err)
	}

	if note == nil {
		return domain.ErrNoteNotFound
	}

	note.Archive()

	if err := s.repo.Save(ctx, note); err != nil {
		return fmt.Errorf("failed to save note: %w", err)
	}

	return nil
}
