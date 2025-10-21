package fetch

import (
	"context"

	"HATCH_APP/internal/note/shared/model"
	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/internal/shared/errs"
)

type UseCaseOutput struct {
	Notes []model.Note
}

type UseCase struct {
	repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context) (UseCaseOutput, error) {
	notes, err := u.repo.List(ctx)

	if err != nil {
		return UseCaseOutput{}, errs.NewDatabaseError(err)
	}

	return UseCaseOutput{
		Notes: notes,
	}, nil
}
