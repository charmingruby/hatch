package create_note_test

import (
	"HATCH_APP/examples/vertical/domain"
	"HATCH_APP/examples/vertical/feature/create_note"
	"HATCH_APP/examples/vertical/mocks"
	"HATCH_APP/internal/shared/errs"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepository
	usecase create_note.UseCase
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := create_note.NewService(repo)

	return &suite{
		repo:    repo,
		usecase: service,
	}
}

func Test_UseCase_Execute(t *testing.T) {
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

		id, err := s.usecase.Execute(t.Context(), title, content)

		require.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("should return a DatabaseError when there is a datasource error", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("Create", mock.Anything, mock.Anything).
			Return(errors.New("unhealthy repo")).
			Once()

		id, err := s.usecase.Execute(t.Context(), title, content)

		assert.Empty(t, id)
		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
