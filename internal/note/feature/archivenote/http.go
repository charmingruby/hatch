package archivenote

import (
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"net/http"
)

func (f *Feature) HTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")

	log := o11y.FromContext(ctx).
		With("feature", "ArchiveNote").
		With("note_id", id)

	if err := f.service.ArchiveNote(ctx, id); err != nil {
		httpx.WriteError(log, w, err)
		return
	}

	httpx.WriteEmptyResponse(w)
}
