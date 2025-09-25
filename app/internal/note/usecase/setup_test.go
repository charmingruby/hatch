package usecase_test

import (
	"testing"

	"HATCH_APP/internal/note/usecase"
	"HATCH_APP/test/gen/note/mocks"
)

type suite struct {
	repo    *mocks.NoteRepository
	usecase usecase.Service
}

func setup(t *testing.T) suite {
	repo := mocks.NewNoteRepository(t)

	return suite{
		repo:    repo,
		usecase: usecase.New(repo),
	}
}
