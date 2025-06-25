package service_test

import (
	"github/charmingruby/habits/internal/device/mocks"
	"github/charmingruby/habits/internal/device/service"
	"testing"
)

func setupTest(t *testing.T) (*service.Service, *mocks.DeviceRepository) {
	repo := mocks.NewDeviceRepository(t)
	svc := service.New(repo)
	return svc, repo
}
