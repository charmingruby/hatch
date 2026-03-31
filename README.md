# Hatch

Feature-driven Go template built on declarative composition. Read a module, see everything it does.

---

## Architecture

**Vertical Slice Architecture** — each feature is a self-contained unit that owns its domain logic, use cases, and infrastructure within the same bounded context.

```
internal/
├── shared/              ← Project-specific code shared across modules
│   ├── auth/            ← Auth middleware, guards
│   └── events/          ← Events
└── note/                ← Example module
    ├── note.go          ← Module entry: transports, listeners, public contracts
    ├── domain/          ← Entities, value objects, repository contracts
    ├── feature/         ← One folder per use case
    ├── infra/           ← Persistence, external integrations
    └── mocks/           ← Test doubles
```

---

## Features

Each feature defines its dependencies, exposes behavior via methods, and owns its tests.

```
feature/createnote/
├── feature.go   ← Constructor + dependency wiring
├── service.go   ← Domain logic
├── http.go      ← HTTP handler
├── event.go     ← Event handlers
└── grpc.go      ← gRPC handlers
```

```go
// feature.go
type Feature struct { service *Service }
func New(repo domain.NoteRepository) *Feature { ... }

// http.go
func (f *Feature) CreateNoteEndpoint(w http.ResponseWriter, r *http.Request) { ... }
```

### Declarative Composition

The module entry point declares all communication explicitly — no hidden behavior:

```go
createNoteF := createnote.New(repo)
listNotesF := listnotes.New(repo)

// HTTP
r.Route("/api/v1/notes", func(r chi.Router) {
    r.Post("/", createNoteF.CreateNoteEndpoint)
    r.Get("/", listNotesF.ListNotesEndpoint)
})

// Events
bus.On("user.created", createNoteF.OnUserCreated)

// gRPC
pb.RegisterNoteServiceServer(grpcServer, &gRPCHandler{...})
```

---

## Shared Library (`pkg/`)

Modules import from `pkg/` for cross-cutting concerns — never from other modules. Organized by **capabilities**, not technology.

```
pkg/
├── core/              ← Primitives, app errors
├── resource/          ← Shared connections (redis/, postgres/)
├── cache/             ← Capability: caching (redis/)
├── lock/              ← Capability: distributed locking (redis/, postgres/)
├── store/             ← Capability: data persistence (postgres/)
├── transport/
│   ├── httpx/         ← HTTP transport
│   └── messagebus/    ← Messaging (rabbitmq/)
├── o11y/              ← Observability
└── validator/         ← Input validation
```

Connections created once in `resource/`, injected into capabilities:

```go
redisPool := resource.NewRedisPool(config)
cache := cache.NewRedis(redisPool)
locker := lock.NewRedis(redisPool)
```

---

## Cross-Module Communication

Modules expose a **minimal facade** only when other bounded contexts need access.

- **Minimal surface** — export only what's strictly needed
- **Stable contracts** — facade shields internal changes
- **Explicit dependencies** — consumers depend on facade, not internals
- **Prefer async** — use queues when possible

---

## Testing

Each feature owns its tests. Use `mocks/` for test doubles, helpers in `test/`.

---

**Simple. Explicit. Modular. Production-ready.**
