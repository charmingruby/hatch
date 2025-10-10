package health

import (
	"HATCH_APP/internal/health/live"
	"HATCH_APP/internal/health/ready"
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
	live.New(log, r)
	ready.New(log, r, db)
}

var Module = fx.Module("health",
	fx.Invoke(New),
)
