# Hatch

A production-ready Go template for building modular, feature-driven applications with explicit, declarative behaviors.

---

## Why Hatch?

- **Fast iteration** — Deliver features independently
- **Declarative features** — Each feature exposes its behavior explicitly
- **Go-idiomatic** — No magic, just clear Go code

---

## Architecture

Hatch uses **Vertical Slice Architecture**. Each feature is a self-contained unit that owns its domain logic, use cases, and infrastructure — all within the same bounded context.

```
internal/
└── note/
    ├── note.go              ← Module entry: declares all transports, listeners, and public contracts
    ├── domain/              ← Entities, value objects, repository contracts
    ├── feature/             ← One folder per use case
    │   ├── createnote/
    │   ├── archivenote/
    │   └── listnotes/
    ├── infra/               ← Persistence, external integrations
    └── mocks/               ← Test doubles for interfaces
```

### Features

Each feature folder is a small, readable unit. A feature defines its dependencies, exposes its behavior via methods, and owns its tests.

```go
// feature.go — Constructor and dependency wiring
type Feature struct { service *Service }
func New(repo domain.NoteRepository) *Feature { ... }

// http.go — Exposes HTTP behavior
func (f *Feature) HTTP(w http.ResponseWriter, r *http.Request) { ... }
```

The module entry point is a single declaration of all communication the module participates in — HTTP routes, event listeners, public API contracts, and anything else:

```go
createNote := createnote.New(repo)
notes.Post("/", createNote.HTTP)
```

### Shared Packages (`pkg/`)

An internal library with common functionalities shared across modules. Flat and provider-agnostic:

```
pkg/
├── core/          ← Common primitives (ID generation, etc.)
├── database/      ← Datasource errors and provider connectors
├── transport/httpx/ ← HTTP server, request parsing, responses
├── o11y/          ← Observability (structured logging)
└── validator/     ← Input validation
```

### Test Infrastructure

Reusable helpers for consistent testing live under `test/`.

---

## Principles

1. **Modules own their boundaries** — Each module is self-contained with explicit dependencies
2. **Features drive delivery** — Ship use cases end-to-end in isolation
3. **Domain defines contracts** — Interfaces live in the domain, infrastructure implements them
4. **No magic** — Everything is readable Go; no frameworks, no reflection tricks
5. **Declarative wiring** — Features expose methods; modules compose them explicitly
6. **Flat and discoverable** — Structure expresses intent; avoid nesting unless it clarifies
7. **Test what matters** — Each feature owns its tests with real dependencies where practical

---

## Cross-Module Communication

Modules expose a minimal facade when other bounded contexts need access. The facade protects internals and keeps coupling low.

```
internal/
├── note/
│   ├── note.go          ← Public facade
│   ├── domain/
│   ├── feature/
│   └── infra/
```

- **Minimal surface** — Export only what's strictly needed
- **Stable contracts** — The facade shields internal changes
- **Explicit dependencies** — Consumers depend on the facade, not on internals

---

**Simple. Explicit. Modular. Production-ready.**
