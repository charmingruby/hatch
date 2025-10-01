# Hatch

> Production-ready Go template for building scalable, maintainable applications from day one.

Hatch is a **battle-tested Go project template** built on modular architecture principles. Each feature is an independent module that can be easily removed or extracted into a separate service. Designed to accelerate development across all project types—from quick PoCs to enterprise-grade systems.

## Why Hatch?

- **Battle-tested architecture** - Proven patterns for real-world production systems
- **Flexible & scalable** - Works for small proofs-of-concept and large-scale applications
- **Extends for any use case** - Monorepo support, infrastructure as code, event-driven architecture
- **Modern by default** - REST APIs, event-driven architecture, Docker, CI/CD ready
- **Clean third-party integration** - Modular, maintainable, easy to swap providers
- **Test-first approach** - Built-in testing conventions, mocks, and tooling
- **Go idioms enforced** - Follows Go best practices and community standards

## Architecture

### Modular Design

Self-contained modules that can be easily extracted into separate services:

```
Application
├─────────────────┬─────────────────┬─────────────────┐
│   Module A      │   Module B      │   Module C      │
│                 │                 │                 │
│  ┌───────────┐  │  ┌───────────┐  │  ┌───────────┐  │
│  │    HTTP   │  │  │    HTTP   │  │  │    HTTP   │  │
│  └─────┬─────┘  │  └─────┬─────┘  │  └─────┬─────┘  │
│        │        │        │        │        │        │
│  ┌─────▼─────┐  │  ┌─────▼─────┐  │  ┌─────▼─────┐  │
│  │  Use Case │  │  │  Use Case │  │  │  Use Case │  │
│  └─────┬─────┘  │  └─────┬─────┘  │  └─────┬─────┘  │
│        │        │        │        │        │        │
│  ┌─────▼─────┐  │  ┌─────▼─────┐  │  ┌─────▼─────┐  │
│  │Repository │  │  │Repository │  │  │Repository │  │
│  └───────────┘  │  └───────────┘  │  └───────────┘  │
│                 │                 │                 │
└─────────────────┴─────────────────┴─────────────────┘
         │                 │                 │
         └─────────────────┼─────────────────┘
                           │
                      ┌────▼────┐
                      │Database │
                      └─────────┘
```

### Core Principles

- **Module Independence** - Each module is self-contained and can be removed or extracted
- **Clear Boundaries** - Modules communicate through well-defined interfaces
- **Minimal Coupling** - Shared code lives in `pkg/` and `internal/shared/`
- **Service Ready** - Any module can become a microservice without refactoring

## Project Structure

### Repository Layout

```
├── cmd/api/            # Application entry point (main.go)
├── config/             # Configuration management
├── internal/           # Private application code
│   ├── note/           # Example feature module
│   └── shared/         # Cross-cutting concerns
├── pkg/                # Public reusable libraries
├── test/               # Test utilities and mocks
├── db/                 # Database migrations
└── Makefile            # Development commands
```

### Feature Module Layout

Each feature inside `internal/` follows this structure:

```
internal/MODULE/
  MODULE.go              # Barrel file - dependency wiring
  model/                 # Domain entities
  dto/                   # Data Transfer Objects
  usecase/               # Business logic + tests
  repository/            # Data access contracts + implementations
  http/endpoint/         # HTTP handlers
  messaging/             # Event publishers/subscribers (optional)
  external/              # Third-party integrations (optional)
```

**Example:** See `app/internal/note/` for a complete reference implementation.

## Documentation

### For Developers
- **[Application Guidelines](app/docs/guidelines.md)** - Complete development guide with Go best practices
- **[Structure Documentation](docs/structure.md)** - Repository organization for single-app and monorepo setups

### For AI Agents
- **[AGENTS.md](AGENTS.md)** - Comprehensive context about architecture, patterns, and code modification guidelines
