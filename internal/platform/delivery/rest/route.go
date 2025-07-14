package rest

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")

	v1.POST("/healthz", makeHealthCheckHandler())
}
