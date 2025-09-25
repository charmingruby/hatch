package health

import (
	"HATCH_APP/internal/health/http/endpoint"
	"HATCH_APP/pkg/database/postgres"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *postgres.Client) {
	endpoint.New(r, db).Register()
}
