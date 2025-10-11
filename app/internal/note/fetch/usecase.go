package fetch

import (
	"context"

	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/internal/shared/customerr"
)

type UseCase struct {
	repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context) (Output, error) {
	notes, err := u.repo.List(ctx)

	if err != nil {
		return Output{}, customerr.NewDatabaseError(err)
	}

	return Output{
		Notes: notes,
	}, nil
}
