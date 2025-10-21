# Hatch – AI Agent Context

> **Purpose:** Help agents contribute to Hatch effectively — fast, simple, and aligned with its philosophy.
> “Simplicity is the ultimate sophistication.” — Leonardo da Vinci

---

## Context

Hatch is a **pragmatic Go template** using a **modular, package-by-feature** architecture.
It favors **clarity over abstraction** and **production value over theory**.

The goal: **ship production-quality features fast**, without unnecessary layers or complexity.

---

## Architecture in One Glance

Each module represents a **bounded context** (e.g., `note`, `user`)
and contains its own features (create, fetch, etc.), each fully self-contained.

```
internal/note/
├── create/          → Feature: POST /notes
│   ├── handler.go   → Transport layer (HTTP, gRPC, messaging, etc.)
│   ├── dto.go       → Input/Output structs
│   └── usecase_test.go
├── fetch/           → Feature: GET /notes
│   ├── handler.go
│   ├── usecase.go
│   └── usecase_test.go
└── shared/
    ├── model/       → Domain entities
    └── repository/  → Repo interface + impl (Postgres)
```

**Each feature = one directory.**
No global services, no tangled layers, no abstractions unless necessary.

> 🧠 Although most examples use **HTTP**, the same structure applies to **any transport** — messaging, gRPC, CLI, etc.
> The delivery layer changes, but the separation (transport → use case → repository) remains the same.

---

## Agent Guidelines

### 1. **Be Pragmatic**

* Only do what’s needed to solve the current request.
* Avoid introducing abstractions or patterns unless explicitly requested.

### 2. **Follow the Pattern**

* Each feature lives inside its own folder (`create`, `fetch`, etc.).
* Transport (HTTP, gRPC, messaging) → parses, validates, and delegates.
* UseCases → contain business logic.
* Repositories → talk to the database.

### 3. **Don’t Over-Engineer**

* No factories, no layers for the sake of layering.
* If something isn’t reused 3+ times, **don’t generalize it**.

### 4. **Write for Production**

* Focus on maintainability and reliability.
* Avoid clever tricks — prefer clear and boring Go code.

### 5. **Test Every Use Case**

* Each `usecase.go` must have a corresponding `usecase_test.go`.
* Mock dependencies, test behavior, not implementation details.

---

## Never Do

* ❌ Add new architectural layers or frameworks
* ❌ Bypass use cases (transport → repository directly)
* ❌ Write SQL inside use cases
* ❌ Add “helpers” or “utils” without purpose
* ❌ Refactor unrelated code
* ❌ Use global vars or `panic()`
* ❌ Cross-import between modules

---

## Always Do

* ✅ Keep changes local to the feature
* ✅ Reuse existing patterns and naming
* ✅ Respect the directory structure
* ✅ Use explicit error handling
* ✅ Keep each file focused and readable

---

## Reference

* Entry point: `cmd/api/main.go`
* Example module: `internal/note/`
* DI: Uber Fx
* Architecture: Transport → UseCase → Repository → DB

---

## Quick Reminder

Hatch isn’t an enterprise framework — it’s **a fast, minimal foundation**.
Agents should think: *“What’s the simplest production-ready change that works?”*
