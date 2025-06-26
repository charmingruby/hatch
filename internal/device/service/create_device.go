package service

import (
	"github/charmingruby/pack/internal/device/model"
)

func (s *Service) CreateDevice(in CreateDeviceInput) error {
	device := model.NewDevice(model.DeviceInput{
		HardwareID:   in.HardwareID,
		HardwareType: in.HardwareType,
	})

	if err := s.deviceRepo.Create(device); err != nil {
		return err
	}

	return nil
}
