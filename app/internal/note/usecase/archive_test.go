package usecase_test

import (
	"errors"
	"testing"
	"time"

	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/shared/errs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_UseCase_Archive(t *testing.T) {
	t.Run("should archive successfully", func(t *testing.T) {
		s := setupSuite(t)

		n := domain.NewNote("title", "content")

		s.repo.On("FindByID", t.Context(), n.ID).
			Return(n, nil).
			Once()

		s.repo.On("Save", t.Context(), mock.MatchedBy(func(note domain.Note) bool {
			return note.ID == n.ID &&
				note.Archived &&
				note.UpdatedAt != nil &&
				time.Since(*note.UpdatedAt) < time.Second
		})).
			Return(nil).
			Once()

		err := s.service.Archive(t.Context(), n.ID)

		require.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindByID fails", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", t.Context(), "nonexistent").
			Return(domain.Note{}, errors.New("repo down")).
			Once()

		err := s.service.Archive(t.Context(), "nonexistent")

		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when Save fails", func(t *testing.T) {
		s := setupSuite(t)

		n := domain.NewNote("title", "content")

		s.repo.On("FindByID", t.Context(), n.ID).
			Return(n, nil).
			Once()

		s.repo.On("Save", t.Context(), mock.Anything).
			Return(errors.New("save error")).
			Once()

		err := s.service.Archive(t.Context(), n.ID)

		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when note ID is empty", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", mock.Anything, "invalid-id").
			Return(domain.Note{}, nil).
			Once()

		err := s.service.Archive(t.Context(), "invalid-id")

		require.Error(t, err)
		var notFoundErr *errs.NotFoundError
		assert.ErrorAs(t, err, &notFoundErr)
	})
}
