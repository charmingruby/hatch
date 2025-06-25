package service_test

import (
	"errors"
	"github/charmingruby/habits/internal/device/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Service_CreateDevice_Success(t *testing.T) {
	svc, repo := setupTest(t)

	repo.On("Create", mock.Anything).Return(nil)

	err := svc.CreateDevice(dto.CreateDeviceInput{
		HardwareID:   "1",
		HardwareType: "Solar",
		TriggeredAt:  time.Now(),
	})

	assert.NoError(t, err)
}

func Test_Service_CreateDevice_RepositoryErr(t *testing.T) {
	svc, repo := setupTest(t)

	repo.On("Create", mock.Anything).Return(errors.New("operation error"))

	err := svc.CreateDevice(dto.CreateDeviceInput{
		HardwareID:   "1",
		HardwareType: "Solar",
		TriggeredAt:  time.Now(),
	})

	assert.Error(t, err)
}
