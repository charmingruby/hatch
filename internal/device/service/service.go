package service

import (
	"github/charmingruby/gew/internal/device/dto"
	"github/charmingruby/gew/internal/device/repository"
)

type UseCase interface {
	CreateDevice(dto dto.CreateDeviceInput) error
}

type Service struct {
	deviceRepo repository.DeviceRepository
}

func New(deviceRepo repository.DeviceRepository) *Service {
	return &Service{
		deviceRepo: deviceRepo,
	}
}
