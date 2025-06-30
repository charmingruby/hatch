package service

import (
	"context"
	"github/charmingruby/pack/internal/device/delivery/broker"
	"github/charmingruby/pack/internal/device/model"
	"github/charmingruby/pack/pkg/core/errs"
)

func (s *Service) CreateDevice(ctx context.Context, in CreateDeviceInput) (CreateDeviceOuput, error) {
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

	if err := s.firmwarePub.DispatchDeviceRegistered(broker.DeviceRegisteredMessage{
		DeviceID: device.ID,
	}); err != nil {
		return CreateDeviceOuput{}, err
	}

	return CreateDeviceOuput{
		DeviceID: device.ID,
	}, nil
}
