# Adding New Module

Complete guide to create a new feature module in Hatch.

## Steps

1. Create directory structure
2. Implement layers in order: model � DTO � use case + **tests** � repository � handlers
3. Wire dependencies in barrel file with Fx module
4. Register module in `main.go`

**CRITICAL**: Every use case MUST have unit tests. No exceptions.

## Directory Setup

```bash
mkdir -p internal/billing/{model,dto,usecase,repository/postgres,http/endpoint}
```

## Complete Implementation Example: Billing Module

### 1. Model (internal/billing/model/subscription.go)

```go
package model

import "time"

type Subscription struct {
	ID        string
	UserID    string
	PlanID    string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSubscription(userID, planID string) Subscription {
	now := time.Now()
	return Subscription{
		ID:        id.New(),
		UserID:    userID,
		PlanID:    planID,
		Status:    "active",
		CreatedAt: now,
		UpdatedAt: now,
	}
}
```

### 2. DTO (internal/billing/dto/create_subscription_dto.go)

```go
package dto

type CreateSubscriptionInput struct {
	UserID string
	PlanID string
}

type CreateSubscriptionOutput struct {
	ID string
}
```

### 3. Use Case Interface (internal/billing/usecase/usecase.go)

```go
package usecase

import (
	"context"
	"HATCH_APP/internal/billing/dto"
	"HATCH_APP/internal/billing/repository"
)

type Service interface {
	CreateSubscription(ctx context.Context, input dto.CreateSubscriptionInput) (dto.CreateSubscriptionOutput, error)
}

type UseCase struct {
	repo repository.SubscriptionRepository
}

func New(repo repository.SubscriptionRepository) UseCase {
	return UseCase{repo: repo}
}
```

### 4. Use Case Implementation (internal/billing/usecase/create_subscription.go)

```go
package usecase

import (
	"context"
	"HATCH_APP/internal/billing/dto"
	"HATCH_APP/internal/billing/model"
	"HATCH_APP/internal/shared/customerr"
)

func (u UseCase) CreateSubscription(ctx context.Context, input dto.CreateSubscriptionInput) (dto.CreateSubscriptionOutput, error) {
	subscription := model.NewSubscription(input.UserID, input.PlanID)

	if err := u.repo.Create(ctx, subscription); err != nil {
		return dto.CreateSubscriptionOutput{}, customerr.NewDatabaseError(err)
	}

	return dto.CreateSubscriptionOutput{ID: subscription.ID}, nil
}
```

### 4b. REQUIRED: Test Setup (internal/billing/usecase/setup_test.go)

```go
package usecase_test

import (
	"testing"
	"HATCH_APP/internal/billing/usecase"
	"HATCH_APP/test/gen/billing/mocks"
)

type suite struct {
	repo    *mocks.SubscriptionRepository
	usecase usecase.Service
}

func setup(t *testing.T) suite {
	repo := mocks.NewSubscriptionRepository(t)

	return suite{
		repo:    repo,
		usecase: usecase.New(repo),
	}
}
```

### 4c. REQUIRED: Unit Tests (internal/billing/usecase/create_subscription_test.go)

```go
package usecase_test

import (
	"errors"
	"testing"

	"HATCH_APP/internal/billing/dto"
	"HATCH_APP/internal/billing/model"
	"HATCH_APP/internal/shared/customerr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CreateSubscription(t *testing.T) {
	userID := "user-123"
	planID := "plan-456"

	t.Run("should create successfully", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Create", t.Context(), mock.MatchedBy(func(sub model.Subscription) bool {
			return sub.UserID == userID && sub.PlanID == planID
		})).
			Return(nil).
			Once()

		output, err := s.usecase.CreateSubscription(t.Context(), dto.CreateSubscriptionInput{
			UserID: userID,
			PlanID: planID,
		})

		require.NoError(t, err)
		assert.NotEmpty(t, output.ID)
	})

	t.Run("should return a DatabaseError when there is a datasource error", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Create", t.Context(), mock.Anything).
			Return(errors.New("unhealthy repo")).
			Once()

		output, err := s.usecase.CreateSubscription(t.Context(), dto.CreateSubscriptionInput{
			UserID: userID,
			PlanID: planID,
		})

		assert.Zero(t, output)
		require.Error(t, err)

		var targetErr *customerr.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
```

### 5. Repository Interface (internal/billing/repository/repository.go)

```go
package repository

import (
	"context"
	"HATCH_APP/internal/billing/model"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription model.Subscription) error
	FindByID(ctx context.Context, id string) (model.Subscription, error)
}
```

### 6. Repository Implementation (internal/billing/repository/postgres/subscription_repository.go)

```go
package postgres

import (
	"context"
	"database/sql"
	"time"

	"HATCH_APP/internal/billing/model"
	"HATCH_APP/pkg/database/postgres"

	"github.com/jmoiron/sqlx"
)

type SubscriptionRepo struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewSubscriptionRepo(db *sqlx.DB) (*SubscriptionRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range subscriptionQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil, postgres.NewPreparationErr(queryName, "subscription", err)
		}
		stmts[queryName] = stmt
	}

	return &SubscriptionRepo{db: db, stmts: stmts}, nil
}

func (r *SubscriptionRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]
	if !ok {
		return nil, postgres.NewStatementNotPreparedErr(queryName, "subscription")
	}
	return stmt, nil
}

func (r *SubscriptionRepo) Create(ctx context.Context, sub model.Subscription) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(createSubscription)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, &sub.ID, &sub.UserID, &sub.PlanID, &sub.Status, &sub.CreatedAt, &sub.UpdatedAt)
	return err
}

func (r *SubscriptionRepo) FindByID(ctx context.Context, id string) (model.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findSubscriptionByID)
	if err != nil {
		return model.Subscription{}, err
	}

	var sub model.Subscription
	if err := stmt.QueryRowContext(ctx, id).Scan(
		&sub.ID, &sub.UserID, &sub.PlanID, &sub.Status, &sub.CreatedAt, &sub.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Subscription{}, nil
		}
		return model.Subscription{}, err
	}

	return sub, nil
}
```

