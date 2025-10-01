# Hatch - AI Agent Context

## Project Overview

**Hatch** is a production-ready Go template implementing Clean Architecture with modular design and Go best practices.

### Architecture
- **Domain**: Entities (`internal/*/model/`)
- **Use Case**: Business logic (`internal/*/usecase/`)
- **Interface Adapters**: Handlers, repositories (`internal/*/http/`, `internal/*/repository/`)
- **External**: Third-party integrations (`internal/*/external/`)

### Principles
- Modular, interface-driven, explicit contracts via DTOs
- Fail-fast validation, dependency inversion

## Directory Structure

```
├── app/                     # Main application
│   ├── cmd/api/             # Entry point
│   ├── config/               # Configuration
│   ├── internal/            # Private code
│   │   ├── MODULE/          # Feature module (note, user, billing)
│   │   │   ├── model/       # Domain entities
│   │   │   ├── dto/         # Shared contract between use case and delivery Layer I/O
│   │   │   ├── usecase/     # Business logic
│   │   │   ├── repository/  # Data access
│   │   │   ├── http/        # HTTP handlers
│   │   │   ├── messaging/   # Events
│   │   │   └── external/    # Third-party contracts
│   │   └── shared/          # Cross-cutting concerns
│   ├── pkg/                 # Public libraries
│   ├── test/                # Mocks & test utils
│   └── db/                  # Migrations
├── apps/                    # Monorepo 
└── infra/                   # Infrastructure (k8s, helm, terraform) 
```

### Bootstrap
1. Replace `HATCH_APP` with your Go module path
2. Update Docker images in **workflows**
3. Configure CI/CD secrets
4. Replace `internal/note` with your features

## Module Structure

```go
internal/note/
  note.go               # Barrel file: wires deps, public API
  model/note.go         # Domain entity
  dto/*.go              # Input/Output contracts
  usecase/
    usecase.go          # Service interface
    create_note.go      # Implementation
  repository/
    repository.go       # Interface
    postgres/note_repository.go
  http/endpoint/*.go    # HTTP handlers
```

**Dependency Flow:**
```
main.go → module.New(router, db)
  → creates repository
  → creates use case
  → creates endpoints
  → registers routes
```

**Shared Code:**
- `internal/shared/` - HTTP utils, errors, storage, messaging
- `pkg/` - Logger, database, validator, ID gen

**External Services:**
```go
internal/billing/external/
  payment_gateway.go         // Interface
  stripe/payment_gateway.go  // Implementation
```

## Patterns

### HTTP Layer
- REST APIs: `http/rest/`
- gRPC Services: `http/grpc/`
- GraphQL APIs: `http/gql/`

### Messaging
- Events in `MODULE/messaging/event/`
- Handlers in `MODULE/messaging/subscriber/`
- Convention: `order.created` → `onOrderCreated()`

### Database
- PostgreSQL with sqlx
- Migrations in `app/db/migrations/`
- Repository pattern: interface + postgres implementation
- SQL queries in `*_query.go` files

### Testing
- Unit tests next to code
- Mocks via mockery in `test/gen/MODULE/mocks/`
- Naming: `should <result> when <condition>`
- Commands: `make test`, `make test-coverage`, `make mock`

### Error Handling
- Return errors, never panic
- Custom domain errors in `internal/shared/customerr/`
- Wrap with context: `fmt.Errorf("msg: %w", err)`

## Go Conventions

### Naming
- **Interfaces**: `NoteRepository`, `Service`
- **Files**: snake_case (`create_note_dto.go`)
- **Packages**: lowercase, single word (`note`, `usecase`)

### Dependency Injection
- Constructor pattern: `New()` functions
- Pass interfaces as parameters
- Wire in `main.go` and barrel files

### Context
- First parameter in use cases
- Propagate through layers

## AI Agent Guidelines

### Adding New Module
1. Create `internal/MODULE/`
2. Add barrel file: `MODULE.go` with `New()`
3. Define model, DTOs, use case interface
4. Implement use cases, repository, endpoints
5. Wire in barrel file
6. Register in `main.go`

### Modifying Code
- Preserve interface signatures
- Follow existing patterns
- Update tests
- Respect layer boundaries

### Adding Dependencies
- Third-party in `external/` or `pkg/`
- Interface first, implementation second
- Use dependency injection

## Quick Reference

### File Locations
- Entry: `app/cmd/api/main.go`
- Config: `app/config/config.go`
- Example: `app/internal/note/`
- Shared: `app/internal/shared/`
- Packages: `app/pkg/`
- Migrations: `app/db/migrations/`
- Mocks: `app/test/gen/*/mocks/`

### What NOT to Do
- ❌ Global variables for dependencies
- ❌ Skip layers (handler → repository)
- ❌ Business logic in handlers
- ❌ `panic()` for errors
- ❌ Ignore context cancellation
- ❌ SQL in use cases
- ❌ Cross-module `internal/` imports (use `shared/`) 