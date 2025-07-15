package service

import (
	"context"
	"github/charmingruby/pack/internal/example/model"
	"github/charmingruby/pack/pkg/core/errs"
)

func (s *Service) CreateDevice(ctx context.Context, in CreateDeviceInput) (CreateDeviceOutput, error) {
	deviceFound, err := s.deviceRepo.FindByHardwareIDAndType(ctx, in.HardwareID, in.HardwareType)

	if err != nil {
		return CreateDeviceOutput{}, err
	}

	if deviceFound.ID != "" {
		return CreateDeviceOutput{}, errs.NewResourceAlreadyExistsError("device")
	}

	device := model.NewDevice(model.DeviceInput{
		HardwareID:   in.HardwareID,
		HardwareType: in.HardwareType,
	})

	if err := s.deviceRepo.Create(ctx, device); err != nil {
		return CreateDeviceOutput{}, err
	}

	return CreateDeviceOutput{
		DeviceID: device.ID,
	}, nil
}
