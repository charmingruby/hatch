package createnote

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

func (s *Service) CreateNote(ctx context.Context, title, content string) (string, error) {
	note := domain.NewNote(title, content)

	if err := s.repo.Create(ctx, note); err != nil {
		return "", apperr.Internal("failed to create note", err)
	}

	return note.ID, nil
}
