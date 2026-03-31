# Hatch Agent Context

**Read README.md first** — full architecture details there.

## Structure

```
internal/<module>/
├── <module>.go        ← Public facade + transport declarations
├── domain/            ← Entities, value objects, repository interfaces
├── feature/<name>/    ← Self-contained use case
│   ├── feature.go     ← Constructor + dependency wiring
│   ├── event.go       ← Event handlers
│   ├── service.go     ← Domain Logic
│   ├── grpc.go        ← gRPC handlers
│   └── http.go        ← HTTP handlers
├── infra/             ← Repository implementations, external services
└── mocks/             ← Test doubles
```

## Rules

**Always:**
- Changes stay inside the active module
- Follow existing patterns and naming
- Wire features declaratively in `<module>.go`

**Never:**
- Add architectural layers
- Bypass use cases (handler → provider)
- Cross-import between modules
- Use globals or `panic()`
- Refactor unrelated code

## When in Doubt

1. Check `internal/note/` as reference
2. Pick the simplest change that keeps modules independent
