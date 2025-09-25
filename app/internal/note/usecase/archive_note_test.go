package usecase_test

import (
	"errors"
	"testing"
	"time"

	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/note/model"
	"HATCH_APP/internal/shared/customerr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ArchiveNote(t *testing.T) {
	t.Run("should archive successfully", func(t *testing.T) {
		s := setup(t)
		n := model.NewNote("title", "content")

		s.repo.On("FindByID", t.Context(), n.ID).
			Return(n, nil).
			Once()

		s.repo.On("Save", t.Context(), mock.MatchedBy(func(note model.Note) bool {
			return note.ID == n.ID &&
				note.Archived &&
				note.UpdatedAt != nil &&
				time.Since(*note.UpdatedAt) < time.Second
		})).
			Return(nil).
			Once()

		err := s.usecase.ArchiveNote(t.Context(), dto.ArchiveNoteInput{
			ID: n.ID,
		})

		require.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindByID fails", func(t *testing.T) {
		s := setup(t)

		s.repo.On("FindByID", t.Context(), "nonexistent").
			Return(model.Note{}, errors.New("repo down")).
			Once()

		err := s.usecase.ArchiveNote(t.Context(), dto.ArchiveNoteInput{
			ID: "nonexistent",
		})

		require.Error(t, err)

		var targetErr *customerr.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when Save fails", func(t *testing.T) {
		s := setup(t)
		n := model.NewNote("title", "content")

		s.repo.On("FindByID", t.Context(), n.ID).
			Return(n, nil).
			Once()

		s.repo.On("Save", t.Context(), mock.Anything).
			Return(errors.New("save error")).
			Once()

		err := s.usecase.ArchiveNote(t.Context(), dto.ArchiveNoteInput{
			ID: n.ID,
		})

		require.Error(t, err)

		var targetErr *customerr.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
