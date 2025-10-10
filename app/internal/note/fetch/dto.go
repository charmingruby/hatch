package fetch

import "HATCH_APP/internal/note/shared/model"

type Output struct {
	Notes []model.Note `json:"notes"`
}
