package usecase

import (
	"HATCH_APP/internal/note/domain"
	"context"
)

type UseCase interface {
	Archive(ctx context.Context, id string) error
	Create(ctx context.Context, title, content string) (string, error)
	Fetch(ctx context.Context) ([]domain.Note, error)
}

type Service struct {
	repo domain.NoteRepository
}

func NewService(repo domain.NoteRepository) *Service {
	return &Service{repo: repo}
}
