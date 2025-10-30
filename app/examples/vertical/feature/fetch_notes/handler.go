package fetch_notes

import (
	"HATCH_APP/examples/vertical/domain"
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/telemetry"
	"errors"

	"github.com/gin-gonic/gin"
)

type Response = []domain.Note

func NewHTTPHandler(log *telemetry.Logger, uc UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/FetchNotes: request received")

		notes, err := uc.Execute(ctx)
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/FetchNotes: database error",
					"error", databaseErr.Unwrap(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/FetchNotes: unknown error", "error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		var res = notes

		log.InfoContext(ctx, "endpoint/FetchNotes: finished successfully")

		rest.SendOKResponse(
			c,
			"",
			res,
		)
	}
}
