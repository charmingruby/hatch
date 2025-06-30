// Package service contains the core logic of the application.
package service

import (
	"context"
	"github/charmingruby/pack/internal/device/delivery/broker"
	"github/charmingruby/pack/internal/device/repository"
)

// CreateDeviceInput is the input of the CreateDevice usecase.
type CreateDeviceInput struct {
	HardwareID   string
	HardwareType string
}

// CreateDeviceOuput is the output of the CreateDevice usecase.
type CreateDeviceOuput struct {
	DeviceID string
}

// UseCase provides the contracts of the usecases.
type UseCase interface {
	CreateDevice(ctx context.Context, in CreateDeviceInput) (CreateDeviceOuput, error)
}

// Service is the implementation of the contracts of the UseCase.
type Service struct {
	deviceRepo  repository.DeviceRepository
	firmwarePub broker.FirmwarePublisher
}

// New creates the Service instance.
//
// Parameters:
//   - repository.DeviceRepository: devices repository.
//   - broker.FirmwarePublisher: firmware publisher.
//
// Returns:
//   - *Service: usecase implementation.
func New(deviceRepo repository.DeviceRepository, firmwarePub broker.FirmwarePublisher) *Service {
	return &Service{
		deviceRepo:  deviceRepo,
		firmwarePub: firmwarePub,
	}
}
