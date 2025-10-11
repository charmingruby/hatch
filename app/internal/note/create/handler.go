package create

import (
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/transport/http"
	"HATCH_APP/pkg/logger"
	"errors"

	"github.com/gin-gonic/gin"
)

func registerRoute(
	log *logger.Logger,
	api *gin.RouterGroup,
	uc UseCase,
) {
	api.POST("/", func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/CreateNote: request received")

		req, err := http.ParseRequest[Input](c)
		if err != nil {
			log.ErrorContext(
				ctx,
				"endpoint/CreateNote: unable to parse payload",
				"error", err.Error(),
			)

			http.SendBadRequestResponse(c, err.Error())
		}

		op, err := uc.Execute(ctx, Input{
			Title:   req.Title,
			Content: req.Content,
		})
		if err != nil {
			var databaseErr *customerr.DatabaseError
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
	})
}
