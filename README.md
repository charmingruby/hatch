# Hatch

> Production-ready Go template that delivers value fast without the complexity.

"Simplicity is the ultimate sophistication." — Leonardo da Vinci

Hatch is a **pragmatic Go project template** built on **package by feature** principles. It solves the same problems as complex "enterprise" architectures, but with radical simplicity. Get the benefits of clean architecture without the over-engineering—ship features faster while maintaining quality.

## Why Hatch?

**The Problem with Most Templates:**
- Over-engineered from day one (10+ layers, abstractions everywhere)
- Takes hours to understand before writing a single line
- More time maintaining architecture than delivering features

**Hatch's Philosophy:**
- ✅ **Simple by default** - Start minimal, grow naturally
- ✅ **Fast time-to-value** - Ship your first feature in minutes
- ✅ **Package by feature** - Find everything in one place
- ✅ **Battle-tested** - Proven in production, not theoretical
- ✅ **Easy to understand** - Onboard in 30 minutes

**Flexible for any project type:**
- POCs & Labs - Quick experiments without architectural overhead
- Microservices - Single service with clear boundaries
- Monorepos - Multiple services with shared infrastructure
- IaC Integration - Organize Terraform/Pulumi alongside your code
- Production Apps - Battle-tested patterns that scale

## Architecture

### Package by Feature: Simple Yet Powerful

Everything related to a feature lives together:

```text
internal/note/             ← Single cohesive module
│
├── create/                ← Feature: Create notes
│   ├── create.go          → Wiring & registration
│   ├── handler.go         → HTTP: POST /notes
│   ├── usecase.go         → Business Logic
│   ├── dto.go             → Input/Output types
│   └── usecase_test.go    → Unit tests
│
├── fetch/                 ← Feature: Fetch notes
│   ├── fetch.go
│   ├── handler.go         → HTTP: GET /notes
│   ├── usecase.go         → Business Logic
│   ├── dto.go             → Input/Output types
│   └── usecase_test.go
│
├── shared/                ← Shared within note module only
│   ├── model/
│   │   └── note.go        → Entity: Note domain model
│   └── repository/
│       ├── repository.go  → Interface: NoteRepository
│       └── postgres/
│           └── note_repository.go → Implementation
│
└── note.go                ← Module aggregator: Wires all features
```

### Visual Architecture

```text
┌────────────────────────────────────────────────────────────────┐
│                         NOTE MODULE                            │
│                                                                │
│  Each feature is a vertical slice with its own:                │
│  • Delivery                                                    │
│  • Business Logic                                              │
│  • DTOs                                                        │
│  • Tests                                                       │
│                                                                │
│  ╔══════════════╗  ╔══════════════╗  ╔══════════════╗          │
│  ║   CREATE     ║  ║    FETCH     ║  ║   ARCHIVE    ║          │
│  ║   FEATURE    ║  ║   FEATURE    ║  ║   FEATURE    ║          │
│  ╠══════════════╣  ╠══════════════╣  ╠══════════════╣          │
│  ║ Handler      ║  ║ Handler      ║  ║ Handler      ║          │
│  ║ UseCase      ║  ║ UseCase      ║  ║ UseCase      ║          │
│  ║ DTOs         ║  ║ DTOs         ║  ║ DTOs         ║          │
│  ║ Tests        ║  ║ Tests        ║  ║ Tests        ║          │
│  ╚══════╤═══════╝  ╚══════╤═══════╝  ╚══════╤═══════╝          │
│         │                 │                 │                  │
│         └─────────────────┼─────────────────┘                  │
│                           │                                    │
│         ┌─────────────────▼─────────────────┐                  │
│         │    SHARED (within module)         │                  │
│         │  • Note Entity (model)            │                  │
│         │  • Repository Interface           │                  │
│         │  • PostgreSQL Implementation      │                  │
│         └───────────────────────────────────┘                  │
└────────────────────────────────────────────────────────────────┘
                           │
                   ┌───────▼────────┐
                   │   PostgreSQL   │
                   │    Database    │
                   └────────────────┘
```

### Why This Works

**❌ Traditional "Enterprise" Approach:**
```text
internal/
├── domain/
│   ├── entity/
│   ├── value_object/
│   └── aggregate/
├── application/
│   ├── service/
│   ├── dto/
│   └── mapper/
├── infrastructure/
│   ├── persistence/
│   ├── messaging/
│   └── adapter/
└── presentation/
    ├── http/
    └── grpc/

Result:
• 8+ directories to navigate for one feature
• 30+ files for a simple CRUD
• Hours to understand where things go
```

**✅ Hatch's Approach:**
```text
internal/
├── note/
│   ├── create/        ← Everything for "create note" in ONE place
│   ├── fetch/         ← Everything for "fetch notes" in ONE place
│   └── shared/        ← Only what's actually shared
└── user/
    ├── register/      ← Everything for "register user" in ONE place
    └── login/         ← Everything for "login user" in ONE place

Result:
• 1 directory per feature
• 4-5 files for a complete feature
• 5 minutes to understand and start coding
```

