###################
# Testing         #
###################
.PHONY: mock
mock:
	mockery --output test/gen/CHANGE_BY_YOUR_MODULE/mocks --dir internal/CHANGE_BY_YOUR_MODULE --all

.PHONY: test
test: mock
	go test ./...
