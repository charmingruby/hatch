package archivenote

import (
	"HATCH_APP/internal/common/errs"
	"HATCH_APP/internal/note/domain"
	"context"
	"time"
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
