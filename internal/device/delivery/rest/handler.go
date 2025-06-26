package rest

import (
	"github/charmingruby/pack/internal/device/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDeviceRequest struct {
	HardwareID   string `json:"hardware_id"   validate:"required,min=1"`
	HardwareType string `json:"hardware_type" validate:"required,min=1"`
}

func CreateDevice(svc service.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateDeviceRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := svc.CreateDevice(service.CreateDeviceInput{
			HardwareID:   req.HardwareID,
			HardwareType: req.HardwareType,
		}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.Status(http.StatusCreated)
	}
}
