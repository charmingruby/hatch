package archivenote

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"errors"
	"net/http"
)

func NewHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		log := o11y.FromContext(ctx).With("feature", "ArchiveNote")

		id := r.PathValue("id")
		log = log.With("note_id", id)

		if err := svc.Execute(ctx, id); err != nil {
			if errors.Is(err, domain.ErrNoteNotFound) {
				log.WarnContext(ctx, "note not found", "error", err)

				httpx.WriteNotFoundResponse(w, err.Error())
				return
			}

			log.ErrorContext(ctx, "execute archive note failed", "error", err)

			httpx.WriteInternalServerErrorResponse(w)
			return
		}

		httpx.WriteEmptyResponse(w)
	}
}
