package service

import (
	"context"
	"github/charmingruby/pack/internal/device/model"
)

func (s *Service) CreateDevice(in CreateDeviceInput) (CreateDeviceOuput, error) {
	device := model.NewDevice(model.DeviceInput{
		HardwareID:   in.HardwareID,
		HardwareType: in.HardwareType,
	})

	if err := s.deviceRepo.Create(context.Background(), device); err != nil {
		return CreateDeviceOuput{}, err
	}

	return CreateDeviceOuput{
		DeviceID: device.ID,
	}, nil
}
