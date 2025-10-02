# Adding Third-Party Integration

Guide to integrate external services (payment gateways, APIs, etc.) following the external layer pattern.

## Overview

The external layer provides an abstraction for third-party services, making them swappable and testable. Multiple implementations can exist for the same interface (e.g., Stripe and PayPal for payments).

## Steps

1. Create `MODULE/external/` directory
2. Define interface contract
3. Create provider implementation directory (`external/PROVIDER/`)
4. Implement client
5. Use interface in use cases (dependency injection)
6. Configure provider selection via environment variables

## Complete Example: Payment Gateway Integration

### 1. Interface Contract (internal/billing/external/payment_gateway.go)

Define what operations the external service provides:

```go
package external

import "context"

type PaymentGateway interface {
	ProcessPayment(ctx context.Context, req ProcessPaymentRequest) (PaymentResult, error)
	RefundPayment(ctx context.Context, paymentID string) (RefundResult, error)
	GetPaymentStatus(ctx context.Context, paymentID string) (PaymentStatus, error)
}

type ProcessPaymentRequest struct {
	Amount      int    // Amount in cents
	Currency    string // ISO currency code (USD, EUR, etc)
	CustomerID  string
	Description string
}

type PaymentResult struct {
	PaymentID     string
	Status        string
	ProcessedAt   time.Time
	ReceiptURL    string
}

type RefundResult struct {
	RefundID    string
	Amount      int
	ProcessedAt time.Time
}

type PaymentStatus struct {
	PaymentID string
	Status    string // "pending", "succeeded", "failed"
	Amount    int
}
```

### 2. Stripe Implementation (internal/billing/external/stripe/payment_gateway.go)

```go
package stripe

import (
	"context"
	"time"

	"HATCH_APP/internal/billing/external"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/refund"
)

type StripeClient struct {
	apiKey string
}

func NewStripeClient(apiKey string) *StripeClient {
	stripe.Key = apiKey
	return &StripeClient{apiKey: apiKey}
}

func (s *StripeClient) ProcessPayment(ctx context.Context, req external.ProcessPaymentRequest) (external.PaymentResult, error) {
	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(int64(req.Amount)),
		Currency:    stripe.String(req.Currency),
		Customer:    stripe.String(req.CustomerID),
		Description: stripe.String(req.Description),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return external.PaymentResult{}, err
	}

	return external.PaymentResult{
		PaymentID:   pi.ID,
		Status:      string(pi.Status),
		ProcessedAt: time.Unix(pi.Created, 0),
		ReceiptURL:  pi.Charges.Data[0].ReceiptURL,
	}, nil
}

func (s *StripeClient) RefundPayment(ctx context.Context, paymentID string) (external.RefundResult, error) {
	params := &stripe.RefundParams{
		PaymentIntent: stripe.String(paymentID),
	}

	r, err := refund.New(params)
	if err != nil {
		return external.RefundResult{}, err
	}

	return external.RefundResult{
		RefundID:    r.ID,
		Amount:      int(r.Amount),
		ProcessedAt: time.Unix(r.Created, 0),
	}, nil
}

func (s *StripeClient) GetPaymentStatus(ctx context.Context, paymentID string) (external.PaymentStatus, error) {
	pi, err := paymentintent.Get(paymentID, nil)
	if err != nil {
		return external.PaymentStatus{}, err
	}

	return external.PaymentStatus{
		PaymentID: pi.ID,
		Status:    string(pi.Status),
		Amount:    int(pi.Amount),
	}, nil
}
```

### 3. PayPal Implementation (internal/billing/external/paypal/payment_gateway.go)

Alternative implementation for the same interface:

```go
package paypal

import (
	"context"

	"HATCH_APP/internal/billing/external"
	// PayPal SDK imports
)

type PayPalClient struct {
	clientID     string
	clientSecret string
}

func NewPayPalClient(clientID, clientSecret string) *PayPalClient {
	return &PayPalClient{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (p *PayPalClient) ProcessPayment(ctx context.Context, req external.ProcessPaymentRequest) (external.PaymentResult, error) {
	// PayPal-specific implementation
	// ...
	return external.PaymentResult{}, nil
}

func (p *PayPalClient) RefundPayment(ctx context.Context, paymentID string) (external.RefundResult, error) {
	// PayPal-specific implementation
	// ...
	return external.RefundResult{}, nil
}

func (p *PayPalClient) GetPaymentStatus(ctx context.Context, paymentID string) (external.PaymentStatus, error) {
	// PayPal-specific implementation
	// ...
	return external.PaymentStatus{}, nil
}
```

### 4. Use in Use Case (internal/billing/usecase/process_payment.go)

Use cases depend on the interface, not concrete implementations:

