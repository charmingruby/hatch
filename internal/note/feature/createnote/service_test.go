package createnote_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type serviceSuite struct {
	repo    *mocks.NoteRepository
	service *createnote.Service
}

func setupServiceSuite(t *testing.T) *serviceSuite {
	repo := mocks.NewNoteRepository(t)

	service := createnote.NewService(repo)

	return &serviceSuite{
		repo:    repo,
		service: service,
	}
}

func TestServiceCreateNote(t *testing.T) {
	title := "Hatch"
	content := "Template"

	tests := []struct {
		arrange func(t *testing.T, s *serviceSuite)
		assert  func(t *testing.T, id string, err error)
		name    string
	}{
		{
			name: "should create successfully",
			arrange: func(t *testing.T, s *serviceSuite) {
				s.repo.On("Create", t.Context(), mock.MatchedBy(func(n *domain.Note) bool {
					return n.Title == title &&
						n.Content == content
				})).
					Return(nil).
					Once()
			},
			assert: func(t *testing.T, id string, err error) {
				require.NoError(t, err)
				assert.NotEmpty(t, id)
			},
		},
		{
			name: "should return error when there is a datasource error",
			arrange: func(t *testing.T, s *serviceSuite) {
				s.repo.On("Create", mock.Anything, mock.Anything).
					Return(errors.New("unhealthy repo")).
					Once()
			},
			assert: func(t *testing.T, id string, err error) {
				assert.Empty(t, id)
				require.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			s := setupServiceSuite(t)

			if tc.arrange != nil {
				tc.arrange(t, s)
			}

			id, err := s.service.CreateNote(t.Context(), title, content)

			tc.assert(t, id, err)
		})
	}
}
