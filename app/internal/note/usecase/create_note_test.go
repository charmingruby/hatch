package usecase_test

import (
	"errors"
	"testing"

	"PACK_APP/internal/note/model"
	"PACK_APP/internal/note/usecase"
	"PACK_APP/internal/shared/customerr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CreateNote(t *testing.T) {
	title := "Pack"
	content := "Template"

	t.Run("should create successfully", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Create", t.Context(), mock.MatchedBy(func(n model.Note) bool {
			return n.Title == title &&
				n.Content == content
		})).
			Return(nil).
			Once()

		op, err := s.usecase.CreateNote(t.Context(), usecase.CreateNoteInput{
			Title:   title,
			Content: content,
		})

		require.NoError(t, err)
		assert.NotEmpty(t, op.ID)
	})

	t.Run("should return a DatabaseError when there is a datasource error", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Create", t.Context(), mock.Anything).
			Return(errors.New("unhealthy repo")).
			Once()

		op, err := s.usecase.CreateNote(t.Context(), usecase.CreateNoteInput{
			Title:   title,
			Content: content,
		})

		assert.Zero(t, op)

		require.Error(t, err)

		var targetErr *customerr.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
