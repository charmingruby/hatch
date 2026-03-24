package createnote

import (
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"net/http"
)

type Request struct {
	Title   string `json:"title"   validate:"required,gt=0"`
	Content string `json:"content" validate:"required,gt=0"`
}

type Response struct {
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	ID string `json:"id"`
}

func (f *Feature) HTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := o11y.FromContext(ctx).With("feature", "CreateNote")

	req, err := httpx.ParseRequest[Request](w, r)
	if err != nil {
		log.WarnContext(ctx, "invalid payload", "error", err)

		return
	}

	id, err := f.service.CreateNote(ctx, req.Title, req.Content)
	if err != nil {
		//nolint:gocritic // keep single-case switch for consistency with other handlers and declarative behaviours.
		switch {
		default:
			log.ErrorContext(ctx, "execute create note failed", "error", err)
			httpx.WriteInternalServerErrorResponse(w)
			return
		}
	}

	httpx.WriteCreatedResponse(w, Response{
		Message: "note created",
		Data:    ResponseData{ID: id},
	})
}
