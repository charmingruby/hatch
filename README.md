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
└── pkg/
    ├── event/
    └── errs/
```

**Structure overview:**

* **`domain/`** → Entities, value objects, interfaces, and domain rules
* **`feature/`** → Independent use cases (each subfolder = one use case)
* **`infra/`** → Persistence, messaging, or external integrations
* **`mocks/`** → Generated test doubles for interfaces
* **`pkg/`** → Internal cross-cutting utilities and shared abstractions

---

## Principles

* **Modular by default** – Each module can live independently or evolve into a service
* **Feature-driven** – Deliver use cases end-to-end in isolation
* **Dependency inversion** – Domain defines interfaces, infrastructure implements them
* **Explicit over magical** – No hidden frameworks, just clear Go code
* **Intent-revealing structure** – Folder names express business purpose
* **Simplicity first** – Add layers only when complexity demands it

---

**Simple. Explicit. Modular. Production-ready.**