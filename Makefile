###################
# Testing         #
###################
.PHONY: mock
mock:
	mockery --output test/gen/device/mocks --dir internal/device --all

.PHONY: test
test: mock
	go test ./...
