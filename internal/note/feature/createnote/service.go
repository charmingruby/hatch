package createnote

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

func (s *Service) Execute(ctx context.Context, title, content string) (string, error) {
	note := domain.NewNote(title, content)

	if err := s.repo.Create(ctx, note); err != nil {
		return "", fmt.Errorf("failed to create note: %w", err)
	}

	return note.ID, nil
}
