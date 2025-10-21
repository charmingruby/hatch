package create

import (
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/internal/shared/http"
	"HATCH_APP/pkg/telemetry"
	"errors"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Title   string `json:"title"   binding:"required" validate:"required,gt=0"`
	Content string `json:"content" binding:"required" validate:"required,gt=0"`
}

func RegisterRoute(log *telemetry.Logger, api *gin.RouterGroup, uc UseCase) {
	api.POST("", handle(log, uc))
}

func handle(log *telemetry.Logger, uc UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/CreateNote: request received")

		req, err := http.ParseRequest[Request](c)
		if err != nil {
			log.ErrorContext(
				ctx,
				"endpoint/CreateNote: unable to parse payload",
				"error", err.Error(),
			)

			http.SendBadRequestResponse(c, err.Error())
		}

		op, err := uc.Execute(ctx, UseCaseInput{
			Title:   req.Title,
			Content: req.Content,
		})
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/CreateNote: database error",
					"error", databaseErr.Unwrap().Error(),
				)

				http.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/CreateNote: unknown error", "error", err.Error(),
			)

			http.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/CreateNote: finished successfully")

		http.SendCreatedResponse(c, op.ID, "note")
	}
}
