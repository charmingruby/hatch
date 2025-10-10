package note

import (
	"HATCH_APP/internal/note/archive"
	"HATCH_APP/internal/note/create"
	"HATCH_APP/internal/note/fetch"
	"HATCH_APP/internal/note/shared/repository/postgres"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func New(
	log *logger.Logger,
	r *gin.Engine,
	db *sqlx.DB,
	val *validator.Validator,
) error {
	repo, err := postgres.NewNoteRepo(db)
	if err != nil {
		return err
	}

	create.New(log, r, val, repo)
	fetch.New(log, r, val, repo)
	archive.New(log, r, val, repo)

	return nil
}

var Module = fx.Module("note",
	fx.Invoke(New),
)
