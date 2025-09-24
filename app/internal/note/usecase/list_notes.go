package usecase

import (
	"context"

	"PACK_APP/internal/note/dto"
	"PACK_APP/internal/shared/customerr"
)

func (u UseCase) ListNotes(ctx context.Context) (dto.ListNotesOutput, error) {
	notes, err := u.noteRepo.List(ctx)

	if err != nil {
		return dto.ListNotesOutput{}, customerr.NewDatabaseError(err)
	}

	return dto.ListNotesOutput{
		Notes: notes,
	}, nil
}
