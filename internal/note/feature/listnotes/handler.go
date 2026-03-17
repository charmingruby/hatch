package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"net/http"
)

type Response = []domain.Note

func NewHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		log := o11y.FromContext(ctx).With("feature", "ListNotes")

		notes, err := svc.Execute(ctx)
		if err != nil {
			//nolint:gocritic // keep single-case switch for consistency with other handlers and declarative behaviours.
			switch {
			default:
				log.ErrorContext(ctx, "execute list notes failed", "error", err)
				httpx.WriteInternalServerErrorResponse(w)
				return
			}
		}

		var res = notes

		httpx.WriteOKResponse(
			w,
			"",
			res,
		)
	}
}
