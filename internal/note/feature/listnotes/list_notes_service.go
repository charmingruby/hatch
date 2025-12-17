package listnotes

import (
	"HATCH_APP/internal/common/errs"
	"HATCH_APP/internal/note/domain"
	"context"
)

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
