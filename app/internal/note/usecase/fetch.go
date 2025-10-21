package usecase

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/shared/errs"
	"context"
)

func (s *Service) Fetch(ctx context.Context) ([]domain.Note, error) {
	notes, err := s.repo.List(ctx)
	if err != nil {
		return nil, errs.NewDatabaseError(err)
	}

	return notes, nil
}
