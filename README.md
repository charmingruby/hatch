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
