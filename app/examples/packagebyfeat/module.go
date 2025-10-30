package packagebyfeat

import (
	"HATCH_APP/examples/packagebyfeat/core"
	"HATCH_APP/examples/packagebyfeat/db/postgres"
	"HATCH_APP/examples/packagebyfeat/http"
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
