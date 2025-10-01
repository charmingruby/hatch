# Application

## Getting Started

### Project Setup

**1. Replace Module Path**

Search and replace all occurrences of the placeholder:

```bash
HATCH_APP
```

With your actual Go module path (e.g., `github.com/yourorg/yourproject`).

**2. Update Docker Image**

- Change the image name in `.github/workflows/*.yml`
- Configure repository secrets: `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN`

**3. Environment Variables**

- Copy `.env.example` to `.env`
- Fill in required values for local development
- Never commit `.env` to version control

## Module Architecture

### Module Structure

The `internal/` directory contains the example module `note/`. Use it as a reference for organizing your own feature modules.

**Standard module layout:**

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

### Barrel File Pattern

Each module has a **barrel file** (`MODULE.go`) at its root. This file:

- Is the **only public API** of the module
- Initializes and wires all module dependencies
- Registers HTTP routes or message subscribers
- Returns any errors during initialization

**Example signature:**

```go
func New(r *gin.Engine, db *sqlx.DB) error
```

The barrel file encapsulates complexity and provides a clean integration point for `main.go`.

## Code Organization Principles

### Internal vs Shared vs Pkg

**`internal/MODULE/`** - Feature modules (private to this service)
- Self-contained business capabilities
- Each module is independent
- Modules should not import each other directly

**`internal/shared/`** - Cross-cutting concerns (private to this repository)
- Used by multiple modules but not meant for external use
- Examples: HTTP response helpers, custom error types, common middleware

**`pkg/`** - Reusable libraries (public API, can be imported by other projects)
- Examples: logger wrapper, database client, ID generation utilities
- Should have no dependencies on `internal/`

### External Service Integration

For third-party APIs (Stripe, SendGrid, AWS, etc.):

**Within a single module:**
```
internal/
  billing/
    external/
      payment_gateway.go    # Interface (contract)
      stripe/
        gateway.go          # Stripe implementation
      paypal/
        gateway.go          # PayPal implementation
```

**Shared across modules:**
```
internal/
  shared/
    storage/
      uploader.go           # S3Uploader interface
      s3/
        client.go           # AWS S3 implementation
```

**Rule:** Business logic depends on **interfaces**, never on concrete implementations.

## Communication Layers

### HTTP (Synchronous)

**Supported protocols:**
- `http/rest/` - REST APIs using Gin
- `http/grpc/` - gRPC services (structure ready, implement as needed)
- `http/gql/` - GraphQL APIs (structure ready, implement as needed)

**HTTP Handler Responsibilities:**

1. Parse and validate incoming request
2. Call use case with DTO input
3. Map domain errors to HTTP status codes
4. Return standardized responses

### Messaging (Asynchronous)

For event-driven architectures using NATS, Kafka, RabbitMQ, etc.

**Module-specific events:**

```
internal/
  order/
    messaging/
      event/
        order_placed.go       # Event definition + serialization
      subscriber/
        on_order_placed.go    # Event handler
```

**Naming convention:** Event `order.placed` → Handler `onOrderPlaced()`

**Shared messaging infrastructure:**

- `pkg/messaging/` - Contracts and connections
- `internal/shared/messaging/` - Common messages

Implementations (NATS, Kafka adapters) live alongside the interface they implement.

## Data Transfer Objects (DTOs)

DTOs define use case contracts and decouple application layers. They are also used by delivery handlers for validation, ensuring a single, consistent contract and minimizing the number of places that need changes.

**Location:** `internal/MODULE/dto/`

**Naming:** `operation_name_dto.go` (e.g., `create_note_dto.go`)

**Structure:**

```go
type CreateNoteInput struct {
    Title   string `json:"title" binding:"required,min=1"`
    Content string `json:"content" binding:"required"`
}

type CreateNoteOutput struct {
    ID string `json:"id"`
}
```

**Benefits:**
- Changing a use case signature doesn't require updating handlers, repositories, etc.
- Clear contract between layers
- Easier to test and document

## Testing

### Test Organization

