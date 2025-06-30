package rest

import (
	"errors"
	"github/charmingruby/pack/internal/device/service"
	"github/charmingruby/pack/pkg/errs"
	"github/charmingruby/pack/pkg/http/rest"
	"github/charmingruby/pack/pkg/telemetry/logger"
	"github/charmingruby/pack/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createDeviceRequest struct {
	HardwareID   string `json:"hardware_id"   validate:"required,min=1"`
	HardwareType string `json:"hardware_type" validate:"required,min=1"`
}

func createDeviceHandler(log *logger.Logger, svc service.UseCase, v *validator.Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createDeviceRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			rest.SendBadRequestErrorResponse(ctx, err.Error())
			return
		}

		if err := v.Validate(req); err != nil {
			rest.SendBadRequestErrorResponse(ctx, err.Error())
			return
		}

		op, err := svc.CreateDevice(ctx.Request.Context(), service.CreateDeviceInput{
			HardwareID:   req.HardwareID,
			HardwareType: req.HardwareType,
		})
		if err != nil {
			var resourceAlreadyExistsErr *errs.ResourceAlreadyExistsError
			if errors.As(err, &resourceAlreadyExistsErr) {
				rest.SendConflictErrorResponse(ctx, err.Error())
				return
			}

			log.Error("uncaught error", "error", err)

			rest.SendUncaughtErrorResponse(ctx)
			return
		}

		rest.SendCreatedResponse(ctx, "", op.DeviceID, "device")
	}
}
