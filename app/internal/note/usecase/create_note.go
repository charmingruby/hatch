package usecase

import (
	"context"

	"github.com/charmingruby/pack/internal/note/model"
	"github.com/charmingruby/pack/internal/shared/customerr"
)

type CreateNoteInput struct {
	Title   string
	Content string
}

type CreateNoteOutput struct {
	ID string
}

func (u UseCase) CreateNote(ctx context.Context, input CreateNoteInput) (CreateNoteOutput, error) {
	note := model.NewNote(input.Title, input.Content)

	if err := u.noteRepo.Create(ctx, note); err != nil {
		return CreateNoteOutput{}, customerr.NewDatabaseError(err)
	}

	return CreateNoteOutput{
		ID: note.ID,
	}, nil
}