### 7. SQL Queries (internal/billing/repository/postgres/subscription_query.go)

```go
package postgres

const (
	createSubscription   = "createSubscription"
	findSubscriptionByID = "findSubscriptionByID"
)

func subscriptionQueries() map[string]string {
	return map[string]string{
		createSubscription: `
			INSERT INTO subscriptions (id, user_id, plan_id, status, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`,
		findSubscriptionByID: `
			SELECT id, user_id, plan_id, status, created_at, updated_at
			FROM subscriptions WHERE id = $1
		`,
	}
}
```

### 8. HTTP Handler (internal/billing/http/endpoint/create_subscription_endpoint.go)

```go
package endpoint

import (
	"HATCH_APP/internal/billing/dto"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"
	"errors"

	"github.com/gin-gonic/gin"
)

type CreateSubscriptionRequest struct {
	UserID string `json:"user_id" binding:"required" validate:"required"`
	PlanID string `json:"plan_id" binding:"required" validate:"required"`
}

func (e *Endpoint) CreateSubscription(c *gin.Context) {
	ctx := c.Request.Context()

	e.log.InfoContext(ctx, "endpoint/CreateSubscription: request received")

	var req CreateSubscriptionRequest
	if err := c.BindJSON(&req); err != nil {
		e.log.ErrorContext(ctx, "endpoint/CreateSubscription: unable to parse payload", "error", err.Error())
		rest.SendBadRequestResponse(c, err.Error())
		return
	}

	if err := e.val.Validate(req); err != nil {
		e.log.ErrorContext(ctx, "endpoint/CreateSubscription: invalid payload", "error", err.Error())
		rest.SendBadRequestResponse(c, err.Error())
		return
	}

	output, err := e.service.CreateSubscription(ctx, dto.CreateSubscriptionInput{
		UserID: req.UserID,
		PlanID: req.PlanID,
	})
	if err != nil {
		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			e.log.ErrorContext(ctx, "endpoint/CreateSubscription: database error", "error", databaseErr.Unwrap().Error())
			rest.SendInternalServerErrorResponse(c)
			return
		}

		e.log.ErrorContext(ctx, "endpoint/CreateSubscription: unknown error", "error", err.Error())
		rest.SendInternalServerErrorResponse(c)
		return
	}

	e.log.InfoContext(ctx, "endpoint/CreateSubscription: finished successfully")
	rest.SendCreatedResponse(c, output.ID, "subscription")
}
```

### 9. Endpoint Registration (internal/billing/http/endpoint/endpoint.go)

```go
package endpoint

import (
	"HATCH_APP/internal/billing/usecase"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	log     *logger.Logger
	val     *validator.Validator
	service usecase.Service
}

func New(log *logger.Logger, val *validator.Validator, service usecase.Service) *Endpoint {
	return &Endpoint{log: log, val: val, service: service}
}

func (e *Endpoint) Register(r *gin.Engine) {
	r.POST("/subscriptions", e.CreateSubscription)
}
```

### 10. Barrel File with Fx Module (internal/billing/billing.go)

```go
package billing

import (
	"HATCH_APP/internal/billing/http/endpoint"
	"HATCH_APP/internal/billing/repository/postgres"
	"HATCH_APP/internal/billing/usecase"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func New(log *logger.Logger, r *gin.Engine, db *sqlx.DB) error {
	repo, err := postgres.NewSubscriptionRepo(db)
	if err != nil {
		return err
	}

	uc := usecase.New(repo)
	val := validator.New()

	endpoint.New(log, val, uc).Register(r)
	return nil
}

var Module = fx.Module("billing", fx.Invoke(New))
```

### 11. Register in main.go (cmd/api/main.go)

```go
package main

import (
	"time"

	"HATCH_APP/config"
	"HATCH_APP/internal/billing"  // � Import new module
	"HATCH_APP/internal/health"
	"HATCH_APP/internal/note"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("env: missing", "error", err)
	}

	fx.New(
		fx.Supply(log),
		config.Module,
		postgres.Module,
		rest.Module,
		health.Module,
		note.Module,
		billing.Module,  // � Add new module here
		fx.WithLogger(func() fxevent.Logger {
			return log
		}),
		fx.StartTimeout(30*time.Second),
		fx.StopTimeout(15*time.Second),
	).Run()
}
```

## After Implementation

1. **Generate mocks**: `make mock`
2. **Run tests**: `make test`
3. **Create migration**: `make new-mig name=creates_subscriptions_table`
4. **Edit migration files** in `db/migrations/` (created automatically with timestamp)
5. **Run migration**: `make migrate-up`
6. **Test endpoint**: `curl -X POST http://localhost:8080/subscriptions -d '{"user_id":"user-123","plan_id":"plan-456"}'`

## Reference

See the complete working example in [internal/note/](../../app/internal/note/)
