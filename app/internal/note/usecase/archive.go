package usecase

import (
	"HATCH_APP/internal/shared/errs"
	"context"
	"time"
)

func (s *Service) Archive(ctx context.Context, id string) error {
	note, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return errs.NewDatabaseError(err)
	}

	if note.ID == "" {
		return errs.NewNotFoundError("note")
	}

	now := time.Now()
	note.Archived = true
	note.UpdatedAt = &now

	if err := s.repo.Save(ctx, note); err != nil {
		return errs.NewDatabaseError(err)
	}

	return nil
}
