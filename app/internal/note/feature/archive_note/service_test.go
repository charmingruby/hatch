package archive_note_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/archive_note"
	"HATCH_APP/internal/note/mocks"
	"HATCH_APP/internal/shared/errs"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepository
	usecase archive_note.UseCase
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := archive_note.NewService(repo)

	return &suite{
		repo:    repo,
		usecase: service,
	}
}

func Test_UseCase_Execute(t *testing.T) {
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

		err := s.usecase.Execute(t.Context(), n.ID)

		require.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindByID fails", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", t.Context(), "nonexistent").
			Return(domain.Note{}, errors.New("repo down")).
			Once()

		err := s.usecase.Execute(t.Context(), "nonexistent")

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

		err := s.usecase.Execute(t.Context(), n.ID)

		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when note ID is empty", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", mock.Anything, "invalid-id").
			Return(domain.Note{}, nil).
			Once()

		err := s.usecase.Execute(t.Context(), "invalid-id")

		require.Error(t, err)
		var notFoundErr *errs.NotFoundError
		assert.ErrorAs(t, err, &notFoundErr)
	})
}
