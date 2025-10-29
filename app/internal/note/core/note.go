package core

import (
	"time"

	"HATCH_APP/pkg/id"
)

type Note struct {
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	ID        string     `json:"id"         db:"id"`
	Title     string     `json:"title"      db:"title"`
	Content   string     `json:"content"    db:"content"`
	Archived  bool       `json:"archived"   db:"archived"`
}

func NewNote(title, content string) Note {
	return Note{
		ID:        id.New(),
		Title:     title,
		Content:   content,
		Archived:  false,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
}
