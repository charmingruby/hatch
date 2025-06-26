package rest

import (
	"github/charmingruby/pack/internal/device/service"
	"github/charmingruby/pack/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createDeviceRequest struct {
	HardwareID   string `json:"hardware_id"   validate:"required,min=1"`
	HardwareType string `json:"hardware_type" validate:"required,min=1"`
}

func createDeviceHandler(svc service.UseCase, v *validator.Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createDeviceRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json: " + err.Error()})
			return
		}

		if err := v.Validate(req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := svc.CreateDevice(service.CreateDeviceInput{
			HardwareID:   req.HardwareID,
			HardwareType: req.HardwareType,
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
