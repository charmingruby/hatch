package createnote_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepository
	service *createnote.Service
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := createnote.NewService(repo)

	return &suite{
		repo:    repo,
		service: service,
	}
}

func Test_Service_Execute(t *testing.T) {
	title := "Hatch"
	content := "Template"

	t.Run("should create successfully", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("Create", t.Context(), mock.MatchedBy(func(n domain.Note) bool {
			return n.Title == title &&
				n.Content == content
		})).
			Return(nil).
			Once()

		id, err := s.service.Execute(t.Context(), title, content)

		require.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("should return error when there is a datasource error", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("Create", mock.Anything, mock.Anything).
			Return(errors.New("unhealthy repo")).
			Once()

		id, err := s.service.Execute(t.Context(), title, content)

		assert.Empty(t, id)
		require.Error(t, err)
	})
}
