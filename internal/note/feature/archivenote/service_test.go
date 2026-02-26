package archivenote_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/archivenote"
	"HATCH_APP/internal/note/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepository
	service *archivenote.Service
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := archivenote.NewService(repo)

	return &suite{
		repo:    repo,
		service: service,
	}
}

func Test_Service_Execute(t *testing.T) {
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

		err := s.service.Execute(t.Context(), n.ID)

		require.NoError(t, err)
	})

	t.Run("should return error when FindByID fails", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", t.Context(), "nonexistent").
			Return(domain.Note{}, errors.New("repo down")).
			Once()

		err := s.service.Execute(t.Context(), "nonexistent")

		require.Error(t, err)
	})

	t.Run("should return error when Save fails", func(t *testing.T) {
		s := setupSuite(t)

		n := domain.NewNote("title", "content")

		s.repo.On("FindByID", t.Context(), n.ID).
			Return(n, nil).
			Once()

		s.repo.On("Save", t.Context(), mock.Anything).
			Return(errors.New("save error")).
			Once()

		err := s.service.Execute(t.Context(), n.ID)

		require.Error(t, err)
	})

	t.Run("should return ErrNoteNotFound when note ID is empty", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", mock.Anything, "invalid-id").
			Return(domain.Note{}, nil).
			Once()

		err := s.service.Execute(t.Context(), "invalid-id")

		require.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrNoteNotFound)
	})
}
