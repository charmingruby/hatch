# Modifying Existing Code

Guide to safely extend and modify existing modules while maintaining architectural integrity.

## Critical Rules

-  Preserve interface signatures (breaking changes ripple through codebase)
-  Follow existing patterns in similar files (`internal/note/` as reference)
-  Update tests alongside changes
-  Respect layer boundaries (never skip layers)
-  Generate mocks after interface changes (`make mock`)

## Common Modifications

### 1. Adding New Method to Existing Interface

Complete flow: Repository í Use Case í Handler

#### Example: Adding Delete Operation to Note Module

**Step 1: Update Repository Interface**

```go
// internal/note/repository/repository.go
package repository

import (
	"context"
	"HATCH_APP/internal/note/model"
)

type NoteRepository interface {
	Create(ctx context.Context, note model.Note) error
	FindByID(ctx context.Context, id string) (model.Note, error)
	List(ctx context.Context) ([]model.Note, error)
	Save(ctx context.Context, note model.Note) error
	Delete(ctx context.Context, id string) error  // ê New method
}
```

**Step 2: Implement in Repository**

```go
// internal/note/repository/postgres/note_repository.go
func (r *NoteRepo) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(deleteNote)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id)
	return err
}
```

**Step 3: Add SQL Query**

```go
// internal/note/repository/postgres/note_query.go
const (
	createNote   = "createNote"
	findNoteByID = "findNoteByID"
	listNotes    = "listNotes"
	saveNote     = "saveNote"
	deleteNote   = "deleteNote"  // ê Add constant
)

func noteQueries() map[string]string {
	return map[string]string{
		createNote: `
			INSERT INTO notes (id, title, content, archived, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`,
		findNoteByID: `
			SELECT id, title, content, archived, created_at, updated_at
			FROM notes WHERE id = $1
		`,
		listNotes: `
			SELECT id, title, content, archived, created_at, updated_at
			FROM notes ORDER BY created_at DESC
		`,
		saveNote: `
			UPDATE notes SET archived = $1, updated_at = $2
			WHERE id = $3
		`,
		deleteNote: "DELETE FROM notes WHERE id = $1",  // ê Add query
	}
}
```

**Step 4: Add DTO**

```go
// internal/note/dto/delete_note_dto.go
package dto

type DeleteNoteInput struct {
	ID string
}
```

**Step 5: Add Use Case Implementation**

```go
// internal/note/usecase/delete_note.go
package usecase

import (
	"context"
	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/shared/customerr"
)

func (u UseCase) DeleteNote(ctx context.Context, input dto.DeleteNoteInput) error {
	if err := u.noteRepo.Delete(ctx, input.ID); err != nil {
		return customerr.NewDatabaseError(err)
	}
	return nil
}
```

**Step 6: Add Use Case Tests**

```go
// internal/note/usecase/delete_note_test.go
package usecase_test

import (
	"errors"
	"testing"

	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/shared/customerr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DeleteNote(t *testing.T) {
	noteID := "note-123"

	t.Run("should delete successfully", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Delete", t.Context(), noteID).
			Return(nil).
			Once()

		err := s.usecase.DeleteNote(t.Context(), dto.DeleteNoteInput{ID: noteID})

		require.NoError(t, err)
	})

	t.Run("should return a DatabaseError when there is a datasource error", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Delete", t.Context(), noteID).
			Return(errors.New("unhealthy repo")).
			Once()

		err := s.usecase.DeleteNote(t.Context(), dto.DeleteNoteInput{ID: noteID})

		require.Error(t, err)

		var targetErr *customerr.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
```

**Step 7: Update Use Case Interface**

```go
// internal/note/usecase/usecase.go
package usecase

import (
	"context"
	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/note/repository"
)

type Service interface {
	CreateNote(ctx context.Context, input dto.CreateNoteInput) (dto.CreateNoteOutput, error)
	ListNotes(ctx context.Context, input dto.ListNotesInput) (dto.ListNotesOutput, error)
	ArchiveNote(ctx context.Context, input dto.ArchiveNoteInput) error
	DeleteNote(ctx context.Context, input dto.DeleteNoteInput) error  // ê New method
}

type UseCase struct {
	noteRepo repository.NoteRepository
}

func New(noteRepo repository.NoteRepository) UseCase {
	return UseCase{noteRepo: noteRepo}
}
```

**Step 8: Add HTTP Handler**

```go
// internal/note/http/endpoint/delete_note_endpoint.go
package endpoint

import (
	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"
	"errors"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) DeleteNote(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	e.log.InfoContext(ctx, "endpoint/DeleteNote: request received", "id", id)

	err := e.service.DeleteNote(ctx, dto.DeleteNoteInput{ID: id})
	if err != nil {
		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			e.log.ErrorContext(ctx, "endpoint/DeleteNote: database error", "error", databaseErr.Unwrap().Error())
			rest.SendInternalServerErrorResponse(c)
			return
		}

		e.log.ErrorContext(ctx, "endpoint/DeleteNote: unknown error", "error", err.Error())
		rest.SendInternalServerErrorResponse(c)
		return
	}

	e.log.InfoContext(ctx, "endpoint/DeleteNote: finished successfully")
	rest.SendNoContentResponse(c)
}
```

**Step 9: Register Route**

```go
// internal/note/http/endpoint/endpoint.go
func (e *Endpoint) Register(r *gin.Engine) {
	r.POST("/notes", e.CreateNote)
	r.GET("/notes", e.ListNotes)
	r.PATCH("/notes/:id/archive", e.ArchiveNote)
	r.DELETE("/notes/:id", e.DeleteNote)  // ê Add route
}
```

**Step 10: Generate Mocks & Test**

```bash
make mock  # Regenerate mocks with new interface methods
make test  # Run all tests
```

### 2. Adding New Field to Existing Model

When adding a new field to a domain model:

**Step 1: Update Model**

```go
// internal/note/model/note.go
type Note struct {
	ID        string
	Title     string
	Content   string
	Archived  bool
	Tags      []string  // ê New field
	CreatedAt time.Time
	UpdatedAt time.Time
}
```

**Step 2: Update Database Migration**

```sql
-- db/migrations/000003_adds_tags_to_notes.up.sql
ALTER TABLE notes ADD COLUMN tags TEXT[] DEFAULT '{}';
```

```sql
-- db/migrations/000003_adds_tags_to_notes.down.sql
ALTER TABLE notes DROP COLUMN tags;
```

**Step 3: Update Repository Queries**

```go
// Update all SELECT statements to include new field
func noteQueries() map[string]string {
	return map[string]string{
		createNote: `
			INSERT INTO notes (id, title, content, archived, tags, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`,
		findNoteByID: `
			SELECT id, title, content, archived, tags, created_at, updated_at
			FROM notes WHERE id = $1
		`,
		// ... update other queries
	}
}
```

**Step 4: Update Repository Methods**

```go
func (r *NoteRepo) Create(ctx context.Context, note model.Note) error {
	// ...
	_, err = stmt.ExecContext(ctx,
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Archived,
		pq.Array(&note.Tags),  // ê Handle array type
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	return err
}

func (r *NoteRepo) FindByID(ctx context.Context, id string) (model.Note, error) {
	// ...
	var note model.Note
	if err := stmt.QueryRowContext(ctx, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Archived,
		pq.Array(&note.Tags),  // ê Handle array type
		&note.CreatedAt,
		&note.UpdatedAt,
	); err != nil {
		// ...
	}
	return note, nil
}
```

**Step 5: Update DTOs (if needed)**

```go
// internal/note/dto/create_note_dto.go
type CreateNoteInput struct {
	Title   string
	Content string
	Tags    []string  // ê Add if accepting tags from input
}
```

**Step 6: Update Tests**

```go
func Test_CreateNote(t *testing.T) {
	t.Run("should create successfully", func(t *testing.T) {
		s := setup(t)

		s.repo.On("Create", t.Context(), mock.MatchedBy(func(n model.Note) bool {
			return n.Title == "Test" &&
				n.Content == "Content" &&
				len(n.Tags) == 2  // ê Verify new field
		})).Return(nil).Once()

		// ...
	})
}
```

**Step 7: Run Migration**

```bash
make migrate-up
```

### 3. Adding New Error Type

**Step 1: Define Error Type**

```go
// internal/shared/customerr/not_found.go
package customerr

import "fmt"

type NotFoundError struct {
	Resource string
	ID       string
}

func NewNotFoundError(resource, id string) *NotFoundError {
	return &NotFoundError{
		Resource: resource,
		ID:       id,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: %s", e.Resource, e.ID)
}
```

**Step 2: Use in Use Case**

```go
func (u UseCase) DeleteNote(ctx context.Context, input dto.DeleteNoteInput) error {
	note, err := u.noteRepo.FindByID(ctx, input.ID)
	if err != nil {
		return customerr.NewDatabaseError(err)
	}

	if note.ID == "" {
		return customerr.NewNotFoundError("note", input.ID)
	}

	if err := u.noteRepo.Delete(ctx, input.ID); err != nil {
		return customerr.NewDatabaseError(err)
	}

	return nil
}
```

**Step 3: Handle in HTTP Layer**

```go
func (e *Endpoint) DeleteNote(c *gin.Context) {
	// ...
	err := e.service.DeleteNote(ctx, dto.DeleteNoteInput{ID: id})
	if err != nil {
		var notFoundErr *customerr.NotFoundError
		if errors.As(err, &notFoundErr) {
			rest.SendNotFoundResponse(c, notFoundErr.Error())
			return
		}

		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			rest.SendInternalServerErrorResponse(c)
			return
		}

		rest.SendInternalServerErrorResponse(c)
		return
	}
	// ...
}
```

## Checklist Before Committing

- [ ] All layers updated (repository í use case í handler)
- [ ] Tests written for new functionality
- [ ] Mocks regenerated (`make mock`)
- [ ] All tests passing (`make test`)
- [ ] Interface signatures preserved (no breaking changes)
- [ ] Migration created (if database changes)
- [ ] Error handling added
- [ ] Logging added to handlers

## Reference

- Example module: [internal/note/](../../app/internal/note/)
- See [APPLICATION.md](../application.md) for architecture details
