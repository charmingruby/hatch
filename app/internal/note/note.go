package note

import (
	"PACK_APP/internal/note/http/endpoint"
	"PACK_APP/internal/note/repository/postgres"
	"PACK_APP/internal/note/usecase"
	"PACK_APP/pkg/logger"
	"PACK_APP/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(
	log *logger.Logger,
	r *gin.Engine,
	db *sqlx.DB,
) error {
	repo, err := postgres.NewNoteRepo(db)
	if err != nil {
		return err
	}

	uc := usecase.New(repo)

	val := validator.New()

	endpoint.New(r, log, val, uc).Register()

	return nil
}
