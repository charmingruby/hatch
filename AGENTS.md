# Agent Context

> **Quick Start:** Read the main README.md first for full architecture details.

---

## What is Hatch?

A **production-ready Go template** focused on:

- **Modularity** - Self-contained features
- **Clarity** - Flat structure, explicit dependencies
- **Speed** - Ship features fast without fighting architecture

---

## Core Rules

### Always Do

- Keep changes **inside the module** you're working on
- Follow existing patterns and naming
- Respect dependency flow
- Test every use case
- Keep files small and focused
- Name files inside each feature folder by their role (`handler.go`, `service.go`, etc.)

### Never Do

- Add new architectural layers
- Bypass use cases (handler → provider directly)
- Cross-import between modules
- Use global vars or `panic()`
- Refactor unrelated code

---

## When in Doubt

1. Check **README.md** for full architecture explanation
2. Look at `internal/note/` as reference implementation
3. Ask: "What's the **simplest change** that keeps modules independent?"

---

**Remember:** Hatch is not a framework — it's a clear, idiomatic Go foundation. Keep it simple.
