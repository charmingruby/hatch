package service

import (
	"context"
	"github/charmingruby/pack/internal/device/delivery/broker"
	"github/charmingruby/pack/internal/device/model"
	"github/charmingruby/pack/pkg/errs"
)

// CreateDevice creates a new device from hardware.
//
// Parameters:
//   - context.Context: shared context.
//   - CreateDeviceInput: input with hardware informations.
//
// Returns:
//   - CreateDeviceOutput: ouput with the created device id.
//   - error: if there is any logic error.
func (s *Service) CreateDevice(ctx context.Context, in CreateDeviceInput) (CreateDeviceOuput, error) {
	deviceFound, err := s.deviceRepo.FindByHardwareIDAndType(ctx, in.HardwareID, in.HardwareType)

	if err != nil {
		return CreateDeviceOuput{}, err
	}

	if deviceFound.ID != "" {
		return CreateDeviceOuput{}, errs.NewResourceAlreadyExistsError("device")
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
