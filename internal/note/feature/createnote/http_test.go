package createnote_test

import (
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/infra/database/postgres"
	"HATCH_APP/pkg/transport/httpx"
	"HATCH_APP/test/container"
	"HATCH_APP/test/httptest"
	"bytes"
	"encoding/json"
	"net/http"
	stdhttptest "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type httpSuite struct {
	repo *postgres.NoteRepository
	feat *createnote.Feature
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
		feat: createnote.New(repo),
	}
}

func TestCreateNoteEndpoint(t *testing.T) {
	s := setupHTTPSuite(t)
	httptest.Init()

	payload := createnote.Request{
		Title:   "Test Note",
		Content: "Test Content",
	}

	tests := []struct {
		tc   httptest.Case
		name string
	}{
		{
			name: "should create note successfully",
			tc: httptest.Case{
				ArrangeRequest: func() *http.Request {
					body, _ := json.Marshal(payload)

					return stdhttptest.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewReader(body))
				},
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				ExpectStatus: http.StatusCreated,
				CheckResponse: func(t *testing.T, body []byte) {
					resp, err := httptest.ParseResponse[createnote.Response](body)

					require.NoError(t, err)
					assert.NotEmpty(t, resp.Data.ID)
					assert.Equal(t, "note created", resp.Message)
				},
			},
		},
		{
			name: "should return 400 when body is invalid json",
			tc: httptest.Case{
				ArrangeRequest: func() *http.Request {
					return stdhttptest.NewRequest(
						http.MethodPost,
						"/api/v1/notes",
						bytes.NewReader([]byte("invalid json")),
					)
				},
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				ExpectStatus: http.StatusBadRequest,
				CheckResponse: func(t *testing.T, body []byte) {
					resp, err := httptest.ParseResponse[httpx.ErrorResponse](body)

					require.NoError(t, err)
					assert.Contains(t, resp.Message, "invalid payload")
				},
			},
		},
		{
			name: "should return 400 when title is empty",
			tc: httptest.Case{
				ArrangeRequest: func() *http.Request {
					p := payload
					p.Title = ""

					body, _ := json.Marshal(p)

					return stdhttptest.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewReader(body))
				},
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				ExpectStatus: http.StatusBadRequest,
				CheckResponse: func(t *testing.T, body []byte) {
					resp, err := httptest.ParseResponse[httpx.ErrorResponse](body)

					require.NoError(t, err)
					assert.Contains(t, resp.Message, "invalid payload")
				},
			},
		},
		{
			name: "should return 400 when content is empty",
			tc: httptest.Case{
				ArrangeRequest: func() *http.Request {
					p := payload
					p.Content = ""

					body, _ := json.Marshal(p)

					return stdhttptest.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewReader(body))
				},
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				ExpectStatus: http.StatusBadRequest,
				CheckResponse: func(t *testing.T, body []byte) {
					resp, err := httptest.ParseResponse[httpx.ErrorResponse](body)

					require.NoError(t, err)
					assert.Contains(t, resp.Message, "invalid payload")
				},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			httptest.Run(t, s.feat.CreateNoteEndpoint, tc.tc)
		})
	}
}
