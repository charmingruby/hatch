package dto

import "HATCH_APP/internal/note/model"

type ListNotesOutput struct {
	Notes []model.Note `json:"notes"`
}
