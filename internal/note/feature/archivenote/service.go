package archivenote

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/core/apperr"
	"context"
)

type Service struct {
	noteRepo domain.NoteRepository
}

func NewService(noteRepo domain.NoteRepository) *Service {
	return &Service{
		noteRepo: noteRepo,
	}
}

func (s *Service) ArchiveNote(ctx context.Context, id string) error {
	note, err := s.noteRepo.FindByID(ctx, id)

	if err != nil {
		return apperr.Internal("failed to find note", err)
	}

	if note == nil {
		return apperr.NotFound("note not found")
	}

	note.Archive()

	if err := s.noteRepo.Save(ctx, note); err != nil {
		return apperr.Internal("failed to save note", err)
	}

	return nil
}
