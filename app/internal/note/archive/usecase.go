package archive

import (
	"context"
	"time"

	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/internal/shared/errs"
)

type UseCaseInput struct {
	ID string `json:"id"`
}

type UseCase struct {
	repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context, input UseCaseInput) error {
	note, err := u.repo.FindByID(ctx, input.ID)

	if err != nil {
		return errs.NewDatabaseError(err)
	}

	if note.ID == "" {
		return errs.NewNotFoundError("note")
	}

	now := time.Now()
	note.Archived = true
	note.UpdatedAt = &now

	if err := u.repo.Save(ctx, note); err != nil {
		return errs.NewDatabaseError(err)
	}

	return nil
}
