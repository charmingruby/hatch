package usecase_test

import (
	"HATCH_APP/internal/note/usecase"
	"HATCH_APP/test/gen/note/mocks"
	"testing"
)

type suite struct {
	repo    *mocks.NoteRepository
	service *usecase.Service
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := usecase.NewService(repo)

	return &suite{
		repo:    repo,
		service: service,
	}
}
