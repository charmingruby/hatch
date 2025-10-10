package repository

import (
	"HATCH_APP/internal/note/shared/model"
	"context"
)

type NoteRepo interface {
	FindByID(ctx context.Context, id string) (model.Note, error)
	Create(ctx context.Context, note model.Note) error
	List(ctx context.Context) ([]model.Note, error)
	Save(ctx context.Context, note model.Note) error
}
