package fetch_notes_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/fetch_notes"
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/test/gen/note/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepository
	usecase fetch_notes.UseCase
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := fetch_notes.NewService(repo)

	return &suite{
		repo:    repo,
		usecase: service,
	}
}

func Test_UseCase_Execute(t *testing.T) {
	t.Run("should list notes successfully", func(t *testing.T) {
		s := setupSuite(t)

		notes := []domain.Note{
			domain.NewNote("title1", "content1"),
			domain.NewNote("title2", "content2"),
		}

		s.repo.On("List", t.Context()).
			Return(notes, nil).
			Once()

		notes, err := s.usecase.Execute(t.Context())

		require.NoError(t, err)
		assert.Len(t, notes, 2)
		assert.Equal(t, "title1", notes[0].Title)
		assert.Equal(t, "title2", notes[1].Title)
	})

	t.Run("should return DatabaseError when List fails", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("List", t.Context()).
			Return(nil, errors.New("db error")).
			Once()

		op, err := s.usecase.Execute(t.Context())

		assert.Zero(t, op)
		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
