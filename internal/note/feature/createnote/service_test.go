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

type suite struct {
	repo    *mocks.NoteRepository
	service *createnote.Service
}

func setupSuite(t *testing.T) *suite {
	repo := mocks.NewNoteRepository(t)

	service := createnote.NewService(repo)

	return &suite{
		repo:    repo,
		service: service,
	}
}

func Test_Service_Execute(t *testing.T) {
	title := "Hatch"
	content := "Template"

	tests := []struct {
		arrange func(t *testing.T, s *suite)
		assert  func(t *testing.T, id string, err error)
		name    string
	}{
		{
			name: "should create successfully",
			arrange: func(t *testing.T, s *suite) {
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
			arrange: func(t *testing.T, s *suite) {
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
			s := setupSuite(t)

			if tc.arrange != nil {
				tc.arrange(t, s)
			}

			id, err := s.service.Execute(t.Context(), title, content)

			tc.assert(t, id, err)
		})
	}
}
