# Hatch

Hatch is a pragmatic Go project template designed for rapid feature development and effortless service extraction. Modular, decoupled, and production-ready.

---

## Why Hatch?

* ✅ **Service-ready** – Extract modules to microservices without rewrites
* ✅ **Fast development** – Add features without fighting the architecture
* ✅ **Go-idiomatic** – Simple, explicit, and dependency-free
* ✅ **Screaming Architecture** – Project structure reveals business intent at first glance

---

## Architecture

Start simple and evolve as your domain grows.

### Module Organization

Hatch is **pattern-agnostic** and encourages choosing the right structure for each module's complexity. The architecture evolves with your domain understanding, not predetermined templates.

**Start simple, evolve deliberately.** Begin with flat structures and introduce layers only when complexity demands them. Mix approaches freely across modules—a CRUD module can stay flat while a complex workflow adopts vertical slices.

The key is **intentional evolution**: refactor when you feel the pain of the current structure, not because a pattern prescribes it.
There is a section to help with implementation details at `examples/vertical or packagebyfeat`.

#### Vertical Slice (Feature-First)

```text
internal/
├── order/
│   ├── domain/
│   │   ├── order.go              # Aggregate root
│   │   ├── order_item.go         # Value object
│   │   ├── repository.go         # Interface
│   │   └── events.go             # Domain events
│   ├── feature/
│   │   ├── create_order/
│   │   │   ├── service.go        # Use case logic
│   │   │   ├── handler.go        # HTTP entry point
│   │   │   └── dto.go            # Request/response
│   │   ├── cancel_order/
│   │   │   ├── service.go
│   │   │   └── handler.go
│   │   └── list_orders/
│   │       ├── service.go
│   │       ├── handler.go
│   │       └── query.go          # Read model
│   ├── infra/
│   │   ├── postgres/
│   │   │   └── repository.go     # Persistence implementation
│   │   └── events/
│   │       └── publisher.go      # Event infrastructure
│   └── module.go                 # Dependency injection
├── catalog/
│   └── ...
└── shared/                        # Cross-cutting concerns
    ├── event/
    ├── logging/
    └── errors/
```

**Use for:** Complex business rules, rich domain behavior, strong invariants, workflows with multiple steps.

**Key traits:**
- Domain entities encapsulate business rules
- Each `feature/` folder represents one use case with clear business purpose
- Infrastructure depends on domain, never the reverse

#### Simplified: Package by Feature (Flat Structure)

```text
internal/
├── note/
│   ├── note.go                   # Entity
│   ├── service.go                # Business logic
│   ├── handler.go                # HTTP handlers
│   ├── repository.go             # Interface
│   ├── event.go                  # Event definitions
│   └── module.go
├── tag/
│   └── ...
└── shared/
```

**Use for:** CRUD operations, simple validations, thin business logic, data orchestration.

**When to simplify:**
- Behavior is primarily data transformation
- Few or no business invariants to enforce
- Logic fits comfortably in a single service file
- Module is unlikely to grow significantly

---

## Decision Guide

Ask yourself:

**Choose Vertical Slice when:**
- Business rules are complex or will grow
- Multiple features operate on the same entity
- You need clear boundaries between use cases
- Domain logic deserves its own layer

**Choose Flat Structure when:**
- Module is essentially CRUD with validation
- All operations are similar in complexity
- Simpler structure aids velocity
- Over-engineering would obscure simplicity

**Remember:** You can start flat and refactor to vertical slices when complexity justifies it. The module boundary remains the same.

---

## Principles

* **Modular by default** – Each module can live independently or become its own service
* **Feature-driven** – Organize by what the code does, not by technical layers
* **Dependency inversion** – Domain defines interfaces, infrastructure implements them
* **Explicit over magical** – No hidden framework behaviors, just clear Go code
* **Intent-revealing structure** – Architecture screams the business domain, not technical details
* **Simplicity when possible** – Don't add layers until complexity demands them

---

**Simple. Explicit. Modular. Production-ready.**