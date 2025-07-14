package service_test

import (
	"errors"
	"github/charmingruby/pack/internal/device/model"
	"github/charmingruby/pack/internal/device/service"
	"github/charmingruby/pack/pkg/core/errs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Service_CreateDevice_Success(t *testing.T) {
	svc, repo := setupTest(t)

	repo.On("FindByHardwareIDAndType", mock.Anything, mock.Anything, mock.Anything).
		Return(model.Device{}, nil)

	repo.On("Create", mock.Anything, mock.Anything).Return(nil)

	_, err := svc.CreateDevice(t.Context(), service.CreateDeviceInput{
		HardwareID:   "1",
		HardwareType: "Solar",
	})

	assert.NoError(t, err)
}

func Test_Service_CreateDevice_DeviceAlreadyExistsErr(t *testing.T) {
	svc, repo := setupTest(t)

	repo.On("FindByHardwareIDAndType", mock.Anything, "1", "Solar").
		Return(model.Device{ID: "existing-id"}, nil)

	_, err := svc.CreateDevice(t.Context(), service.CreateDeviceInput{
		HardwareID:   "1",
		HardwareType: "Solar",
	})

	require.Error(t, err)

	var alreadyExistsErr *errs.ResourceAlreadyExistsError
	assert.ErrorAs(t, err, &alreadyExistsErr)
}

func Test_Service_CreateDevice_RepositoryErr(t *testing.T) {
	svc, repo := setupTest(t)

	repo.On("FindByHardwareIDAndType", mock.Anything, mock.Anything, mock.Anything).
		Return(model.Device{}, nil)

	repo.On("Create", mock.Anything, mock.Anything).Return(errors.New("operation error"))

	_, err := svc.CreateDevice(t.Context(), service.CreateDeviceInput{
		HardwareID:   "1",
		HardwareType: "Solar",
	})

	assert.Error(t, err)
}
