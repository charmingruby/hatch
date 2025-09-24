package usecase

import (
	"context"

	"PACK_APP/internal/note/dto"
	"PACK_APP/internal/note/repository"
)

type Service interface {
	CreateNote(ctx context.Context, input dto.CreateNoteInput) (dto.CreateNoteOutput, error)
	ListNotes(ctx context.Context) (dto.ListNotesOutput, error)
	ArchiveNote(ctx context.Context, input dto.ArchiveNoteInput) error
}

type UseCase struct {
	noteRepo repository.NoteRepository
}

func New(noteRepo repository.NoteRepository) Service {
	return UseCase{
		noteRepo: noteRepo,
	}
}
