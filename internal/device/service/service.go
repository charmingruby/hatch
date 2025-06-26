package service

import (
	"github/charmingruby/pack/internal/device/repository"
)

type CreateDeviceInput struct {
	HardwareID   string
	HardwareType string
}

type UseCase interface {
	CreateDevice(in CreateDeviceInput) error
}

type Service struct {
	deviceRepo repository.DeviceRepository
}

func New(deviceRepo repository.DeviceRepository) *Service {
	return &Service{
		deviceRepo: deviceRepo,
	}
}
