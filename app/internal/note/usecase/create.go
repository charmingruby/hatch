package usecase

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/shared/errs"
	"context"
)

func (s *Service) Create(ctx context.Context, title, content string) (string, error) {
	note := domain.NewNote(title, content)

	if err := s.repo.Create(ctx, note); err != nil {
		return "", errs.NewDatabaseError(err)
	}

	return note.ID, nil
}
