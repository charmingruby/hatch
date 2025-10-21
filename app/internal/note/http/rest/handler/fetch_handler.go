package handler

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/internal/note/usecase"
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/telemetry"
	"errors"

	"github.com/gin-gonic/gin"
)

type Response = []domain.Note

func FetchHandler(log *telemetry.Logger, uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/ListNotes: request received")

		notes, err := uc.Fetch(ctx)
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ListNotes: database error",
					"error", databaseErr.Unwrap().Error(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ListNotes: unknown error", "error", err.Error(),
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
