package usecase_test

import (
	"testing"

	"github.com/charmingruby/pack/internal/note/usecase"
	"github.com/charmingruby/pack/test/gen/note/mocks"
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
