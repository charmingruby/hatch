# Hatch

Hatch is a pragmatic Go project template designed for rapid feature development and effortless service extraction. Modular, decoupled, and production-ready.

---

## Why Hatch?

* ✅ **Service-ready** – Extract modules to microservices without rewrites
* ✅ **Fast development** – Add features without fighting the architecture
* ✅ **Go-idiomatic** – Simple, explicit, and dependency-free

---

## Architecture

Start simple and evolve as your domain grows.

### Module Patterns

Choose based on complexity. Mix both patterns in the same project.

#### Simpler Module Approach (Package by Feature)

```text
internal/note/
├── core
│   ├── service.go
│   ├── note.go
│   ├── event.go
│   └── repository.go
├── event
│   └── subscriber.go
├── http
│   ├── route.go
│   └── handler.go
├── infra
│   └── postgres
│       └── note_repository.go
├── mocks/
└── module.go
```

**Use for:** CRUD, simple business logic, orchestration.

#### Richer Domain Module (Vertical Slice)

```text
internal
├── order
│    ├── domain/
│    │   ├── order.go
│    │   ├── order_repository.go
│    │   ├── order_item.go
│    │   └── order_created_event.go
│    ├── feature/
│    │   ├── create_order/
│    │   │   ├── subscriber.go
│    │   │   ├── service.go
│    │   │   └── dto.go
│    │   ├── fetch_orders/
│    │   │   ├── handler.go
│    │   │   └── service.go
│    │   └── cancel_order/
│    ├── infra/
│    │   ├── repository/
│    │   │   └── postgres
│    │   │       └── order_repository.go
│    │   └── events/
│    │       └── order_created_event.go
│    ├── mocks/
│    └── module.go   
├── catalog
└── shared
```

**Use for:** Complex rules, rich behavior, strong invariants.

Each `feature/` folder = one complete use case.

---

## Principles

* **Modular by default** – Each module can live independently or become its own service
* **Feature-driven** – Organize by what the code does, not by technical layers
* **Dependency inversion** – Domain defines interfaces, infrastructure implements them
* **Explicit over magical** – No hidden framework behaviors, just clear Go code

---

**Simple. Explicit. Modular. Production-ready.**