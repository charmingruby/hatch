package domain

import (
	"context"
)

type NoteRepository interface {
	FindByID(ctx context.Context, id string) (Note, error)
	Create(ctx context.Context, note Note) error
	List(ctx context.Context) ([]Note, error)
	Save(ctx context.Context, note Note) error
}
