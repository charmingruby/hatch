# Hatch - AI Agent Context

**Hatch** is a production-ready Go template using Clean Architecture with modular design.

📖 **Read first**: [APPLICATION.MD](docs/application.md) | [LAYOUT.MD](docs/layout.md)

## Quick Context

- **Language**: Go
- **Architecture**: Clean Architecture (HTTP/Messaging → Use Case → Repository → Database)
- **DI Framework**: Uber Fx
- **Entry point**: `app/cmd/api/main.go`
- **Reference module**: `app/internal/note/`

## Agent Commands

Detailed step-by-step guides with complete code examples:

- **[Adding New Module](docs/agents/new-module.md)** - Create a new feature module from scratch
- **[Event-Driven Communication](docs/agents/adding-event-driven-communication.md)** - Add async messaging between modules
- **[Third-Party Integration](docs/agents/adding-third-party-integration.md)** - Integrate external services (Stripe, etc)
- **[Modifying Existing Code](docs/agents/modifying-existing-code.md)** - Safely extend existing modules

## Critical Rules (NEVER violate)

### Architectural
- ❌ No layer skipping (handler → repository directly)
- ❌ No business logic in handlers (handlers only parse, validate, delegate)
- ❌ No SQL in use cases (only in `repository/postgres/*_query.go`)
- ❌ No cross-module `internal/` imports (use `shared/` for shared code)

### Go Patterns
- ❌ No global variables for dependencies
- ❌ No `panic()` for error handling (return errors explicitly)
- ❌ No ignoring context cancellation
- ❌ No exporting internal services in barrel files (only `New()` and `Module`)

### Testing
- ✅ **REQUIRED**: Unit tests for ALL use cases (non-negotiable)
- ❌ No use case implementation without corresponding tests
- ❌ No tests without mocks for dependencies
- ❌ No testing implementation details (test behavior, not internals)

## Quick Reference

### File Patterns
- **Model**: `internal/MODULE/model/entity_name.go`
- **DTO**: `internal/MODULE/dto/operation_name_dto.go`
- **Use Case**: `internal/MODULE/usecase/operation_name.go`
- **Use Case Tests**: `internal/MODULE/usecase/operation_name_test.go`
- **Test Setup**: `internal/MODULE/usecase/setup_test.go`
- **Repository**: `internal/MODULE/repository/postgres/entity_repository.go`
- **Queries**: `internal/MODULE/repository/postgres/entity_query.go`
- **Handler**: `internal/MODULE/http/endpoint/operation_endpoint.go`
- **Event**: `internal/MODULE/messaging/event/event_name.go`
- **Subscriber**: `internal/MODULE/messaging/subscriber/on_event_name.go`
- **External**: `internal/MODULE/external/PROVIDER/service_name.go`
- **Barrel**: `internal/MODULE/MODULE.go`

### Key Files
- Entry: [cmd/api/main.go](app/cmd/api/main.go)
- Example module: [internal/note/](app/internal/note/)
- Example barrel: [internal/note/note.go](app/internal/note/note.go)
- Example use case: [internal/note/usecase/create_note.go](app/internal/note/usecase/create_note.go)
- Example use case test: [internal/note/usecase/create_note_test.go](app/internal/note/usecase/create_note_test.go)
- Example repository: [internal/note/repository/postgres/note_repository.go](app/internal/note/repository/postgres/note_repository.go)
- Example handler: [internal/note/http/endpoint/create_note_endpoint.go](app/internal/note/http/endpoint/create_note_endpoint.go)
