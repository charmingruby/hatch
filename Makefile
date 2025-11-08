MIGRATIONS_PATH="db/migration"
DATABASE_URL=postgres://postgres:postgres@localhost:5432/hatch?sslmode=disable

###################
# Build           #
###################
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api/main.go

###################
# Database        #
###################
.PHONY: mig-up
mig-up: ## Runs the migrations up
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

.PHONY: mig-down
mig-down: ## Runs the migrations down
	migrate -path ${MIGRATIONS_PATH} -database "$(DATABASE_URL)" down

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(NAME)

###################
# Testing         #
###################
.PHONY: mock
mock:
	mockery --output internal/note/mocks --dir internal/note --all

.PHONY: test
test: mock
	go test ./... -race

###################
# Linting         #
###################
.PHONY: lint
lint:
	go fmt ./...
	golangci-lint run --fix ./...
