package createnote

import (
	"HATCH_APP/internal/note/domain"
)

type Feature struct {
	service *Service
}

func New(repo domain.NoteRepository) *Feature {
	return &Feature{
		service: NewService(repo),
	}
}
