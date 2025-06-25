package service

import (
	"github/charmingruby/gew/internal/device/dto"
	"github/charmingruby/gew/internal/device/model"
)

func (s *Service) CreateDevice(dto dto.CreateDeviceInput) error {
	device := model.NewDevice(model.DeviceInput{
		HardwareID:   dto.HardwareID,
		HardwareType: dto.HardwareType,
		IconURL:      "dummy for moment",
	})

	if err := s.deviceRepo.Create(device); err != nil {
		return err
	}

	return nil
}
