# Hatch – AI Agent Context

> **Purpose:** Help agents contribute to Hatch effectively — fast, simple, and aligned with its philosophy.

---

## Context

Hatch is a **pragmatic Go template** built for **clarity and speed**.
It uses a **modular architecture** where each module is **self-contained and service-ready**.

The goal: **ship production-quality features fast**, with clean separation and easy extraction to microservices.

---

## Architecture in One Glance

Each module represents a **bounded context** (e.g., `note`, `user`)
and is fully independent with clear layer separation.

```
internal/note/
├── domain/              → Business logic (zero dependencies)
│   ├── note.go          → Entities & rules
│   └── repository.go    → Repository interface
├── usecase/             → Application logic
│   ├── create.go        → Feature implementation
│   ├── create_test.go
│   ├── fetch.go
│   └── archive.go
├── infra/               → External world
│   ├── repository/postgres/  → DB implementation
│   │   ├── note_repository.go
│   │   └── note_query.go
│   └── http/            → HTTP layer
│       ├── handler/     → Request handlers
│       │   ├── create_handler.go
│       │   ├── fetch_handler.go
│       │   └── archive_handler.go
│       └── route.go     → Route registration
└── module.go            → Module wiring
```

**Flow:** REST → Handler → UseCase → Repository (interface) ← PostgresRepository (impl)

**Dependencies point inward:** `infra/` → `usecase/` → `domain/`

> 🧠 Although examples use **HTTP**, the same structure applies to **any transport** — messaging, gRPC, CLI, etc.
> Transport changes, separation remains.

---

## Agent Guidelines

### 1. **Be Pragmatic**

* Only do what's needed to solve the current request.
* Avoid introducing abstractions or patterns unless explicitly requested.

### 2. **Follow the Pattern**

* **Domain** → Pure business logic, no dependencies
* **UseCase** → Application features, depends on domain interfaces
* **Infra** → External integrations (DB, HTTP, events), implements domain interfaces
* Each handler = one file, each usecase = one file

### 3. **Don't Over-Engineer**

* No factories, no layers for the sake of layering.
* If something isn't reused 3+ times, **don't generalize it**.

### 4. **Write for Production**

* Focus on maintainability and reliability.
* Avoid clever tricks — prefer clear and boring Go code.

### 5. **Test Every Use Case**

* Each `usecase.go` must have a corresponding `usecase_test.go`.
* Mock repository interfaces, test behavior, not implementation.

---

## Never Do

* ❌ Add new architectural layers or frameworks
* ❌ Bypass use cases (handler → repository directly)
* ❌ Write SQL inside use cases
* ❌ Add "helpers" or "utils" without purpose
* ❌ Refactor unrelated code
* ❌ Use global vars or `panic()`
* ❌ Cross-import between modules
* ❌ Put business logic in handlers or repositories

---

## Always Do

* ✅ Keep changes local to the module
* ✅ Reuse existing patterns and naming
* ✅ Respect the layer separation (domain ← usecase ← infra)
* ✅ Use explicit error handling
* ✅ Keep each file focused and readable
* ✅ Domain logic stays in domain/, queries in repository/

---

## Module Structure Rules

**domain/** - Business entities and repository contracts
* Pure Go, no external dependencies
* Defines interfaces that infra implements

**usecase/** - Application features
* Depends on domain interfaces
* Contains all business orchestration

**infra/** - External world
* **repository/postgres/** - Database queries and implementation
* **http/** - HTTP handlers and routes
* Implements domain interfaces

**module.go** - Dependency injection for the module

---

## Reference

* Entry point: `cmd/api/main.go`
* Example module: `internal/note/`
* Shared utilities: `internal/shared/`
* Reusable packages: `pkg/`

---

## Quick Reminder

Hatch isn't an enterprise framework — it's **a clean, modular foundation**.
Agents should think: *"What's the simplest change that keeps layers decoupled and modules independent?"*