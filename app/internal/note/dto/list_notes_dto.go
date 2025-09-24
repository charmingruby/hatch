package dto

import "PACK_APP/internal/note/model"

type ListNotesOutput struct {
	Notes []model.Note `json:"notes"`
}
