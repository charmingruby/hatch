package dto

type CreateNoteInput struct {
	Title   string `json:"title"   binding:"required" validate:"required,gt=0"`
	Content string `json:"content" binding:"required" validate:"required,gt=0"`
}

type CreateNoteOutput struct {
	ID string
}
