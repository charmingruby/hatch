# Application

## Quick Start

1. Replace `HATCH_APP` with your Go module path (e.g., `github.com/yourorg/project`)
2. Update Docker image name in `.github/workflows/*.yml` and configure secrets
3. Copy `.env.example` → `.env` and fill required values

## Architecture

### Module Structure

Standard layout for feature modules in `internal/MODULE/`:

```
internal/
  MODULE/
    MODULE.go            # Barrel file - wires all dependencies
    model/               # Domain entities (business objects)
    dto/                 # Data Transfer Objects (use case contracts)
    usecase/             # Business logic layer
      usecase.go         # Service interface definition
      *.go               # Use case implementations
      *_test.go          # Unit tests
    repository/          # Data access layer
      repository.go      # Repository interface
      postgres/          # PostgreSQL implementation
        *_repository.go
        *_query.go       # SQL queries separated from logic
    http/                # Presentation layer
      endpoint/
        endpoint.go      # Handler registration
        *.go             # HTTP handlers
    messaging/           # Event-driven communication (optional)
      event/             # Event definitions
      subscriber/        # Event handlers
    external/            # Third-party service integrations (optional)
```

The `internal/note/` module is fully implemented and serves as a reference example.

### Dependency Injection

Hatch uses **[Uber Fx](https://uber-go.github.io/fx/)** for dependency injection and application lifecycle management.

#### Key Concepts
- **Modules** - Self-contained units that provide and/or invoke dependencies
- **Providers** (`fx.Provide`) - Functions that create and return dependencies
- **Invocations** (`fx.Invoke`) - Functions that consume dependencies to perform initialization
- **Lifecycle Hooks** - Startup and shutdown hooks for managing resources

#### Module Types
**Provider Modules** - Export dependencies for other modules to consume (e.g., `config.Module`, `postgres.Module`)
**Invoke Modules** - Consume dependencies to perform initialization (e.g., `note.Module`, `health.Module`)
**Lifecycle Modules** - Manage resource lifecycle with startup/shutdown hooks (e.g., `rest.Module`, `postgres.Module`)

See [cmd/api/main.go](../app/cmd/api/main.go), [config/config.go](../app/config/config.go), [internal/note/note.go](../app/internal/note/note.go), and [pkg/database/postgres/postgres.go](../app/pkg/database/postgres/postgres.go) for reference implementations.

### Directory Principles

- **`internal/MODULE/`** - Feature modules (private, independent)
- **`internal/shared/`** - Cross-cutting concerns (middleware, errors, helpers)
- **`pkg/`** - Reusable libraries (public API, no `internal/` deps)

## Layers

Request flow (outside → inside):

### HTTP
Entry point for synchronous requests.
- Handlers parse input, call use cases, map errors, and respond
- Supports REST, gRPC, GraphQL

**Implementation(REST example):**
- Routes: `MODULE/http/endpoint/endpoint.go`
- Handlers: `MODULE/http/endpoint/*.go`

### Messaging
Entry point for asynchronous events.
- Event-driven communication via NATS/Kafka/RabbitMQ
- Subscribers handle events and trigger use cases

**Implementation:**
- Events: `MODULE/messaging/event/*.go` (parsing and mapping logic, e.g., protobuf → model)
- Subscribers: `MODULE/messaging/subscriber/*.go` (consume events, invoke use cases)

### DTO
Contracts between presentation and business layers.
- Define input/output structures
- Enable validation at boundaries

**Implementation:**
- Location: `MODULE/dto/operation_name_dto.go`

### Use Case
Business logic and orchestration.
- Coordinates repositories, external services, and domain models
- Implements application workflows
- Returns domain entities or DTOs

**Implementation:**
- Interface: `MODULE/usecase/usecase.go`
- Implementations: `MODULE/usecase/*.go`
- Tests: `MODULE/usecase/*_test.go`

### Repository
Data access abstraction.
- Interfaces hide implementation details
- Implementations handle database operations
- Queries separated for clarity

**Implementation:**
- Interface: `MODULE/repository/repository.go`
- PostgreSQL: `MODULE/repository/postgres/MODULE_repository.go`
- SQL queries: `MODULE/repository/postgres/*_query.go`

### External
Third-party service integrations.
- Defines contracts for external APIs and services
- Multiple implementations per contract (e.g., Stripe, PayPal for payment)
- Called by use cases when external communication is needed

**Implementation:**
- Contract: `MODULE/external/service_name.go`
- Implementations: `MODULE/external/PROVIDER/client.go`

Example:
```
internal/
  billing/
    external/
      payment_gateway.go    # Interface contract
      stripe/
        payment_gateway.go          # Stripe implementation
      paypal/
        payment_gateway.go          # PayPal implementation
```

## Testing

- Tests: `*_test.go` alongside code
- Naming: `Test<Action>_Should<Result>_When<Condition>`
- Mocks: Generated via `mockery` in `test/gen/MODULE/mocks/`
- Commands: `make test`, `make test-coverage`, `make mock`

## Best Practices

### Go Standards
- **Errors:** Return explicitly, use custom types, wrap with context
- **Context:** First param in use cases/repos, propagate through layers
- **Naming:** Interfaces by behavior, packages lowercase single-word, files snake_case
- **DI:** Constructor pattern, explicit interfaces, no globals, wire at startup

### Database
- Driver: `sqlx`
- Migrations: `db/migrations/` (`000001_creates_notes_table.up.sql`, `000001_creates_notes_table.down.sql`)
- Commands: `make migrate-up`, `make migrate-down`

### Adding a Module
1. Create `internal/newmodule/` with barrel file `newmodule.go`
2. Add `model/`, `dto/`, `usecase/`, `repository/`, `http/` subdirs
3. Implement layers: model → DTO → use case → repository → handlers
4. Create Fx module in barrel file:
   ```go
   func New(log *logger.Logger, r *gin.Engine, db *sqlx.DB) error {
       // Wire dependencies
       return nil
   }

   var Module = fx.Module("newmodule", fx.Invoke(New))
   ```
5. Register module in `cmd/api/main.go` by adding `newmodule.Module` to `fx.New()`
