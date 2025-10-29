package core_test

import (
	"HATCH_APP/internal/note/core"
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
	usecase core.UseCase
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := core.NewService(repo)

	return &suite{
		repo:    repo,
		usecase: service,
	}
}

func Test_UseCase_ArchiveNote(t *testing.T) {
	t.Run("should archive successfully", func(t *testing.T) {
		s := setupSuite(t)

		n := core.NewNote("title", "content")

		s.repo.On("FindByID", t.Context(), n.ID).
			Return(n, nil).
			Once()

		s.repo.On("Save", t.Context(), mock.MatchedBy(func(note core.Note) bool {
			return note.ID == n.ID &&
				note.Archived &&
				note.UpdatedAt != nil &&
				time.Since(*note.UpdatedAt) < time.Second
		})).
			Return(nil).
			Once()

		err := s.usecase.ArchiveNote(t.Context(), n.ID)

		require.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindByID fails", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", t.Context(), "nonexistent").
			Return(core.Note{}, errors.New("repo down")).
			Once()

		err := s.usecase.ArchiveNote(t.Context(), "nonexistent")

		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when Save fails", func(t *testing.T) {
		s := setupSuite(t)

		n := core.NewNote("title", "content")

		s.repo.On("FindByID", t.Context(), n.ID).
			Return(n, nil).
			Once()

		s.repo.On("Save", t.Context(), mock.Anything).
			Return(errors.New("save error")).
			Once()

		err := s.usecase.ArchiveNote(t.Context(), n.ID)

		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when note ID is empty", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("FindByID", mock.Anything, "invalid-id").
			Return(core.Note{}, nil).
			Once()

		err := s.usecase.ArchiveNote(t.Context(), "invalid-id")

		require.Error(t, err)
		var notFoundErr *errs.NotFoundError
		assert.ErrorAs(t, err, &notFoundErr)
	})
}

func Test_UseCase_FetchNotes(t *testing.T) {
	t.Run("should list notes successfully", func(t *testing.T) {
		s := setupSuite(t)

		notes := []core.Note{
			core.NewNote("title1", "content1"),
			core.NewNote("title2", "content2"),
		}

		s.repo.On("List", t.Context()).
			Return(notes, nil).
			Once()

		notes, err := s.usecase.FetchNotes(t.Context())

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

		op, err := s.usecase.FetchNotes(t.Context())

		assert.Zero(t, op)
		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}

func Test_UseCase_CreateNote(t *testing.T) {
	title := "Hatch"
	content := "Template"

	t.Run("should create successfully", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("Create", t.Context(), mock.MatchedBy(func(n core.Note) bool {
			return n.Title == title &&
				n.Content == content
		})).
			Return(nil).
			Once()

		id, err := s.usecase.CreateNote(t.Context(), title, content)

		require.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("should return a DatabaseError when there is a datasource error", func(t *testing.T) {
		s := setupSuite(t)

		s.repo.On("Create", mock.Anything, mock.Anything).
			Return(errors.New("unhealthy repo")).
			Once()

		id, err := s.usecase.CreateNote(t.Context(), title, content)

		assert.Empty(t, id)
		require.Error(t, err)

		var targetErr *errs.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
