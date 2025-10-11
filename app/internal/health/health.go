package health

import (
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func register(
	log *logger.Logger,
	r *gin.Engine,
	db *postgres.Client,
) {

	registerRoutes(log, r, db)
}

var Module = fx.Module("health",
	fx.Invoke(register),
)
