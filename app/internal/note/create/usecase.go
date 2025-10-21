package create

import (
	"context"

	"HATCH_APP/internal/note/shared/model"
	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/internal/shared/errs"
)

type UseCaseInput struct {
	Title   string
	Content string
}

type UseCaseOutput struct {
	ID string
}

type UseCase struct {
	repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context, input UseCaseInput) (UseCaseOutput, error) {
	note := model.NewNote(input.Title, input.Content)

	if err := u.repo.Create(ctx, note); err != nil {
		return UseCaseOutput{}, errs.NewDatabaseError(err)
	}

	return UseCaseOutput{
		ID: note.ID,
	}, nil
}
