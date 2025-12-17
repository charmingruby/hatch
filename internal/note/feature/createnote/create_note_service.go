package createnote

import (
	"HATCH_APP/internal/common/errs"
	"HATCH_APP/internal/note/domain"
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

func (s *Service) Execute(ctx context.Context, title, content string) (string, error) {
	note := domain.NewNote(title, content)

	if err := s.repo.Create(ctx, note); err != nil {
		return "", errs.NewDatabaseError(err)
	}

	return note.ID, nil
}
