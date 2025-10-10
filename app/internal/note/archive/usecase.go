package archive

import (
	"context"
	"time"

	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/internal/shared/customerr"
)

type Service interface {
	Execute(ctx context.Context, input Input) error
}

type UseCase struct {
	repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context, input Input) error {
	note, err := u.repo.FindByID(ctx, input.ID)

	if err != nil {
		return customerr.NewDatabaseError(err)
	}

	if note.ID == "" {
		return customerr.NewNotFoundError("note")
	}

	now := time.Now()
	note.Archived = true
	note.UpdatedAt = &now

	if err := u.repo.Save(ctx, note); err != nil {
		return customerr.NewDatabaseError(err)
	}

	return nil
}
