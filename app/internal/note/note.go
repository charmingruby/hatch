package note

import (
	"HATCH_APP/internal/note/archive"
	"HATCH_APP/internal/note/create"
	"HATCH_APP/internal/note/fetch"
	"HATCH_APP/internal/note/shared/repository/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func register(
	log *logger.Logger,
	r *gin.Engine,
	db *sqlx.DB,
) error {
	repo, err := postgres.NewNoteRepo(db)
	if err != nil {
		return err
	}

	api := r.Group("/api/v1/notes")

	create.New(log, api, repo)
	fetch.New(log, api, repo)
	archive.New(log, api, repo)

	return nil
}

var Module = fx.Module("note",
	fx.Invoke(register),
)
