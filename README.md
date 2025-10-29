# Hatch

Production-ready Go template built for speed and clarity.

Hatch is a pragmatic Go template for rapid feature development and easy service extraction. Modular, decoupled, and production-ready.

## Why Hatch?

* ✅ **Modular** – Self-contained bounded contexts
* ✅ **Service-ready** – Extract modules to microservices easily
* ✅ **Fast development** – Add features without fighting the architecture
* ✅ **Flat & clear** – No deep nesting or confusing layers
* ✅ **Go idiomatic** – Explicit dependencies, no hidden magic

## Architecture

```text
app/
├── cmd/api/main.go           # Entry point
├── config/config.go          # Configuration
├── db/migration/             # SQL migrations
├── internal/
│   ├── note/                 # Bounded Context: Notes
│   │   ├── domain/           # Shared entities & interfaces
│   │   ├── feature/          # Vertical Slices per use case
│   │   │   ├── create_note/
│   │   │   ├── fetch_notes/
│   │   │   └── archive_note/
│   │   ├── infra/            # External implementations (DB, cache, etc.)
│   │   ├── mocks/            # Module specific mocks
│   │   └── module.go         # Dependency wiring
│   └── shared/               # Cross-cutting concerns (errors, HTTP utils)
└── pkg/                      # Reusable packages (DB client, logger, validator)
```

## Layer Responsibilities

* **domain/** – Entities and business rules, no external dependencies
* **feature/** – Self-contained vertical slice per use case
* **infra/** – External services (DB, cache, queues, APIs)
* **shared/** – Utilities and cross-cutting concerns

## Template overrides

1. Replace `HATCH_APP` with your Go module path (e.g., `github.com/yourorg/project`)
2. Copy `.env.example` → `.env` and fill required values
3. Update Docker image name in `.github/workflows/*.yml` and configure repository secrets

## Benefits

* Rapid development with isolated features
* Easy extraction to microservices
* Flexible providers for testing and deployment
* Clear, flat structure for quick navigation

## Principles

* **Modularity** – Bounded contexts
* **Decoupling** – Depend on interfaces, not implementations
* **Extractability** – Modules can become microservices
* **Vertical Slice** – Each use case encapsulates its own logic
* **Provider Pattern** – Pluggable external dependencies

---

Simple. Explicit. Modular. Production-ready.