package note

import (
	"HATCH_APP/internal/note/core"
	"HATCH_APP/internal/note/http"
	"HATCH_APP/internal/note/infra/postgres"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewModule(log *telemetry.Logger, r *gin.Engine, db *sqlx.DB) error {
	repo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	usecase := core.NewService(repo)

	http.RegisterRoutes(log, r, usecase)

	return nil
}
