package usecase

import (
	"context"

	"github.com/charmingruby/pack/internal/note/model"
	"github.com/charmingruby/pack/internal/shared/customerr"
)

type ListNotesOutput struct {
	Notes []model.Note
}

func (u UseCase) ListNotes(ctx context.Context) (ListNotesOutput, error) {
	notes, err := u.noteRepo.List(ctx)

	if err != nil {
		return ListNotesOutput{}, customerr.NewDatabaseError(err)
	}

	return ListNotesOutput{
		Notes: notes,
	}, nil
}
