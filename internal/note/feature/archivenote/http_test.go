package archivenote_test

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/feature/archivenote"
	"HATCH_APP/internal/note/infra/database/postgres"
	"HATCH_APP/pkg/transport/httpx"
	"HATCH_APP/test/container"
	"HATCH_APP/test/testutil"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestHTTP(t *testing.T) {
	tests := []struct {
		name           string
		arrange        func(s *httpSuite) *http.Request
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name: "should archive note successfully",
			arrange: func(s *httpSuite) *http.Request {
				note := domain.NewNote("Test Note", "Test Content")

				err := s.repo.Create(t.Context(), note)
				require.NoError(t, err)

				req := httptest.NewRequest(http.MethodPatch, "/api/v1/notes/"+note.ID, nil)

				return testutil.InjectPathParam(req, "id", note.ID)
			},
			expectedStatus: http.StatusNoContent,
			checkResponse: func(t *testing.T, body []byte) {
				assert.Empty(t, body)
			},
		},
		{
			name: "should return 404 when note not found",
			arrange: func(s *httpSuite) *http.Request {
				id := "nonexistent-id"

				req := httptest.NewRequest(http.MethodPatch, "/api/v1/notes/"+id, nil)

				return testutil.InjectPathParam(req, "id", id)
			},
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, body []byte) {
				var resp httpx.Response

				err := json.Unmarshal(body, &resp)

				require.NoError(t, err)
				assert.Contains(t, resp.Message, "note not found")
			},
		},
	}

	testutil.Init()

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			s := setupHTTPSuite(t)

			req := tc.arrange(s)

			req = testutil.RequestInjection(req)

			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			s.feat.HTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			tc.checkResponse(t, rec.Body.Bytes())
		})
	}
}
