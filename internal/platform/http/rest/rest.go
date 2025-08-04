package rest

import (
	"github/charmingruby/pack/internal/platform/http/rest/endpoint"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	endpoint.New(r)
}
