package createnote_test

import (
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/infra/database/postgres"
	"HATCH_APP/pkg/transport/httpx"
	"HATCH_APP/test/container"
	"HATCH_APP/test/testutil"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestHTTP(t *testing.T) {
	tests := []struct {
		name           string
		arrange        func(payload createnote.Request) *http.Request
		checkResponse  func(t *testing.T, body []byte)
		expectedStatus int
	}{
		{
			name: "should create note successfully",
			arrange: func(payload createnote.Request) *http.Request {
				body, _ := json.Marshal(payload)

				req := httptest.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewReader(body))

				return req
			},
			expectedStatus: http.StatusCreated,
			checkResponse: func(t *testing.T, body []byte) {
				data, resp, err := testutil.ParseResponse[createnote.ResponseData](body)

				require.NoError(t, err)
				assert.NotEmpty(t, data.ID)
				assert.Equal(t, "note created", resp.Message)
			},
		},
		{
			name: "should return 400 when body is invalid json",
			arrange: func(payload createnote.Request) *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewReader([]byte("invalid json")))

				return req
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body []byte) {
				var resp httpx.Response

				err := json.Unmarshal(body, &resp)

				require.NoError(t, err)
				assert.Contains(t, resp.Message, "invalid payload")
			},
		},
		{
			name: "should return 400 when title is empty",
			arrange: func(payload createnote.Request) *http.Request {
				payload.Title = ""

				body, _ := json.Marshal(payload)

				req := httptest.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewReader(body))

				return req
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body []byte) {
				var resp httpx.Response

				err := json.Unmarshal(body, &resp)

				require.NoError(t, err)
				assert.Contains(t, resp.Message, "invalid payload")
			},
		},
		{
			name: "should return 400 when content is empty",
			arrange: func(payload createnote.Request) *http.Request {
				payload.Content = ""

				body, _ := json.Marshal(payload)

				req := httptest.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewReader(body))

				return req
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body []byte) {
				var resp httpx.Response

				err := json.Unmarshal(body, &resp)

				require.NoError(t, err)
				assert.Contains(t, resp.Message, "invalid payload")
			},
		},
	}

	testutil.Init()

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			s := setupHTTPSuite(t)

			payload := createnote.Request{
				Title:   "Test Note",
				Content: "Test Content",
			}

			req := tc.arrange(payload)

			req = testutil.RequestInjection(req)

			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			s.feat.HTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)

			tc.checkResponse(t, rec.Body.Bytes())
		})
	}
}
