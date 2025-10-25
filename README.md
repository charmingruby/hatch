# Hatch

Production-ready Go template built for speed and clarity.

Hatch is a pragmatic Go project template designed for rapid feature development and effortless service extraction. Modular, decoupled, and production-ready.

## Why Hatch?

- ✅ **Modular by design** - Self-contained features that don't interfere
- ✅ **Service-ready** - Extract modules to microservices without rewrites
- ✅ **Fast development** - Add features without fighting the architecture
- ✅ **Flat & clear** – No confusing layers or excessive nesting
- ✅ **Go Way** - Idiomatic patterns, explicit dependencies, no magic

## Architecture

```text
internal/note/             ← Self-contained feature (bounded context)
│
├── domain/                ← Business entities & rules
│   ├── note.go
│   └── repository.go      → Interfaces for external dependencies
│
├── usecase/               ← Application logic
│   ├── create.go
│   ├── fetch.go
│   ├── archive.go
│   └── usecase.go         → Shared setup or interface
│
├── provider/              ← External service providers
│   ├── postgres/          → Database implementation
│   │   ├── note_repository.go
│   │   └── note_query.go
│   ├── redis/             → Cache implementation
│   ├── rabbitmq/          → Queue implementation
│   └── sendgrid/          → Email service implementation
│
├── http/rest/             ← HTTP layer
│   ├── handler/
│   │   ├── create_handler.go
│   │   ├── fetch_handler.go
│   │   └── archive_handler.go
│   └── route.go
│
└── module.go              ← Dependency wiring for the feature
```

## Project Structure

```text
.
├── cmd/api/main.go              # Entry point
├── config/                      # Configuration
├── db/migration/                # Database migrations
├── internal/
│   ├── note/                    # Bounded context
│   │   ├── domain/              → Business logic & interfaces
│   │   ├── usecase/             → Application logic
│   │   ├── provider/            → External service providers
│   │   ├── http/                → HTTP handlers
│   │   └── module.go            → Dependency injection
│   └── shared/                  # Cross-cutting concerns
│       ├── errs/                → Error types
│       └── http/                → HTTP utilities
├── pkg/                         # Reusable packages
└── test/gen/                    # Generated mocks
```

## Layer Responsibilities

**domain/** - Business entities and interface contracts
- Core business logic
- Entity definitions
- Repository and service interfaces
- No external dependencies

**usecase/** - Application logic orchestration
- Coordinates business operations
- Implements use cases
- Depends only on domain interfaces

**provider/** - External service implementations
- Database repositories (Postgres, MySQL)
- Cache providers (Redis, Memcached)
- Message queues (RabbitMQ, Kafka)
- Third-party APIs (SendGrid, Stripe, Twilio)
- Storage services (S3, MinIO)
- All external integrations live here

**http/** - HTTP presentation layer
- Request/response handling
- Route definitions
- Input validation

## Key Benefits

**Rapid Development**
- Add features in isolated modules
- No cascade of changes across codebase
- Clear boundaries = predictable changes

**Service Extraction**
- Move `internal/note/` to new repo
- Add `cmd/main.go`, done
- Zero architectural rewrites

**Provider Flexibility**
- Swap implementations without touching business logic
- Mock providers easily for testing
- Deploy with different providers per environment

**Go Way**
- Flat over nested – Shallow folder hierarchy for quick navigation
- Purposeful packages – Clear boundaries and naming
- Explicit dependencies – No hidden magic

## Core Principles

**Modularity** - Features are self-contained  
**Decoupling** - Layers depend on interfaces  
**Extractability** - Modules → Microservices naturally  
**Provider Pattern** - External dependencies are pluggable

---

Simple. Explicit. Modular. Production-ready.