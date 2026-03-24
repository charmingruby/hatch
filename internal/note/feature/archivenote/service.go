package archivenote

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

func (s *Service) ArchiveNote(ctx context.Context, id string) error {
	note, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return apperr.Internal("failed to find note", err)
	}

	if note == nil {
		return apperr.NotFound("note not found")
	}

	note.Archive()

	if err := s.repo.Save(ctx, note); err != nil {
		return apperr.Internal("failed to save note", err)
	}

	return nil
}
