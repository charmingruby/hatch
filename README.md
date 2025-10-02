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
│  │   REST    │  │  │   gRPC    │  │  │   Event   │  │
│  └─────┬─────┘  │  └─────┬─────┘  │  └─────┬─────┘  │
│        │        │        │        │        │        │
│  ┌─────▼─────┐  │  ┌─────▼─────┐  │  ┌─────▼─────┐  │
│  │ Use Case  │  │  │ Use Case  │  │  │ Use Case  │  │
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
- **Minimal Coupling** - Shared code lives in common packages
- **Service Ready** - Any module can become a microservice without refactoring

## Documentation

- **[Application Guidelines](docs/app.md)** - Complete development guide with Go best practices
- **[Project Structure Documentation](docs/project-structure.md)** - Repository organization and structure details
- **[AGENTS.md](AGENTS.md)** - Comprehensive context for AI agents about architecture and patterns
