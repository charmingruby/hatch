package core

import (
	"HATCH_APP/internal/shared/errs"
	"context"
	"time"
)

type UseCase interface {
	ArchiveNote(ctx context.Context, id string) error
	CreateNote(ctx context.Context, title, content string) (string, error)
	FetchNotes(ctx context.Context) ([]Note, error)
}

type Service struct {
	repo NoteRepository
}

func NewService(repo NoteRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ArchiveNote(ctx context.Context, id string) error {
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

func (s *Service) CreateNote(ctx context.Context, title, content string) (string, error) {
	note := NewNote(title, content)

	if err := s.repo.Create(ctx, note); err != nil {
		return "", errs.NewDatabaseError(err)
	}

	return note.ID, nil
}

func (s *Service) FetchNotes(ctx context.Context) ([]Note, error) {
	notes, err := s.repo.List(ctx)
	if err != nil {
		return nil, errs.NewDatabaseError(err)
	}

	return notes, nil
}
