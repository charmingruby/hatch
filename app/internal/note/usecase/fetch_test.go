package usecase_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/shared/errs"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UseCase_Fetch(t *testing.T) {
	t.Run("should list notes successfully", func(t *testing.T) {
		s := setupSuite(t)

		notes := []domain.Note{
			domain.NewNote("title1", "content1"),
			domain.NewNote("title2", "content2"),
		}

		s.repo.On("List", t.Context()).
			Return(notes, nil).
			Once()

		notes, err := s.service.Fetch(t.Context())

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

		op, err := s.service.Fetch(t.Context())

		assert.Zero(t, op)
		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
