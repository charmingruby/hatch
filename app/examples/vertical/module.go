package note

import (
	"HATCH_APP/examples/vertical/feature/archive_note"
	"HATCH_APP/examples/vertical/feature/create_note"
	"HATCH_APP/examples/vertical/feature/fetch_notes"
	"HATCH_APP/examples/vertical/infra/db/postgres"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewModule(log *telemetry.Logger, r *gin.Engine, db *sqlx.DB) error {
	repo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	api := r.Group("/api/v1/notes")

	create_note.NewFeature(log, api, repo)
	archive_note.NewFeature(log, api, repo)
	fetch_notes.NewFeature(log, api, repo)

	return nil
}
