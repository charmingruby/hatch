# Hatch

Feature-driven Go template built on declarative composition. Read a module, see everything it does.

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

Each feature is a small, readable unit. A feature defines its dependencies, exposes its behavior via methods, and owns its tests.

```go
// feature.go — Constructor and dependency wiring
type Feature struct { service *Service }
func New(repo domain.NoteRepository) *Feature { ... }

// http.go — Exposes HTTP behavior
func (f *Feature) HTTP(w http.ResponseWriter, r *http.Request) { ... }
```

**Declarative composition:** The module entry point is a single declaration of all communication the module participates in. Everything the module does is explicitly composed here:

```go
createNote := createnote.New(repo)
listNotes := listnotes.New(repo)

// HTTP
notes.Post("/", createNote.HTTP)
notes.Get("/", listNotes.HTTP)

// Event listener
bus.On("user.created", createNote.OnUserCreated)

// gRPC
pb.RegisterNoteServiceServer(grpcServer, createNote.GRPC())
```

Routes, event listeners, gRPC services — all explicitly wired. No hidden behavior, no implicit registration.

### Shared Library (`pkg/`)

Internal library consumed by all modules. Modules import from `pkg/` for cross-cutting concerns — never from other modules. Organize packages as the project grows.

```
pkg/
├── core/          ← Primitives, app errors (apperr/), etc.
├── database/      ← Datasource errors and provider connectors
├── transport/httpx/ ← HTTP server, request parsing, responses
├── o11y/          ← Observability
├── validator/     ← Input validation
└── ...
```

### Test Utilities

Reusable helpers for consistent testing live under `test/`.

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
