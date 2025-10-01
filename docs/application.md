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

**Barrel file** (`MODULE.go`): Only public API, initializes dependencies, registers routes.
```go
func New(r *gin.Engine, db *sqlx.DB) error
```

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
4. Wire in barrel file and register in `main.go` via `newmodule.New(router, db)`