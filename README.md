# Hatch

Hatch is a pragmatic Go project template designed for rapid feature development and effortless service extraction. Modular, decoupled, and production-ready.

---

## Why Hatch?

- ✅ **Service-ready** – Extract modules into microservices without rewrites
- ✅ **Fast iteration** – Deliver features independently and incrementally
- ✅ **Go-idiomatic** – Simple, explicit, and dependency-free

> Hatch ships a complete vertical-slice stack meant to scale products and complex contexts with clear boundaries and opinionated workflows. If your problem is simpler, doesn’t demand this much ceremony, or you prefer to dial structure in later, consider [F-Hatch](https://github.com/charmingruby/f-hatch) for a flatter starter.

---

## Architecture

Hatch follows the Vertical Slice Architecture, where each feature is a self-contained unit that owns its domain logic, use cases, and infrastructure.

This approach enables fast delivery, clear boundaries, and effortless scalability as the domain evolves.

### Module Organization

Each module lives under `internal/` and follows a consistent vertical-slice layout:
domain logic, features, and infrastructure are grouped inside the same bounded context.

```text
internal/
├── note/
│   ├── register.go          ← Defines module behavior and interface layer
│   ├── domain/
│   │   ├── note_repository.go
│   │   └── note.go
│   ├── feature/
│   │   ├── archivenote/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── service_test.go
│   │   │   └── feature.go
│   │   ├── createnote/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── service_test.go
│   │   │   └── feature.go
│   │   └── listnotes/
│   │       ├── handler.go
│   │       ├── service.go
│   │       ├── service_test.go
│   │       └── feature.go
│   ├── infra/
│   │   └── db/
│   │       └── postgres/
│   │           ├── note_query.go
│   │           └── note_repository.go
│   ├── mocks/
│   │   ├── NoteRepository.go
│   │   └── UseCase.go
│   └── note.go
```

**Structure overview:**

- **`{MODULE_NAME}.go`** → Defines module behavior and interface layer (HTTP, gRPC, CLI, etc)
- **`domain/`** → Entities, value objects, interfaces, and domain rules
- **`feature/`** → Independent use cases (each subfolder = one use case; files inside use role-based names like `handler.go`, `service.go`, `service_test.go`, and `feature.go`)
- **`infra/`** → Persistence, messaging, or external integrations
- **`mocks/`** → Generated test doubles for interfaces

### Declarative features

Think of each feature folder as a small declaration of behavior. It states which dependencies it needs, which transports it exposes, and which background jobs or listeners it runs. Everything stays in plain Go so you can read a feature and instantly know what it does.

**Route-only feature (pure HTTP):** mirrors `internal/note/feature/createnote/feature.go`. The function name is the declaration: “This feature provides an HTTP handler for creating notes.”

```go
func Route(repo domain.NoteRepository) http.HandlerFunc {
    service := NewService(repo)
    handler := NewHandler(service)
    return handler
}
```

From module code you can read `notes.Post("/", createnote.Route(repo))` and understand the behavior without opening other files.

**Feature exporting multiple behaviors:** when the same use case must also emit/consume events, cron jobs, etc., just expose more functions. Each export describes one capability, and everything still returns native Go types.

```go
// Declarative: "this feature serves HTTP"
func Route(repo domain.NoteRepository, bus eventbus.Bus) http.HandlerFunc {
    svc := NewService(repo, bus)
    return NewHandler(svc)
}

// Declarative: "this feature listens to events"
func Listener(repo domain.NoteRepository, bus eventbus.Bus) eventbus.Listener {
    svc := NewService(repo, bus)
    return eventbus.Listener{
        Event: "note.created",
        Handle: func(ctx context.Context, payload any) error {
            return svc.HandleEvent(ctx, payload)
        },
    }
}
```

`internal/{module}/module.go` stays simple and declarative: call whichever functions a feature exposes, mount routes on the HTTP router, and register listeners on the event bus. There is no framework magic—just reading Go functions to understand the module’s behavior map.

### Shared Packages

Infrastructure helpers live under `pkg/`. The database helpers are intentionally flat so each provider ships in a single Go file that owns its own connection logic and error values:

```text
pkg/
├── database/
│   ├── error.go           ← Common datasource errors
│   └── postgres.go        ← Provider-specific connect helpers + sentinels
├── http/
├── id/
├── o11y/
└── validator/
```

Add new providers by dropping another `<provider>.go` beside `postgres.go`; avoid nested directories so all database concerns stay discoverable at a glance.

---

## Principles

- **Modular by default** – Each module can live independently or evolve into a service
- **Feature-driven** – Deliver use cases end-to-end in isolation
- **Dependency inversion** – Domain defines interfaces, infrastructure implements them
- **Explicit over magical** – No hidden frameworks, just clear Go code
- **Intent-revealing structure** – Folder names express business purpose
- **Simplicity first** – Add layers only when complexity demands it

---

## Cross-Module Communication

When a bounded context needs to expose functionality to other modules, the simplest approach is to create a minimal public API (facade) that exports only what is strictly necessary.

```text
internal/
├── note/
│   ├── note.go                ← Public facade: exports minimal API for other modules
│   ├── register.go
│   ├── domain/
│   ├── feature/
│   └── infra/
```

**Guidelines:**

- **Minimal surface** – Export only what other modules actually need
- **Stable contracts** – The facade acts as a stable boundary, protecting internal changes
- **Explicit dependencies** – Consumers depend on the facade, not on internal implementation

---

**Simple. Explicit. Modular. Production-ready.**
