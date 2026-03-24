package listnotes_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/listnotes"
	"HATCH_APP/internal/note/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type suite struct {
	repo    *mocks.NoteRepository
	service *listnotes.Service
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := listnotes.NewService(repo)

	return &suite{
		repo:    repo,
		service: service,
	}
}

func TestServiceListNotes(t *testing.T) {
	tests := []struct {
		arrange func(t *testing.T, s *suite)
		assert  func(t *testing.T, notes []*domain.Note, err error)
		name    string
	}{
		{
			name: "should list notes successfully",
			arrange: func(t *testing.T, s *suite) {
				ns := []*domain.Note{
					domain.NewNote("title1", "content1"),
					domain.NewNote("title2", "content2"),
				}

				s.repo.On("List", t.Context()).
					Return(ns, nil).
					Once()
			},
			assert: func(t *testing.T, notes []*domain.Note, err error) {
				require.NoError(t, err)
				assert.Len(t, notes, 2)
				assert.Equal(t, "title1", notes[0].Title)
				assert.Equal(t, "title2", notes[1].Title)
			},
		},
		{
			name: "should return error when List fails",
			arrange: func(t *testing.T, s *suite) {
				s.repo.On("List", t.Context()).
					Return(nil, errors.New("db error")).
					Once()
			},
			assert: func(t *testing.T, notes []*domain.Note, err error) {
				assert.Zero(t, notes)
				require.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			s := setupSuite(t)

			if tc.arrange != nil {
				tc.arrange(t, s)
			}

			notes, err := s.service.ListNotes(t.Context())

			tc.assert(t, notes, err)
		})
	}
}
