package note

import (
	"HATCH_APP/internal/note/feature/archivenote"
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/feature/listnotes"
	"HATCH_APP/internal/note/infra/db/postgres"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Scaffold(log *telemetry.Logger, r *gin.Engine, db *sqlx.DB) error {
	repo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	createNoteHandler := createnote.NewFeature(log, repo)
	listNotesHandler := listnotes.NewFeature(log, repo)
	archiveNoteHandler := archivenote.NewFeature(log, repo)

	api := r.Group("/api/v1/notes")
	{
		api.POST("", createNoteHandler)
		api.GET("", listNotesHandler)
		api.PATCH(":id", archiveNoteHandler)
	}

	return nil
}