```go
package usecase

import (
	"context"

	"HATCH_APP/internal/billing/dto"
	"HATCH_APP/internal/billing/external"
	"HATCH_APP/internal/billing/repository"
	"HATCH_APP/internal/shared/customerr"
)

type UseCase struct {
	repo           repository.SubscriptionRepository
	paymentGateway external.PaymentGateway  // Interface, not concrete type
}

func New(repo repository.SubscriptionRepository, pg external.PaymentGateway) UseCase {
	return UseCase{
		repo:           repo,
		paymentGateway: pg,
	}
}

func (u UseCase) ProcessPayment(ctx context.Context, input dto.ProcessPaymentInput) (dto.ProcessPaymentOutput, error) {
	// Get subscription
	subscription, err := u.repo.FindByID(ctx, input.SubscriptionID)
	if err != nil {
		return dto.ProcessPaymentOutput{}, customerr.NewDatabaseError(err)
	}

	// Process payment via external service
	result, err := u.paymentGateway.ProcessPayment(ctx, external.ProcessPaymentRequest{
		Amount:      input.Amount,
		Currency:    input.Currency,
		CustomerID:  subscription.UserID,
		Description: "Subscription payment",
	})
	if err != nil {
		return dto.ProcessPaymentOutput{}, customerr.NewExternalServiceError(err)
	}

	// Update subscription with payment result
	// ...

	return dto.ProcessPaymentOutput{
		PaymentID: result.PaymentID,
		Status:    result.Status,
	}, nil
}
```

### 5. Wire in Barrel File (internal/billing/billing.go)

Configure which provider to use based on config:

```go
package billing

import (
	"HATCH_APP/config"
	"HATCH_APP/internal/billing/external"
	"HATCH_APP/internal/billing/external/paypal"
	"HATCH_APP/internal/billing/external/stripe"
	"HATCH_APP/internal/billing/http/endpoint"
	"HATCH_APP/internal/billing/repository/postgres"
	"HATCH_APP/internal/billing/usecase"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func New(log *logger.Logger, r *gin.Engine, db *sqlx.DB, cfg *config.Config) error {
	repo, err := postgres.NewSubscriptionRepo(db)
	if err != nil {
		return err
	}

	// Select payment gateway based on config
	var paymentGateway external.PaymentGateway
	switch cfg.PaymentProvider {
	case "stripe":
		paymentGateway = stripe.NewStripeClient(cfg.StripeAPIKey)
	case "paypal":
		paymentGateway = paypal.NewPayPalClient(cfg.PayPalClientID, cfg.PayPalClientSecret)
	default:
		log.Fatal("invalid payment provider", "provider", cfg.PaymentProvider)
	}

	uc := usecase.New(repo, paymentGateway)
	val := validator.New()

	endpoint.New(log, val, uc).Register(r)
	return nil
}

var Module = fx.Module("billing", fx.Invoke(New))
```

### 6. Configuration (config/config.go)

```go
type Config struct {
	// ... other config

	PaymentProvider    string `env:"PAYMENT_PROVIDER" envDefault:"stripe"`
	StripeAPIKey       string `env:"STRIPE_API_KEY"`
	PayPalClientID     string `env:"PAYPAL_CLIENT_ID"`
	PayPalClientSecret string `env:"PAYPAL_CLIENT_SECRET"`
}
```

### 7. Environment Variables (.env)

```bash
PAYMENT_PROVIDER=stripe
STRIPE_API_KEY=sk_test_xxxxx

# For PayPal:
# PAYMENT_PROVIDER=paypal
# PAYPAL_CLIENT_ID=xxxxx
# PAYPAL_CLIENT_SECRET=xxxxx
```

## Testing External Services

Use mocks for the interface:

```go
func Test_ProcessPayment_ShouldProcessSuccessfully(t *testing.T) {
	s := setup(t)

	s.repo.On("FindByID", t.Context(), "sub-123").
		Return(model.Subscription{UserID: "user-123"}, nil)

	s.paymentGateway.On("ProcessPayment", t.Context(), mock.MatchedBy(func(req external.ProcessPaymentRequest) bool {
		return req.Amount == 1000 && req.Currency == "USD"
	})).
		Return(external.PaymentResult{
			PaymentID: "pay-123",
			Status:    "succeeded",
		}, nil).
		Once()

	output, err := s.usecase.ProcessPayment(t.Context(), dto.ProcessPaymentInput{
		SubscriptionID: "sub-123",
		Amount:         1000,
		Currency:       "USD",
	})

	require.NoError(t, err)
	assert.Equal(t, "pay-123", output.PaymentID)
}
```

## Best Practices

1. **Interface first**: Always define the interface before implementation
2. **Provider-agnostic**: Don't leak provider-specific details into use cases
3. **Error handling**: Wrap external errors in custom error types
4. **Retry logic**: Implement retries for transient failures
5. **Timeouts**: Always use context timeouts for external calls
6. **Logging**: Log all external service calls for debugging
7. **Monitoring**: Track success/failure rates and latencies

## Example Error Handling

```go
func (u UseCase) ProcessPayment(ctx context.Context, input dto.ProcessPaymentInput) (dto.ProcessPaymentOutput, error) {
	// Set timeout for external call
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := u.paymentGateway.ProcessPayment(ctx, external.ProcessPaymentRequest{
		Amount:   input.Amount,
		Currency: input.Currency,
	})
	if err != nil {
		u.log.Error("payment processing failed", "error", err)
		return dto.ProcessPaymentOutput{}, customerr.NewExternalServiceError(err)
	}

	return dto.ProcessPaymentOutput{
		PaymentID: result.PaymentID,
		Status:    result.Status,
	}, nil
}
```

## Reference

See [APPLICATION.md](../application.md#external) for more details on the external layer.
