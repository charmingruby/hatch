package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"fmt"
	"net/http"
)

type Response struct {
	Message string         `json:"message"`
	Data    []*domain.Note `json:"data"`
}

func (f *Feature) ListNotesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := o11y.LoggerFromContext(ctx).With("feature", "ListNotes")

	notes, err := f.service.ListNotes(ctx)
	if err != nil {
		httpx.WriteError(log, w, err)
		return
	}

	httpx.WriteOKResponse(w, Response{
		Message: fmt.Sprintf("%d notes listed", len(notes)),
		Data:    notes,
	})
}
