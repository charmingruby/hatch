package usecase

import (
	"context"
	"time"

	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/shared/customerr"
)

func (u UseCase) ArchiveNote(ctx context.Context, input dto.ArchiveNoteInput) error {
	note, err := u.noteRepo.FindByID(ctx, input.ID)

	if err != nil {
		return customerr.NewDatabaseError(err)
	}

	now := time.Now()
	note.Archived = true
	note.UpdatedAt = &now

	if err := u.noteRepo.Save(ctx, note); err != nil {
		return customerr.NewDatabaseError(err)
	}

	return nil
}
