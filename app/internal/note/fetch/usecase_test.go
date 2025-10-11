package fetch_test

import (
	"errors"
	"testing"

	"HATCH_APP/internal/note/fetch"
	"HATCH_APP/internal/note/shared/model"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/test/gen/note/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepo
	usecase fetch.UseCase
}

func setup(t *testing.T) suite {
	repo := mocks.NewNoteRepo(t)

	return suite{
		repo:    repo,
		usecase: fetch.NewUseCase(repo),
	}
}

func Test_Execute(t *testing.T) {
	t.Run("should list notes successfully", func(t *testing.T) {
		s := setup(t)

		notes := []model.Note{
			model.NewNote("title1", "content1"),
			model.NewNote("title2", "content2"),
		}

		s.repo.On("List", t.Context()).
			Return(notes, nil).
			Once()

		op, err := s.usecase.Execute(t.Context())

		require.NoError(t, err)
		assert.Len(t, op.Notes, 2)
		assert.Equal(t, "title1", op.Notes[0].Title)
		assert.Equal(t, "title2", op.Notes[1].Title)
	})

	t.Run("should return DatabaseError when List fails", func(t *testing.T) {
		s := setup(t)

		s.repo.On("List", t.Context()).
			Return(nil, errors.New("db error")).
			Once()

		op, err := s.usecase.Execute(t.Context())

		assert.Zero(t, op)
		require.Error(t, err)

		var targetErr *customerr.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
