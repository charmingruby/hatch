package note

import (
	"HATCH_APP/internal/note/db/repository/postgres"
	"HATCH_APP/internal/note/http/rest"
	"HATCH_APP/internal/note/usecase"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func register(log *telemetry.Logger, r *gin.Engine, db *sqlx.DB) error {
	repo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	usecase := usecase.NewService(repo)

	rest.RegisterRoutes(log, r, usecase)

	return nil
}

var Module = fx.Module("note",
	fx.Invoke(register),
)
