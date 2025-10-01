**# Application

## Quick Start

1. Replace `HATCH_APP` with your Go module path (e.g., `github.com/yourorg/project`)
2. Update Docker image name in `.github/workflows/*.yml` and configure secrets
3. Copy `.env.example` → `.env` and fill required values

## Dependency Injection

Hatch uses **[Uber Fx](https://uber-go.github.io/fx/)** for dependency injection and application lifecycle management.

### Key Concepts

- **Modules** - Self-contained units that provide and/or invoke dependencies
- **Providers** (`fx.Provide`) - Functions that create and return dependencies
- **Invocations** (`fx.Invoke`) - Functions that consume dependencies to perform initialization
- **Lifecycle Hooks** - Startup and shutdown hooks for managing resources

### Module Types

**Provider Modules** - Export dependencies for other modules to consume (e.g., `config.Module`, `postgres.Module`)
**Invoke Modules** - Consume dependencies to perform initialization (e.g., `note.Module`, `health.Module`)
**Lifecycle Modules** - Manage resource lifecycle with startup/shutdown hooks (e.g., `rest.Module`, `postgres.Module`)

See [cmd/api/main.go](../app/cmd/api/main.go), [config/config.go](../app/config/config.go), [internal/note/note.go](../app/internal/note/note.go), and [pkg/database/postgres/postgres.go](../app/pkg/database/postgres/postgres.go) for reference implementations.

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

**Barrel file** (`MODULE.go`): Exports an Fx module with dependency injection:
- Constructor function (`New`) receives dependencies via Fx
- Wires internal components (repository → use case → endpoints)
- Registers HTTP routes
- Exports `Module` variable for main.go registration

### Directory Principles

- **`internal/MODULE/`** - Feature modules (private, independent)
- **`internal/shared/`** - Cross-cutting concerns (middleware, errors, helpers)
- **`pkg/`** - Reusable libraries (public API, no `internal/` deps)

**External services:** Use interfaces. Place in `MODULE/external/` (single module) or `shared/SERVICE/` (multiple modules).

## Layers

### HTTP
- REST (`http/rest/`), gRPC, GraphQL supported
- Handlers: parse → call use case → map errors → respond

### Messaging
- Event-driven via NATS/Kafka/RabbitMQ
- Events: `MODULE/messaging/event/`
- Handlers: `MODULE/messaging/subscriber/`

### DTOs
- Location: `internal/MODULE/dto/operation_name_dto.go`
- Define contracts between layers
- Used for validation in handlers

```go
type CreateNoteInput struct {
    Title string `json:"title" binding:"required"`
}
```

### Repository
- Interface: `repository/repository.go`
- Implementation: `repository/postgres/MODULE_repository.go`
- SQL queries: Separate `*_query.go` files

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
5. Register module in `cmd/api/main.go` by adding `newmodule.Module` to `fx.New()`**