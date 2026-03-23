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

type serviceSuite struct {
	repo    *mocks.NoteRepository
	service *archivenote.Service
}

func setupSuite(t *testing.T) *serviceSuite {
	repo := mocks.NewNoteRepository(t)

	service := archivenote.NewService(repo)

	return &serviceSuite{
		repo:    repo,
		service: service,
	}
}

func TestServiceArchiveNote(t *testing.T) {
	tests := []struct {
		arrange   func(t *testing.T, s *serviceSuite) string
		assertErr func(t *testing.T, err error)
		name      string
	}{
		{
			name: "should archive successfully",
			arrange: func(t *testing.T, s *serviceSuite) string {
				n := domain.NewNote("title", "content")

				s.repo.On("FindByID", t.Context(), n.ID).
					Return(n, nil).
					Once()

				s.repo.On("Save", t.Context(), mock.MatchedBy(func(note *domain.Note) bool {
					return note.ID == n.ID &&
						note.Archived &&
						note.UpdatedAt != nil &&
						time.Since(*note.UpdatedAt) < time.Second
				})).
					Return(nil).
					Once()

				return n.ID
			},
			assertErr: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "should return error when FindByID fails",
			arrange: func(t *testing.T, s *serviceSuite) string {
				s.repo.On("FindByID", t.Context(), "nonexistent").
					Return((*domain.Note)(nil), errors.New("repo down")).
					Once()

				return "nonexistent"
			},
			assertErr: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "should return error when Save fails",
			arrange: func(t *testing.T, s *serviceSuite) string {
				n := domain.NewNote("title", "content")

				s.repo.On("FindByID", t.Context(), n.ID).
					Return(n, nil).
					Once()

				s.repo.On("Save", t.Context(), mock.Anything).
					Return(errors.New("save error")).
					Once()

				return n.ID
			},
			assertErr: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "should return ErrNoteNotFound when note ID is empty",
			arrange: func(t *testing.T, s *serviceSuite) string {
				s.repo.On("FindByID", mock.Anything, "invalid-id").
					Return((*domain.Note)(nil), domain.ErrNoteNotFound).
					Once()

				return "invalid-id"
			},
			assertErr: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.ErrorIs(t, err, domain.ErrNoteNotFound)
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			s := setupSuite(t)

			noteID := tc.arrange(t, s)

			err := s.service.ArchiveNote(t.Context(), noteID)

			tc.assertErr(t, err)
		})
	}
}
