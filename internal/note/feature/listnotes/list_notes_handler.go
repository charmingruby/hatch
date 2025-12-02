package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/pkg/errs"
	"HATCH_APP/internal/pkg/http/rest"
	"HATCH_APP/pkg/o11y/logging"
	"errors"

	"github.com/gin-gonic/gin"
)

type Response = []domain.Note

func NewHTTPHandler(uc UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log := logging.FromContext(ctx)

		log.InfoContext(ctx, "endpoint/ListNotes: request received")

		notes, err := uc.Execute(ctx)
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ListNotes: database error",
					"error", databaseErr.Unwrap(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ListNotes: unknown error", "error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		var res = notes

		log.InfoContext(ctx, "endpoint/ListNotes: finished successfully")

		rest.SendOKResponse(
			c,
			"",
			res,
		)
	}
}
