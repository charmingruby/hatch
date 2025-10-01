package usecase

import (
	"context"

	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/note/repository"
)

type Service interface {
	CreateNote(ctx context.Context, input dto.CreateNoteInput) (dto.CreateNoteOutput, error)
	ListNotes(ctx context.Context) (dto.ListNotesOutput, error)
	ArchiveNote(ctx context.Context, input dto.ArchiveNoteInput) error
}

type UseCase struct {
	noteRepo repository.NoteRepo
}

func New(noteRepo repository.NoteRepo) Service {
	return UseCase{
		noteRepo: noteRepo,
	}
}
