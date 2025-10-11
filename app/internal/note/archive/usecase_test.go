package archive_test

import (
	"errors"
	"testing"
	"time"

	"HATCH_APP/internal/note/archive"
	"HATCH_APP/internal/note/shared/model"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/test/gen/note/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepo
	usecase archive.UseCase
}

func setup(t *testing.T) suite {
	repo := mocks.NewNoteRepo(t)

	return suite{
		repo:    repo,
		usecase: archive.NewUseCase(repo),
	}
}

func Test_Execute(t *testing.T) {
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

		err := s.usecase.Execute(t.Context(), archive.Input{
			ID: n.ID,
		})

		require.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindByID fails", func(t *testing.T) {
		s := setup(t)

		s.repo.On("FindByID", t.Context(), "nonexistent").
			Return(model.Note{}, errors.New("repo down")).
			Once()

		err := s.usecase.Execute(t.Context(), archive.Input{
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

		err := s.usecase.Execute(t.Context(), archive.Input{
			ID: n.ID,
		})

		require.Error(t, err)

		var targetErr *customerr.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
