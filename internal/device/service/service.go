package service

import (
	"context"
	"github/charmingruby/pack/internal/device/delivery/broker"
	"github/charmingruby/pack/internal/device/repository"
)

type CreateDeviceInput struct {
	HardwareID   string
	HardwareType string
}

type CreateDeviceOuput struct {
	DeviceID string
}

type UseCase interface {
	CreateDevice(ctx context.Context, in CreateDeviceInput) (CreateDeviceOuput, error)
}

type Service struct {
	deviceRepo  repository.DeviceRepository
	firmwarePub broker.FirmwarePublisher
}

func New(deviceRepo repository.DeviceRepository, firmwarePub broker.FirmwarePublisher) *Service {
	return &Service{
		deviceRepo:  deviceRepo,
		firmwarePub: firmwarePub,
	}
}
