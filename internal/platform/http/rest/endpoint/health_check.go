package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func makeHealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
