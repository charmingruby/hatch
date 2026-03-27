package createnote

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

func (s *Service) CreateNote(ctx context.Context, title, content string) (string, error) {
	note := domain.NewNote(title, content)

	if err := s.noteRepo.Create(ctx, note); err != nil {
		return "", apperr.Internal("failed to create note", err)
	}

	return note.ID, nil
}