## Core Principles

1. **Package by Feature** - Code organized by business capabilities, not technical layers
2. **Feature Independence** - Each feature is self-contained and can be modified/removed independently
3. **Pragmatic Sharing** - Share code only when proven necessary (3+ uses)
4. **Scalability Path** - Start monolith, extract to microservices when needed (no rewrites)

## Quick Start

**1. Clone and setup**

Clone the repository and copy the environment file:
```bash
git clone https://github.com/charmingruby/hatch
cd hatch/app
cp .env.example .env
```

**2. Replace module name**

Replace `HATCH_APP` with your actual Go module path (e.g., `github.com/yourorg/project`) throughout the project.

**3. Start everything**

Run Docker Compose to start the application, database, and apply migrations automatically:
```bash
docker compose up -d
```

That's it! The application is now running at `http://localhost:3333`.

**4. Test it**

Verify everything is working:
```bash
# Health check
curl http://localhost:3333/livez

# Create your first note
curl -X POST http://localhost:3333/api/v1/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "My first note", "content": "Hello Hatch!"}'
```

### Local Development (without Docker)

If you prefer running the application locally while using Docker only for dependencies:

Start PostgreSQL in Docker:
```bash
docker compose up -d postgres
```

Run migrations:
```bash
make migrate-up
```

Start the application:
```bash
air
```

## Adding a Feature in 5 Minutes

Let's add a "like note" feature. Here's everything you need:

**1. Create the directory:**
```bash
mkdir -p internal/note/like
```

**2. Define types (`dto.go`):**
```go
package like

type Input struct {
    NoteID string `json:"note_id" binding:"required" validate:"required"`
}

type Output struct {
    LikesCount int `json:"likes_count"`
}
```

**3. Implement business logic (`usecase.go`):**
```go
package like

import (
    "context"
    "HATCH_APP/internal/note/shared/repository"
    "HATCH_APP/internal/shared/customerr"
)

type UseCase struct {
    repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
    return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context, input Input) (Output, error) {
    count, err := u.repo.IncrementLikes(ctx, input.NoteID)
    if err != nil {
        return Output{}, customerr.NewDatabaseError(err)
    }

    return Output{LikesCount: count}, nil
}
```

**4. Create HTTP handler (`handler.go`):**
```go
package like

import (
    "HATCH_APP/internal/shared/customerr"
    "HATCH_APP/internal/shared/transport/http"
    "HATCH_APP/pkg/logger"
    "errors"
    "github.com/gin-gonic/gin"
)

func registerRoute(
    log *logger.Logger,
    api *gin.RouterGroup,
    uc UseCase,
) {
    api.POST(":id/like", func(c *gin.Context) {
        ctx := c.Request.Context()

        req, err := http.ParseRequest[Input](c)
        if err != nil {
            http.SendBadRequestResponse(c, err.Error())
            return
        }

        output, err := uc.Execute(ctx, *req)
        if err != nil {
            var dbErr *customerr.DatabaseError
            if errors.As(err, &dbErr) {
                http.SendInternalServerErrorResponse(c)
                return
            }
            http.SendInternalServerErrorResponse(c)
            return
        }

        http.SendOkResponse(c, output)
    })
}
```

**5. Wire it up (`like.go`):**
```go
package like

import (
    "HATCH_APP/internal/note/shared/repository"
    "HATCH_APP/pkg/logger"
    "github.com/gin-gonic/gin"
)

func New(
    log *logger.Logger,
    api *gin.RouterGroup,
    repo repository.NoteRepo,
) {
    registerRoute(log, api, NewUseCase(repo))
}
```

**6. Register in module (`internal/note/note.go`):**
```go
func register(log *logger.Logger, r *gin.Engine, db *sqlx.DB) error {
    repo, err := postgres.NewNoteRepo(db)
    if err != nil {
        return err
    }

    api := r.Group("/api/v1/notes")

    create.New(log, api, repo)
    fetch.New(log, api, repo)
    archive.New(log, api, repo)
    like.New(log, api, repo)  // ← Add this line

    return nil
}
```

**Done!** Feature ready in 5 minutes. Test it:
```bash
curl -X POST http://localhost:8080/api/v1/notes/123/like \
  -H "Content-Type: application/json" \
  -d '{"note_id": "123"}'
```

Compare this to enterprise templates:
- ❌ 8+ files across different directories
- ❌ Multiple abstraction layers to understand
- ❌ 30+ minutes of boilerplate

With Hatch:
- ✅ 5 files in ONE directory
- ✅ Clear, simple structure
- ✅ 5 minutes from idea to working code

## Documentation

- **[AGENTS.md](AGENTS.md)** - Comprehensive context for AI assistants about architecture and patterns
