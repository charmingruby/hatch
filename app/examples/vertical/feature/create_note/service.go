package create_note

import (
	"HATCH_APP/examples/vertical/domain"
	"HATCH_APP/internal/shared/errs"
	"context"
)

type UseCase interface {
	Execute(ctx context.Context, title, content string) (string, error)
}

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
		return "", errs.NewDatabaseError(err)
	}

	return note.ID, nil
}
