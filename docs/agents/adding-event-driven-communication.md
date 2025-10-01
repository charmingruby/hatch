# Adding Event-Driven Communication

Guide to add asynchronous event-driven communication between modules.

## Overview

Events enable loose coupling between modules. One module publishes events, and other modules subscribe to react to them without direct dependencies.

## Steps

1. Create `MODULE/messaging/event/` directory
2. Define event struct with Serialize/Deserialize methods
3. Define topic constant (`module.action`)
4. Create `MODULE/messaging/subscriber/` directory (in consuming module)
5. Implement listener with Subscribe registration
6. Implement handler function (`on<Event>`)
7. Wire subscriber in barrel file

## Complete Example: Subscription Events

### 1. Event Definition (internal/billing/messaging/event/subscription_created.go)

```go
package event

import (
	"encoding/json"
	"time"
)

const SubscriptionCreatedTopic = "subscription.created"

type SubscriptionCreatedMessage struct {
	SourcedAt      time.Time `json:"sourced_at"`
	SubscriptionID string    `json:"subscription_id"`
	UserID         string    `json:"user_id"`
	PlanID         string    `json:"plan_id"`
}

func (m *SubscriptionCreatedMessage) Serialize() ([]byte, error) {
	payload, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (m *SubscriptionCreatedMessage) Deserialize(payload []byte) error {
	if err := json.Unmarshal(payload, m); err != nil {
		return err
	}
	return nil
}
```

### 2. Publishing Events (internal/billing/usecase/create_subscription.go)

Modify the use case to publish the event after creating the subscription:

```go
package usecase

import (
	"context"
	"time"

	"HATCH_APP/internal/billing/dto"
	"HATCH_APP/internal/billing/messaging/event"
	"HATCH_APP/internal/billing/model"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/pkg/messaging"
)

type UseCase struct {
	repo      repository.SubscriptionRepository
	publisher messaging.Publisher  // ê Add publisher dependency
}

func New(repo repository.SubscriptionRepository, pub messaging.Publisher) UseCase {
	return UseCase{
		repo:      repo,
		publisher: pub,
	}
}

func (u UseCase) CreateSubscription(ctx context.Context, input dto.CreateSubscriptionInput) (dto.CreateSubscriptionOutput, error) {
	subscription := model.NewSubscription(input.UserID, input.PlanID)

	if err := u.repo.Create(ctx, subscription); err != nil {
		return dto.CreateSubscriptionOutput{}, customerr.NewDatabaseError(err)
	}

	// Publish event
	msg := event.SubscriptionCreatedMessage{
		SourcedAt:      time.Now(),
		SubscriptionID: subscription.ID,
		UserID:         subscription.UserID,
		PlanID:         subscription.PlanID,
	}

	payload, err := msg.Serialize()
	if err != nil {
		u.log.Warn("failed to serialize event", "error", err)
		// Don't fail the operation if event serialization fails
	} else {
		if err := u.publisher.Publish(ctx, event.SubscriptionCreatedTopic, payload); err != nil {
			u.log.Warn("failed to publish event", "error", err)
			// Don't fail the operation if event publishing fails
		}
	}

	return dto.CreateSubscriptionOutput{ID: subscription.ID}, nil
}
```

### 3. Subscriber Listener (internal/notification/messaging/subscriber/listener.go)

In the consuming module (notification), create a listener:

```go
package subscriber

import (
	"HATCH_APP/internal/billing/messaging/event"
	"HATCH_APP/internal/notification/usecase"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/messaging"
)

type Listener struct {
	log *logger.Logger
	sub messaging.Subscriber
	uc  usecase.Service
}

func New(log *logger.Logger, sub messaging.Subscriber, uc usecase.Service) Listener {
	return Listener{
		log: log,
		sub: sub,
		uc:  uc,
	}
}

func (l *Listener) Listen() error {
	// Subscribe to multiple events
	if err := l.sub.Subscribe(event.SubscriptionCreatedTopic, l.onSubscriptionCreated); err != nil {
		return err
	}

	return nil
}
```

### 4. Event Handler (internal/notification/messaging/subscriber/on_subscription_created.go)

```go
package subscriber

import (
	"context"

	"HATCH_APP/internal/billing/messaging/event"
	"HATCH_APP/internal/notification/dto"
	"HATCH_APP/pkg/messaging"
)

func (l *Listener) onSubscriptionCreated(e messaging.Envelope) error {
	var msg event.SubscriptionCreatedMessage

	if err := msg.Deserialize(e.Message); err != nil {
		l.log.Error("failed to deserialize message", "error", err)
		return err
	}

	l.log.Debug("subscription created event received", "subscription_id", msg.SubscriptionID)

	// Call use case to handle business logic
	ctx := context.Background()
	return l.uc.SendWelcomeEmail(ctx, dto.SendWelcomeEmailInput{
		UserID:         msg.UserID,
		SubscriptionID: msg.SubscriptionID,
	})
}
```

### 5. Wire Subscriber in Barrel File (internal/notification/notification.go)

```go
package notification

import (
	"HATCH_APP/internal/notification/messaging/subscriber"
	"HATCH_APP/internal/notification/usecase"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/messaging"

	"go.uber.org/fx"
)

func New(log *logger.Logger, sub messaging.Subscriber, uc usecase.Service) error {
	listener := subscriber.New(log, sub, uc)

	// Start listening to events
	if err := listener.Listen(); err != nil {
		return err
	}

	return nil
}

var Module = fx.Module("notification", fx.Invoke(New))
```

## Topic Naming Convention

Follow the pattern: `module.action`

Examples:
- `subscription.created`
- `subscription.cancelled`
- `subscription.renewed`
- `payment.processed`
- `payment.failed`
- `user.registered`
- `user.deleted`

## Event Structure Best Practices

1. **Include timestamp**: Always add `SourcedAt time.Time` field
2. **Include IDs**: Reference entities by ID, not full objects
3. **Immutable**: Events represent facts that happened, never modify them
4. **Self-contained**: Include all data needed by consumers
5. **Versioned**: Consider adding version field for schema evolution

```go
type ExampleMessage struct {
	Version   string    `json:"version"`    // e.g., "v1"
	SourcedAt time.Time `json:"sourced_at"`
	EntityID  string    `json:"entity_id"`
	// ... other fields
}
```

## Error Handling

- **Publisher side**: Log errors but don't fail the main operation
- **Subscriber side**: Return errors to trigger retry mechanisms
- Use dead letter queues for persistent failures

## Testing Events

```go
func Test_CreateSubscription_ShouldPublishEvent(t *testing.T) {
	s := setup(t)

	s.repo.On("Create", t.Context(), mock.Anything).Return(nil)
	s.publisher.On("Publish", t.Context(), event.SubscriptionCreatedTopic, mock.Anything).
		Return(nil).
		Once()

	_, err := s.usecase.CreateSubscription(t.Context(), dto.CreateSubscriptionInput{
		UserID: "user-123",
		PlanID: "plan-456",
	})

	require.NoError(t, err)
	s.publisher.AssertExpectations(t)
}
```

## Reference

See [APPLICATION.md](../application.md#messaging) for more details on messaging layer.
