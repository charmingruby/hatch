package create

import (
	"context"

	"HATCH_APP/internal/note/shared/model"
	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/internal/shared/customerr"
)

type UseCase struct {
	repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	note := model.NewNote(input.Title, input.Content)

	if err := u.repo.Create(ctx, note); err != nil {
		return Output{}, customerr.NewDatabaseError(err)
	}

	return Output{
		ID: note.ID,
	}, nil
}
