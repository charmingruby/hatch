package fetch

import (
	"HATCH_APP/internal/note/shared/model"
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/internal/shared/http"
	"HATCH_APP/pkg/telemetry"
	"errors"

	"github.com/gin-gonic/gin"
)

type Response = []model.Note

func RegisterRoute(log *telemetry.Logger, api *gin.RouterGroup, uc UseCase) {
	api.GET("", handle(log, uc))
}

func handle(log *telemetry.Logger, uc UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/ListNotes: request received")

		op, err := uc.Execute(ctx)
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ListNotes: database error",
					"error", databaseErr.Unwrap().Error(),
				)

				http.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ListNotes: unknown error", "error", err.Error(),
			)

			http.SendInternalServerErrorResponse(c)
			return
		}

		var res = op.Notes

		log.InfoContext(ctx, "endpoint/ListNotes: finished successfully")

		http.SendOKResponse(
			c,
			"",
			res,
		)
	}
}
