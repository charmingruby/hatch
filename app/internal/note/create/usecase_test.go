package create_test

import (
	"errors"
	"testing"

	"HATCH_APP/internal/note/create"
	"HATCH_APP/internal/note/shared/model"
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/test/gen/note/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepo
	usecase create.UseCase
}

func setup(t *testing.T) suite {
	repo := mocks.NewNoteRepo(t)

	return suite{
		repo:    repo,
		usecase: create.NewUseCase(repo),
	}
}

func Test_Execute(t *testing.T) {
	title := "Hatch"
	content := "Template"

	t.Run("should create successfully", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Create", t.Context(), mock.MatchedBy(func(n model.Note) bool {
			return n.Title == title &&
				n.Content == content
		})).
			Return(nil).
			Once()

		op, err := s.usecase.Execute(t.Context(), create.UseCaseInput{
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

		op, err := s.usecase.Execute(t.Context(), create.UseCaseInput{
			Title:   title,
			Content: content,
		})

		assert.Zero(t, op)

		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
