package service

import (
	"context"
	"github/charmingruby/pack/internal/example/repository"
)

type CreateDeviceInput struct {
	HardwareID   string
	HardwareType string
}

type CreateDeviceOutput struct {
	DeviceID string
}

type UseCase interface {
	CreateDevice(ctx context.Context, in CreateDeviceInput) (CreateDeviceOutput, error)
}

type Service struct {
	deviceRepo repository.DeviceRepository
}

func New(deviceRepo repository.DeviceRepository) *Service {
	return &Service{
		deviceRepo: deviceRepo,
	}
}
