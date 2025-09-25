package usecase

import (
	"context"

	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/note/model"
	"HATCH_APP/internal/shared/customerr"
)

func (u UseCase) CreateNote(ctx context.Context, input dto.CreateNoteInput) (dto.CreateNoteOutput, error) {
	note := model.NewNote(input.Title, input.Content)

	if err := u.noteRepo.Create(ctx, note); err != nil {
		return dto.CreateNoteOutput{}, customerr.NewDatabaseError(err)
	}

	return dto.CreateNoteOutput{
		ID: note.ID,
	}, nil
}
