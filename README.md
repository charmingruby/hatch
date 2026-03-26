# Hatch

Feature-driven Go template built on declarative composition. Read a module, see everything it does.

---

## Architecture

Hatch uses **Vertical Slice Architecture**. Each feature is a self-contained unit that owns its domain logic, use cases, and infrastructure — all within the same bounded context.

```
internal/
├── shared/              ← Project-specific code shared across modules
│   ├── auth/            ← Auth middleware, guards
│   ├── events/          ← Domain events
│   └── contract/        ← Cross-module interfaces
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
func (f *Feature) XPTOEndpoint(w http.ResponseWriter, r *http.Request) { ... }
```

**Declarative composition:** The module entry point is a single declaration of all communication the module participates in. Everything the module does is explicitly composed here:

```go
createNoteF := createnote.New(repo)
listNotesF := listnotes.New(repo)

// HTTP
notes.Post("/", createNoteF.CreateNoteEndpoint)
notes.Get("/", listNotesF.ListNotesEndpoint)

// Event listener
bus.On("user.created", createNoteF.OnUserCreated)

// gRPC
gRPCHandler := &gRPCHandler{
	createNote: createNoteF,
	listNotes: listNotesF,
	archiveNote: archiveNoteF,
}
 
pb.RegisterNoteServiceServer(grpcServer, gRPCHandler)

// ...

type gRPCHandler struct {
	pb.UnimplementedNoteServiceServer

	createNote  *createnote.Feature
	listNotes   *listnotes.Feature
	archiveNote *archivenote.Feature
}

func (s *gRPCHandler) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	return s.createNote.CreateNote(ctx, req)
}

func (s *gRPCHandler) ListNotes(ctx context.Context, req *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	return s.listNotes.ListNotes(ctx, req)
}

func (s *gRPCHandler) ArchiveNote(ctx context.Context, req *pb.ArchiveNoteRequest) (*pb.ArchiveNoteResponse, error) {
	return s.archiveNote.ArchiveNote(ctx, req)
}
```

Routes, event listeners, gRPC services — all explicitly wired. No hidden behavior, no implicit registration.

### Shared Library (`pkg/`)

General public library. Modules import from `pkg/` for cross-cutting concerns — never from other modules.

```
pkg/
├── core/          ← Primitives, app errors (apperr/), etc.
├── database/      ← Datasource errors and provider connectors
├── transport/httpx/ ← HTTP server, request parsing, responses
├── o11y/          ← Observability
├── validator/     ← Input validation
└── ...
```

### Shared Kernel (`internal/shared/`)

Project-specific code used by more than one module, but too app-specific for `pkg/` — e.g., an auth middleware, events messages.
### Domain Layer

Each module owns its business domain in `domain/` — entities, repository contracts, business rules. Example: `Note.Archive()` encapsulates the logic of archiving a note. Other modules never import `domain/` directly; they go through the facade.

### Test Utilities

Reusable helpers for consistent testing live under `test/`.


### Cross-Module Communication

Modules should expose a minimal facade only when other bounded contexts truly need access. This facade shields internal complexity and maintains low coupling. When appropriate, prefer asynchronous communication via a queue instead.

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
