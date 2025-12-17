package createnote

import (
	"HATCH_APP/internal/common/errs"
	"HATCH_APP/internal/common/http/rest"
	"HATCH_APP/pkg/o11y/logging"
	"errors"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Title   string `json:"title"   binding:"required" validate:"required,gt=0"`
	Content string `json:"content" binding:"required" validate:"required,gt=0"`
}

func NewHTTPHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log := logging.FromContext(ctx)

		log.InfoContext(ctx, "endpoint/CreateNote: request received")

		req, err := rest.ParseRequest[Request](c)
		if err != nil {
			log.ErrorContext(
				ctx,
				"endpoint/CreateNote: unable to parse payload",
				"error", err,
			)

			rest.SendBadRequestResponse(c, err.Error())
			return
		}

		id, err := svc.Execute(ctx, req.Title, req.Content)
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/CreateNote: database error",
					"error", databaseErr.Unwrap(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/CreateNote: unknown error", "error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/CreateNote: finished successfully")

		rest.SendCreatedResponse(c, id, "note")
	}
}
