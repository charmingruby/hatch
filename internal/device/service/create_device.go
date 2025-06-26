package service

import (
	"context"
	"github/charmingruby/pack/internal/device/model"
	"github/charmingruby/pack/pkg/core/errs"
)

func (s *Service) CreateDevice(in CreateDeviceInput) (CreateDeviceOuput, error) {
	ctx := context.Background()

	deviceFound, err := s.deviceRepo.FindByHardwareIDAndType(ctx, in.HardwareID, in.HardwareType)

	if err != nil {
		return CreateDeviceOuput{}, err
	}

	if deviceFound.ID != "" {
		return CreateDeviceOuput{}, errs.NewErrResourceAlreadyExists("device")
	}

	device := model.NewDevice(model.DeviceInput{
		HardwareID:   in.HardwareID,
		HardwareType: in.HardwareType,
	})

	if err := s.deviceRepo.Create(ctx, device); err != nil {
		return CreateDeviceOuput{}, err
	}

	return CreateDeviceOuput{
		DeviceID: device.ID,
	}, nil
}
