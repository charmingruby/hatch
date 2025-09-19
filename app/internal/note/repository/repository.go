package repository

import (
	"context"

	"github.com/charmingruby/pack/internal/note/model"
)

type NoteRepository interface {
	FindByID(ctx context.Context, id string) (model.Note, error)
	Create(ctx context.Context, note model.Note) error
	List(ctx context.Context) ([]model.Note, error)
	Save(ctx context.Context, note model.Note) error
}
