package archivenote_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/archivenote"
	"HATCH_APP/internal/note/infra/store/postgres"
	"HATCH_APP/pkg/transport/httpx"
	"HATCH_APP/test/container"
	"HATCH_APP/test/httptest"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type httpSuite struct {
	repo *postgres.NoteRepository
	feat *archivenote.Feature
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
		feat: archivenote.New(repo),
	}
}

func TestArchiveNoteEndpoint(t *testing.T) {
	s := setupHTTPSuite(t)
	httptest.Init()

	tests := []struct {
		name string
		tc   httptest.Case
	}{
		{
			name: "should archive note successfully",
			tc: httptest.Case{
				ArrangeRequest: func() *http.Request {
					note := domain.NewNote("Test Note", "Test Content")

					err := s.repo.Create(t.Context(), note)
					require.NoError(t, err)

					return httptest.WithParam(
						httptest.NewRequest(http.MethodPatch, "/api/v1/notes/"+note.ID),
						"id",
						note.ID,
					)
				},
				ExpectStatus: http.StatusNoContent,
				CheckResponse: func(t *testing.T, body []byte) {
					assert.Empty(t, body)
				},
			},
		},
		{
			name: "should return 404 when note not found",
			tc: httptest.Case{
				ExpectStatus: http.StatusNotFound,
				CheckResponse: func(t *testing.T, body []byte) {
					resp, err := httptest.ParseResponse[httpx.ErrorResponse](body)

					require.NoError(t, err)
					assert.Equal(t, "note not found", resp.Message)
				},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			httptest.Run(t, s.feat.ArchiveNoteEndpoint, tc.tc)
		})
	}
}
