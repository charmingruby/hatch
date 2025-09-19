package note

import (
	"github.com/charmingruby/pack/internal/note/http/endpoint"
	"github.com/charmingruby/pack/internal/note/repository/postgres"
	"github.com/charmingruby/pack/internal/note/usecase"
	"github.com/charmingruby/pack/pkg/logger"
	"github.com/charmingruby/pack/pkg/validator"
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
