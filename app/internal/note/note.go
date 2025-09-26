package note

import (
	"HATCH_APP/internal/note/http/endpoint"
	"HATCH_APP/internal/note/repository/postgres"
	"HATCH_APP/internal/note/usecase"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(
	r *gin.Engine,
	db *sqlx.DB,
) error {
	repo, err := postgres.NewNoteRepo(db)
	if err != nil {
		return err
	}

	uc := usecase.New(repo)

	val := validator.New()

	endpoint.New(r, val, uc).Register()

	return nil
}
