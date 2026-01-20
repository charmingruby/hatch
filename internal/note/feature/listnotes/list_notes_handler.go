package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/http/rest"
	"HATCH_APP/pkg/o11y"

	"github.com/gin-gonic/gin"
)

type Response = []domain.Note

func NewHTTPHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log := o11y.FromContext(ctx)

		log.InfoContext(ctx, "endpoint/ListNotes: request received")

		notes, err := svc.Execute(ctx)
		if err != nil {
			log.ErrorContext(
				ctx,
				"endpoint/ListNotes: internal error",
				"error", err,
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
