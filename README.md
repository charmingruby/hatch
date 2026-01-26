# Hatch

Hatch is a pragmatic Go project template designed for rapid feature development and effortless service extraction. Modular, decoupled, and production-ready.

---

## Why Hatch?

* ✅ **Service-ready** – Extract modules into microservices without rewrites
* ✅ **Fast iteration** – Deliver features independently and incrementally
* ✅ **Go-idiomatic** – Simple, explicit, and dependency-free

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
│   │   │   ├── archive_note_handler.go
│   │   │   ├── archive_note_service.go
│   │   │   ├── archive_note_service_test.go
│   │   │   └── archive_note.go
│   │   ├── createnote/
│   │   │   ├── create_note_handler.go
│   │   │   ├── create_note_service.go
│   │   │   ├── create_note_service_test.go
│   │   │   └── create_note.go
│   │   └── listnotes/
│   │       ├── list_notes_handler.go
│   │       ├── list_notes_service.go
│   │       ├── list_notes_service_test.go
│   │       └── list_notes.go
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

* **`register.go`** → Defines module behavior and interface layer (HTTP, gRPC, CLI, etc)
* **`domain/`** → Entities, value objects, interfaces, and domain rules
* **`feature/`** → Independent use cases (each subfolder = one use case)
* **`infra/`** → Persistence, messaging, or external integrations
* **`mocks/`** → Generated test doubles for interfaces

---

## Principles

* **Modular by default** – Each module can live independently or evolve into a service
* **Feature-driven** – Deliver use cases end-to-end in isolation
* **Dependency inversion** – Domain defines interfaces, infrastructure implements them
* **Explicit over magical** – No hidden frameworks, just clear Go code
* **Intent-revealing structure** – Folder names express business purpose
* **Simplicity first** – Add layers only when complexity demands it

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

* **Minimal surface** – Export only what other modules actually need
* **Stable contracts** – The facade acts as a stable boundary, protecting internal changes
* **Explicit dependencies** – Consumers depend on the facade, not on internal implementation

---

**Simple. Explicit. Modular. Production-ready.**
