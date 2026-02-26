package note

import (
	"HATCH_APP/internal/note/feature/archivenote"
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/feature/listnotes"
	"HATCH_APP/internal/note/infra/db/postgres"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Register(r *gin.Engine, db *sqlx.DB) error {
	repo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	createNoteHandler := createnote.New(repo)
	listNotesHandler := listnotes.New(repo)
	archiveNoteHandler := archivenote.New(repo)

	api := r.Group("/api/v1/notes")
	{
		api.POST("", createNoteHandler)
		api.GET("", listNotesHandler)
		api.PATCH(":id", archiveNoteHandler)
	}

	return nil
}
