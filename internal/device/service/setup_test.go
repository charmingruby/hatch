package service_test

import (
	"github/charmingruby/gew/internal/device/service"
	"github/charmingruby/gew/test/gen/device/mocks"
	"testing"
)

func setupTest(t *testing.T) (*service.Service, *mocks.DeviceRepository) {
	repo := mocks.NewDeviceRepository(t)
	svc := service.New(repo)
	return svc, repo
}
