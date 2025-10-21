package http

import (
	"HATCH_APP/internal/note/infra/http/handler"
	"HATCH_APP/internal/note/usecase"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(log *telemetry.Logger, r *gin.Engine, uc usecase.UseCase) {
	notes := r.Group("/api/v1/notes")
	{
		notes.POST("", handler.CreateHandler(log, uc))
		notes.GET("", handler.FetchHandler(log, uc))
		notes.PATCH(":id", handler.ArchiveHandler(log, uc))
	}
}
