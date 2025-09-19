package usecase

import (
	"context"

	"PACK_APP/internal/note/repository"
)

type Service interface {
	CreateNote(ctx context.Context, input CreateNoteInput) (CreateNoteOutput, error)
	ListNotes(ctx context.Context) (ListNotesOutput, error)
	ArchiveNote(ctx context.Context, input ArchiveNoteInput) error
}

type UseCase struct {
	noteRepo repository.NoteRepository
}

func New(noteRepo repository.NoteRepository) Service {
	return UseCase{
		noteRepo: noteRepo,
	}
}
