package service_test

import (
	"github/charmingruby/pack/internal/device/service"
	"github/charmingruby/pack/test/gen/device/mocks"
	"testing"
)

func setupTest(t *testing.T) (*service.Service, *mocks.DeviceRepository, *mocks.FirmwarePublisher) {
	repo := mocks.NewDeviceRepository(t)
	pub := mocks.NewFirmwarePublisher(t)
	svc := service.New(repo, pub)

	return svc, repo, pub
}
