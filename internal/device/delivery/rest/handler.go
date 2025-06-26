package rest

import (
	"github/charmingruby/pack/internal/device/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createDeviceRequest struct {
	HardwareID   string `json:"hardware_id" binding:"required"  validate:"required,min=1"`
	HardwareType string `json:"hardware_type" binding:"required" validate:"required,min=1"`
}

func createDeviceHandler(svc service.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createDeviceRequest
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