- **Unit tests**: Live next to the code (`*_test.go` in the same package)
- **Mocks**: Generated via `mockery` in `test/gen/MODULE/mocks/`
- **Test setup helpers**: Use `setup_test.go` for common test fixtures

### Test Naming Convention

Follow the pattern:

```
<action> <expected result> when <condition>
```

**Examples:**

```go
func TestCreateNote_ShouldReturnNote_WhenInputIsValid(t *testing.T)
func TestArchiveNote_ShouldReturnNotFoundError_WhenNoteDoesNotExist(t *testing.T)
func TestListNotes_ShouldReturnEmptyList_WhenNoNotesExist(t *testing.T)
```

This format makes test intent immediately clear.

### Makefile Commands

- `make test` - Run all tests
- `make test-coverage` - Generate and view coverage report
- `make mock` - Regenerate mocks using mockery

## Database & Persistence

### PostgreSQL Setup

- **Driver:** `sqlx` (enhanced SQL operations)
- **Migrations:** SQL files in `db/migrations/`
- **Connection:** Managed via `pkg/database/postgres/`

### Repository Pattern

**Interface location:** `internal/MODULE/repository/repository.go`

**Implementation:** `internal/MODULE/repository/postgres/MODULE_repository.go`

**SQL Queries:** Separate into `*_query.go` files for readability

**Example:**

```
internal/note/repository/
  repository.go              # NoteRepository interface
  postgres/
    note_repository.go       # Implementation
    note_query.go            # SQL queries as constants
```

**Benefits:**
- Easy to swap implementations (e.g., add a Redis cache)
- SQL is isolated and easier to review/optimize
- Repositories are mockable for testing

## Go Best Practices

### Error Handling

- **Always return errors explicitly** - Don't use `panic()` in library code
- **Use custom error types** for domain errors (`internal/shared/customerr/`)
- **Wrap errors with context:** `fmt.Errorf("create note: %w", err)`
- **Check errors at layer boundaries** (handlers, repositories, external services)

### Context Usage

- **Always accept `context.Context` as the first parameter** in use cases and repositories
- **Propagate context through layers:** Handler → Use Case → Repository
- Use context for timeouts, cancellation, and request-scoped values (e.g., trace IDs)

### Naming Conventions

- **Interfaces:** Named after behavior (`NoteRepository`, `Service`, `Uploader`)
- **Packages:** Short, lowercase, single word (`note`, `usecase`, `postgres`)
- **Files:** snake_case (`create_note_dto.go`, `note_repository.go`)
- **Structs:** PascalCase for exported, camelCase for unexported

### Dependency Injection

- **Constructor pattern:** Each component has a `New()` function
- **Explicit dependencies:** Pass interfaces as constructor parameters
- **No global state:** Avoid global variables for dependencies
- **Wire at startup:** `main.go` and barrel files handle all wiring

### Struct Initialization

- Use field names in struct literals: `Config{Port: ":8080", DB: dbConn}`
- Prefer constructors for complex initialization
- Return pointers for large structs or when mutability is expected

## Development Workflow

### Local Development

```bash
make up          # Start Docker Compose (PostgreSQL + API)
make down        # Stop and remove containers
make logs        # View application logs
make test        # Run tests
make mock        # Generate mocks
```

### Database Migrations

```bash
make migrate-up      # Apply all pending migrations
make migrate-down    # Rollback last migration
```

Migrations are in `db/migrations/` and follow the naming pattern generated by `golang-migrate`:
```
001_create_notes_table.up.sql
001_create_notes_table.down.sql
```

### Adding a New Module

1. Create directory: `internal/newmodule/`
2. Add barrel file: `newmodule.go` with `New()` function
3. Create subdirectories: `model/`, `dto/`, `usecase/`, `repository/`, `http/`
4. Implement domain model in `model/`
5. Define DTOs in `dto/`
6. Create use case interface in `usecase/usecase.go`
7. Implement business logic in `usecase/*.go`
8. Create repository interface and PostgreSQL implementation
9. Build HTTP handlers in `http/endpoint/`
10. Wire everything in the barrel file
11. Register module in `main.go` by calling `newmodule.New(router, db)`

---