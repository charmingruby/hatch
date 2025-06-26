package service_test

import (
	"errors"
	"github/charmingruby/gew/internal/device/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Service_CreateDevice_Success(t *testing.T) {
	svc, repo := setupTest(t)

	repo.On("Create", mock.Anything).Return(nil)

	err := svc.CreateDevice(service.CreateDeviceInput{
		HardwareID:   "1",
		HardwareType: "Solar",
	})

	assert.NoError(t, err)
}

func Test_Service_CreateDevice_RepositoryErr(t *testing.T) {
	svc, repo := setupTest(t)

	repo.On("Create", mock.Anything).Return(errors.New("operation error"))

	err := svc.CreateDevice(service.CreateDeviceInput{
		HardwareID:   "1",
		HardwareType: "Solar",
	})

	assert.Error(t, err)
}
