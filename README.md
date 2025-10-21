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
│   ├── feature.go         → Wiring & registration
│   ├── handler.go         → HTTP: POST /notes
│   ├── usecase.go         → Business Logic
│   └── usecase_test.go    → Unit tests
│
├── fetch/                 ← Feature: Fetch notes
│   ├── feature.go
│   ├── handler.go         → HTTP: GET /notes
│   ├── usecase.go         → Business Logic
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
└── module.go              ← Module aggregator: Wires all features
```

### Visual Architecture

```text
┌────────────────────────────────────────────────────────────────┐
│                         NOTE MODULE                            │
│                                                                │
│  Each feature is a vertical slice with its own:                │
│  • Delivery                                                    │
│  • Business Logic                                              │
│  • Tests                                                       │
│                                                                │
│  ╔══════════════╗  ╔══════════════╗  ╔══════════════╗          │
│  ║   CREATE     ║  ║    FETCH     ║  ║   ARCHIVE    ║          │
│  ║   FEATURE    ║  ║   FEATURE    ║  ║   FEATURE    ║          │
│  ╠══════════════╣  ╠══════════════╣  ╠══════════════╣          │
│  ║ Handler      ║  ║ Handler      ║  ║ Handler      ║          │
│  ║ UseCase      ║  ║ UseCase      ║  ║ UseCase      ║          │
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

**2. Implement the business logic (usecase.go):**
```go
package like

import (
    "context"

    "HATCH_APP/internal/note/shared/repository"
    "HATCH_APP/internal/shared/errs"
)

type UseCaseInput struct {
    NoteID string
}

type UseCaseOutput struct {
    LikesCount int
}

type UseCase struct {
    repo repository.NoteRepo
}

func NewUseCase(repo repository.NoteRepo) UseCase {
    return UseCase{repo: repo}
}

func (u UseCase) Execute(ctx context.Context, input UseCaseInput) (UseCaseOutput, error) {
    count, err := u.repo.IncrementLikes(ctx, input.NoteID)
    if err != nil {
        return UseCaseOutput{}, errs.NewDatabaseError(err)
    }

    return UseCaseOutput{LikesCount: count}, nil
}
```

**3. Create the HTTP handler (handler.go):**
```go
package like

import (
    "errors"

    "HATCH_APP/internal/shared/errs"
    "HATCH_APP/internal/shared/http"
    "HATCH_APP/pkg/telemetry"

    "github.com/gin-gonic/gin"
)

type Request struct {
    NoteID string `json:"note_id" binding:"required" validate:"required,uuid4"`
}

type Response struct {
    LikesCount int `json:"likes_count"`
}

func RegisterRoute(log *telemetry.Logger, api *gin.RouterGroup, uc UseCase) {
    api.POST(":id/like", handle(log, uc))
}

func handle(log *telemetry.Logger, uc UseCase) gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()

        log.InfoContext(ctx, "endpoint/LikeNote: request received")

        req, err := http.ParseRequest[Request](c)
        if err != nil {
            log.ErrorContext(ctx, "endpoint/LikeNote: invalid request", "error", err.Error())
            http.SendBadRequestResponse(c, err.Error())
            return
        }

        output, err := uc.Execute(ctx, UseCaseInput{NoteID: req.NoteID})
        if err != nil {
            var dbErr *errs.DatabaseError
            if errors.As(err, &dbErr) {
                log.ErrorContext(ctx, "endpoint/LikeNote: database error", "error", dbErr.Unwrap().Error())
                http.SendInternalServerErrorResponse(c)
                return
            }

            log.ErrorContext(ctx, "endpoint/LikeNote: unexpected error", "error", err.Error())
            http.SendInternalServerErrorResponse(c)
            return
        }

        log.InfoContext(ctx, "endpoint/LikeNote: finished successfully")
        http.SendOkResponse(c, Response{LikesCount: output.LikesCount})
    }
}
```

**4. Wire it up (feature.go):**
```go
package like

import (
    "HATCH_APP/internal/note/shared/repository"
    "HATCH_APP/pkg/telemetry"

    "github.com/gin-gonic/gin"
)

func New(log *telemetry.Logger, api *gin.RouterGroup, repo repository.NoteRepo) {
    usecase := NewUseCase(repo)
    RegisterRoute(log, api, usecase)
}
```

**5. Wire it up (`like.go`):**
```go
package note

import (
	"HATCH_APP/internal/note/archive"
	"HATCH_APP/internal/note/create"
	"HATCH_APP/internal/note/fetch"
	"HATCH_APP/internal/note/shared/repository/postgres"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func Register(log *telemetry.Logger, r *gin.Engine, db *sqlx.DB) error {
    repo, err := postgres.NewNoteRepo(db)
    if err != nil {
        return err
    }

    api := r.Group("/api/v1/notes")

    create.New(log, api, repo)
    fetch.New(log, api, repo)
    archive.New(log, api, repo)
    like.New(log, api, repo) // ← Add this line

    return nil
}

var Module = fx.Module("note",
	fx.Invoke(register),
)
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
