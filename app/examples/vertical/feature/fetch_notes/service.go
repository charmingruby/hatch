package fetch_notes

import (
	"HATCH_APP/examples/vertical/domain"
	"HATCH_APP/internal/shared/errs"
	"context"
)

type UseCase interface {
	Execute(ctx context.Context) ([]domain.Note, error)
}

type Service struct {
	repo domain.NoteRepository
}

func NewService(repo domain.NoteRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Execute(ctx context.Context) ([]domain.Note, error) {
	notes, err := s.repo.List(ctx)
	if err != nil {
		return nil, errs.NewDatabaseError(err)
	}

	return notes, nil
}
