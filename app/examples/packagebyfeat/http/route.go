package http

import (
	"HATCH_APP/examples/packagebyfeat/core"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(log *telemetry.Logger, r *gin.Engine, uc core.UseCase) {
	api := r.Group("/api/v1/notes")
	{
		api.POST("", CreateNoteHandler(log, uc))
		api.GET("", FetchNotesHandler(log, uc))
		api.PATCH(":id", ArchiveNoteHandler(log, uc))
	}
}
