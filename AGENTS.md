# 🧠 Hatch – AI Agent Context

> **Purpose:** Help agents contribute to Hatch effectively — fast, simple, and aligned with its philosophy.

---

## Context

Hatch is a **production-ready Go template** built for **clarity, speed, and modularity**.
Each feature (module) is **self-contained**, **explicitly wired**, and **ready for service extraction**.

The goal: **ship production-quality features fast**, with **clean boundaries** and **minimal nesting**.

---

## Architecture in One Glance

Each module represents a **bounded context** (e.g. `note`, `user`)
and contains everything needed for that feature — domain, use cases, handlers, and repositories.

```
internal/note/
├── domain/                 → Business logic (pure Go)
│   ├── note.go             → Entities & rules
│   └── repository.go       → Repository contract
│
├── usecase/                → Application logic
│   ├── create.go
│   ├── fetch.go
│   ├── archive.go
│   └── usecase.go
│
├── db/repository/postgres/ → Database implementation
│   ├── note_repository.go
│   └── note_query.go
│
├── http/rest/              → HTTP layer
│   ├── handler/
│   │   ├── create_handler.go
│   │   ├── fetch_handler.go
│   │   └── archive_handler.go
│   └── route.go
│
└── module.go               → Dependency wiring for the module
```

---

## Agent Guidelines

### 1. **Be Pragmatic**

* Do only what’s needed to fulfill the request.
* Avoid abstractions or “generic helpers” unless clearly justified.

### 2. **Follow the Pattern**

* **Domain** → Core business logic, pure Go
* **UseCase** → Application orchestration, depends on domain interfaces
* **HTTP / DB** → External adapters implementing domain interfaces
* One handler = one file. One use case = one file.

### 3. **Don’t Over-Engineer**

* No factories, no DI frameworks, no “extra” layers.
* Don’t generalize unless reused 3+ times.

### 4. **Write for Production**

* Prioritize readability, reliability, and testability.
* Avoid clever tricks — **boring Go is good Go**.

### 5. **Test Every Use Case**

* Each `*.go` file in `usecase/` should have a matching test.
* Mock only repository interfaces — test *behavior*, not implementation.

---

## 🚫 Never Do

* ❌ Add new architectural layers or frameworks
* ❌ Bypass use cases (e.g. handler → repository directly)
* ❌ Write SQL or business logic inside use cases
* ❌ Add “helpers” or “utils” without purpose
* ❌ Refactor unrelated code
* ❌ Use global vars or `panic()`
* ❌ Cross-import between modules
* ❌ Hide logic inside handlers or repositories

---

## ✅ Always Do

* ✅ Keep changes local to the module
* ✅ Reuse existing patterns and naming
* ✅ Respect dependency flow (`db/http` → `usecase` → `domain`)
* ✅ Use explicit error handling
* ✅ Keep files small and focused
* ✅ Domain logic stays in `domain/`, queries in `repository/`

---

## Module Structure Rules

**domain/** – Business entities and contracts

* No external dependencies
* Defines repository interfaces

**usecase/** – Application features

* Coordinates domain logic
* Implements business orchestration

**db/repository/** – Data persistence layer

* Implements domain repositories
* SQL and persistence logic live here

**http/rest/** – Transport layer

* Handlers, routes, request/response models

**module.go** – Wires everything together

---

## Reference

* Entry point: `cmd/api/main.go`
* Example module: `internal/note/`
* Shared utilities: `internal/shared/`
* Reusable packages: `pkg/`

---

## Quick Reminder

Hatch is not a framework — it’s a **clear, idiomatic Go foundation**.
When contributing, always ask:

> “What’s the **simplest change** that keeps modules independent and layers decoupled?”