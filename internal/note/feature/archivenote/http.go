package archivenote

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"errors"
	"net/http"
)

func (f *Feature) HTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := o11y.FromContext(ctx).With("feature", "ArchiveNote")

	id := r.PathValue("id")
	log = log.With("note_id", id)

	if err := f.service.ArchiveNote(ctx, id); err != nil {
		switch {
		case errors.Is(err, domain.ErrNoteNotFound):
			log.WarnContext(ctx, "note not found", "error", err)

			httpx.WriteNotFoundResponse(w, httpx.Response{
				Message: "note not found",
			})

			return
		default:
			log.ErrorContext(ctx, "execute archive note failed", "error", err)
			httpx.WriteInternalServerErrorResponse(w)
			return
		}
	}

	httpx.WriteEmptyResponse(w)
}
