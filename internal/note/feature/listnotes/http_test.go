package listnotes_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/listnotes"
	"HATCH_APP/internal/note/infra/database/postgres"
	"HATCH_APP/test/container"
	"HATCH_APP/test/httptest"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type httpSuite struct {
	repo *postgres.NoteRepository
	feat *listnotes.Feature
}

func setupHTTPSuite(t *testing.T) *httpSuite {
	db, dbTeardown := container.SetupPostgres(t)

	t.Cleanup(func() {
		dbTeardown()
	})

	repo, err := postgres.NewNoteRepository(db)
	require.NoError(t, err)

	return &httpSuite{
		repo: repo,
		feat: listnotes.New(repo),
	}
}

func TestHTTP(t *testing.T) {
	s := setupHTTPSuite(t)
	httptest.Init()

	tests := []struct {
		tc   httptest.Case
		name string
	}{
		{
			name: "should return empty list when no notes exist",
			tc: httptest.Case{
				ArrangeRequest: func() *http.Request {
					return httptest.NewRequest(http.MethodGet, "/api/v1/notes")
				},
				ExpectStatus: http.StatusOK,
				CheckResponse: func(t *testing.T, body []byte) {
					resp, err := httptest.ParseResponse[listnotes.Response](body)

					require.NoError(t, err)
					assert.Empty(t, resp.Data)
					assert.Equal(t, "0 notes listed", resp.Message)
				},
			},
		},
		{
			name: "should list notes successfully",
			tc: httptest.Case{
				ArrangeRequest: func() *http.Request {
					note1 := domain.NewNote("First Note", "First Content")
					note2 := domain.NewNote("Second Note", "Second Content")

					err := s.repo.Create(t.Context(), note1)
					require.NoError(t, err)

					err = s.repo.Create(t.Context(), note2)
					require.NoError(t, err)

					return httptest.NewRequest(http.MethodGet, "/api/v1/notes")
				},
				ExpectStatus: http.StatusOK,
				CheckResponse: func(t *testing.T, body []byte) {
					resp, err := httptest.ParseResponse[listnotes.Response](body)

					require.NoError(t, err)
					assert.Len(t, resp.Data, 2)
					assert.Equal(t, "2 notes listed", resp.Message)
				},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			httptest.Run(t, s.feat.HTTP, tc.tc)
		})
	}
}
