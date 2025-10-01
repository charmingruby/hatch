package health

import (
	"HATCH_APP/internal/health/http/endpoint"
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func New(
	log *logger.Logger,
	r *gin.Engine,
	db *postgres.Client,
) {
	endpoint.New(log, r, db).Register()
}

var Module = fx.Module("health",
	fx.Invoke(New),
)
